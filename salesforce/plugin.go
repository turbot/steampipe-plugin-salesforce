package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-salesforce"

type contextKey string

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.From(getFieldFromSObjectMapByColumnName).NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		SchemaMode:   plugin.SchemaModeDynamic,
		TableMapFunc: pluginTableDefinitions,
	}

	return p
}

type dynamicMap struct {
	cols       []*plugin.Column
	keyColumns plugin.KeyColumnSlice
}

func pluginTableDefinitions(ctx context.Context, p *plugin.Plugin) (map[string]*plugin.Table, error) {
	// If unable to connect to salesforce instance, log warning and abort dynamic table creation
	client, err := connectRaw(ctx, p.ConnectionManager, p.Connection)
	if err != nil {
		// do not abort the plugin as static table needs to be generated
		plugin.Logger(ctx).Error("salesforce.pluginTableDefinitions", "connection_error: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
	}

	staticTables := []string{"Account", "AccountContactRole", "Asset", "Contact", "Contract", "Lead", "Opportunity", "OpportunityContactRole", "Order", "Pricebook2", "Product2", "User"}

	var wgd sync.WaitGroup
	wgd.Add(len(staticTables))
	dynamicColumnsMap := map[string]dynamicMap{}
	for _, st := range staticTables {
		go func(staticTable string) {
			defer wgd.Done()
			dynamicCols, dynamicKeyColumns := dynamicColumns(ctx, client, staticTable, p)
			// Ignore if the requested Salesforce object is not present.
			dynamicColumnsMap[staticTable] = dynamicMap{dynamicCols, dynamicKeyColumns}
		}(st)
	}
	wgd.Wait()

	// Initialize tables with static tables with static and dynamic columns(if credentials are set)
	tables := map[string]*plugin.Table{
		"salesforce_account":                  SalesforceAccount(ctx, dynamicColumnsMap["Account"], p),
		"salesforce_account_contact_role":     SalesforceAccountContactRole(ctx, dynamicColumnsMap["AccountContactRole"], p),
		"salesforce_asset":                    SalesforceAsset(ctx, dynamicColumnsMap["Asset"], p),
		"salesforce_contact":                  SalesforceContact(ctx, dynamicColumnsMap["Contact"], p),
		"salesforce_contract":                 SalesforceContract(ctx, dynamicColumnsMap["Contract"], p),
		"salesforce_lead":                     SalesforceLead(ctx, dynamicColumnsMap["Lead"], p),
		"salesforce_opportunity":              SalesforceOpportunity(ctx, dynamicColumnsMap["Opportunity"], p),
		"salesforce_opportunity_contact_role": SalesforceOpportunityContactRole(ctx, dynamicColumnsMap["OpportunityContactRole"], p),
		"salesforce_order":                    SalesforceOrder(ctx, dynamicColumnsMap["Order"], p),
		"salesforce_pricebook":                SalesforcePricebook(ctx, dynamicColumnsMap["Pricebook2"], p),
		"salesforce_product":                  SalesforceProduct(ctx, dynamicColumnsMap["Product2"], p),
		"salesforce_user":                     SalesforceUser(ctx, dynamicColumnsMap["User"], p),
	}

	var re = regexp.MustCompile(`\d+`)
	var substitution = ``
	salesforceTables := []string{}
	config := GetConfig(p.Connection)
	if config.Tables != nil && len(*config.Tables) > 0 {
		for _, tableName := range *config.Tables {
			pluginTableName := "salesforce_" + strcase.ToSnake(re.ReplaceAllString(tableName, substitution))
			if _, ok := tables[pluginTableName]; !ok {
				salesforceTables = append(salesforceTables, tableName)
			}
		}
	}

	if client == nil {
		plugin.Logger(ctx).Error("salesforce.pluginTableDefinitions", "client_not_found: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
		return tables, nil
	}

	var wg sync.WaitGroup
	wg.Add(len(salesforceTables))
	for i, table := range salesforceTables {
		tableName := "salesforce_" + strcase.ToSnake(re.ReplaceAllString(table, substitution))
		ctx = context.WithValue(ctx, contextKey("PluginTableName"), tableName)
		plugin.Logger(ctx).Debug("salesforce.pluginTableDefinitions", "SALESFORCE_OBJECT_NAME", table, "STEAMPIPE_TABLE_NAME", tableName)
		ctx = context.WithValue(ctx, contextKey("SalesforceTableName"), table)
		if tables[tableName] != nil {
			wg.Done()
			continue
		}
		go func(i int) {
			defer wg.Done()
			table := generateDynamicTables(ctx, p)
			// Ignore if the requested Salesforce object is not present.
			if table != nil {
				tables[tableName] = table
			}
		}(i)
	}
	wg.Wait()
	return tables, nil
}

func generateDynamicTables(ctx context.Context, p *plugin.Plugin) *plugin.Table {

	client, err := connectRaw(ctx, p.ConnectionManager, p.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.generateDynamicTables", "connection_error", err)
		return nil
	}

	// Get the query for the metric (required)
	salesforceTableName := ctx.Value(contextKey("SalesforceTableName")).(string)
	tableName := ctx.Value(contextKey("PluginTableName")).(string)

	sObjectMeta := client.SObject(salesforceTableName).Describe()
	if sObjectMeta == nil {
		plugin.Logger(ctx).Error("salesforce.generateDynamicTables", fmt.Sprintf("Object %s not found in salesforce", salesforceTableName))
		return nil
	}

	// Top columns
	cols := []*plugin.Column{}
	// Key columns
	keyColumns := plugin.KeyColumnSlice{}

	salesforceObjectMetadata := *sObjectMeta
	salesforceObjectMetadataAsByte, err := json.Marshal(salesforceObjectMetadata["fields"])
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.generateDynamicTables", "json marshal error", err)
	}

	salesforceObjectFields := []map[string]interface{}{}
	// var queryColumns []string
	err = json.Unmarshal(salesforceObjectMetadataAsByte, &salesforceObjectFields)
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.generateDynamicTables", "json unmarshal error", err)
	}

	for _, properties := range salesforceObjectFields {
		if properties["name"] == nil {
			continue
		}
		fieldName := properties["name"].(string)
		compoundFieldName := properties["compoundFieldName"]
		if compoundFieldName != nil && compoundFieldName.(string) != fieldName {
			continue
		}

		// queryColumns = append(queryColumns, fieldName)
		if properties["soapType"] == nil {
			continue
		}
		soapType := strings.Split((properties["soapType"]).(string), ":")
		fieldType := soapType[len(soapType)-1]

		// Coloumn dynamic generation
		columnFieldName := strcase.ToSnake(fieldName)
		column := plugin.Column{
			Name:        columnFieldName,
			Description: fmt.Sprintf("The %s.", properties["label"].(string)),
			Transform:   transform.FromP(getFieldFromSObjectMap, fieldName),
		}

		// Set column type based on the `soapType` from salesforce schema
		switch fieldType {
		case "string", "ID":
			column.Type = proto.ColumnType_STRING
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", "<>"}})
		case "date", "dateTime":
			column.Type = proto.ColumnType_TIMESTAMP
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		case "boolean":
			column.Type = proto.ColumnType_BOOL
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", "<>"}})
		case "double":
			column.Type = proto.ColumnType_DOUBLE
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		case "int":
			column.Type = proto.ColumnType_INT
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		default:
			column.Type = proto.ColumnType_JSON
		}
		cols = append(cols, &column)

	}

	// query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(queryColumns, ", "), salesforceTableName)

	Table := plugin.Table{
		Name:        tableName,
		Description: fmt.Sprintf("Salesforce %s.", salesforceObjectMetadata["label"]),
		List: &plugin.ListConfig{
			KeyColumns: keyColumns,
			Hydrate:    listSalesforceObjectsByTable(salesforceTableName, cols),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSalesforceObjectbyID(salesforceTableName),
		},
		Columns: cols,
	}

	return &Table
}
