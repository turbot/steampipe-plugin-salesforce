package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

const pluginName = "steampipe-plugin-salesforce"

var d *plugin.QueryData

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
	cols              []*plugin.Column
	keyColumns        plugin.KeyColumnSlice
	salesforceColumns map[string]string
}

func pluginTableDefinitions(ctx context.Context, connection *plugin.Connection) (map[string]*plugin.Table, error) {
	// If unable to connect to salesforce instance, log warning and abort dynamic table creation

	client, err := connectRaw(ctx, d.ConnectionManager, d.Connection)
	if err != nil {
		// do not abort the plugin as static table needs to be generated
		plugin.Logger(ctx).Warn("salesforce.pluginTableDefinitions", "connection_error: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
	}

	staticTables := []string{"Account", "AccountContactRole", "Asset", "Contact", "Contract", "Lead", "Opportunity", "OpportunityContactRole", "Order", "Pricebook2", "Product2", "User", "PermissionSet", "PermissionSetAssignment", "ObjectPermissions"}

	dynamicColumnsMap := map[string]dynamicMap{}
	var mapLock sync.Mutex

	// If Salesforce client was obtained, don't generate dynamic columns for
	// defined static tables
	if client != nil {
		var wgd sync.WaitGroup
		wgd.Add(len(staticTables))
		for _, st := range staticTables {
			go func(staticTable string) {
				defer wgd.Done()
				dynamicCols, dynamicKeyColumns, salesforceCols := dynamicColumns(ctx, client, staticTable, d.Table.Plugin)
				mapLock.Lock()
				dynamicColumnsMap[staticTable] = dynamicMap{dynamicCols, dynamicKeyColumns, salesforceCols}
				defer mapLock.Unlock()
			}(st)
		}
		wgd.Wait()
	}

	// Initialize tables with static tables with static and dynamic columns(if credentials are set)
	tables := map[string]*plugin.Table{
		"salesforce_account":                   SalesforceAccount(ctx, dynamicColumnsMap["Account"], d.Table.Plugin),
		"salesforce_account_contact_role":      SalesforceAccountContactRole(ctx, dynamicColumnsMap["AccountContactRole"], d.Table.Plugin),
		"salesforce_asset":                     SalesforceAsset(ctx, dynamicColumnsMap["Asset"], d.Table.Plugin),
		"salesforce_contact":                   SalesforceContact(ctx, dynamicColumnsMap["Contact"], d.Table.Plugin),
		"salesforce_contract":                  SalesforceContract(ctx, dynamicColumnsMap["Contract"], d.Table.Plugin),
		"salesforce_lead":                      SalesforceLead(ctx, dynamicColumnsMap["Lead"], d.Table.Plugin),
		"salesforce_object_permission":         SalesforceObjectPermission(ctx, dynamicColumnsMap["ObjectPermissions"], d.Table.Plugin),
		"salesforce_opportunity":               SalesforceOpportunity(ctx, dynamicColumnsMap["Opportunity"], d.Table.Plugin),
		"salesforce_opportunity_contact_role":  SalesforceOpportunityContactRole(ctx, dynamicColumnsMap["OpportunityContactRole"], d.Table.Plugin),
		"salesforce_order":                     SalesforceOrder(ctx, dynamicColumnsMap["Order"], d.Table.Plugin),
		"salesforce_permission_set":            SalesforcePermissionSet(ctx, dynamicColumnsMap["PermissionSet"], d.Table.Plugin),
		"salesforce_permission_set_assignment": SalesforcePermissionSetAssignment(ctx, dynamicColumnsMap["PermissionSetAssignment"], d.Table.Plugin),
		"salesforce_pricebook":                 SalesforcePricebook(ctx, dynamicColumnsMap["Pricebook2"], d.Table.Plugin),
		"salesforce_product":                   SalesforceProduct(ctx, dynamicColumnsMap["Product2"], d.Table.Plugin),
		"salesforce_user":                      SalesforceUser(ctx, dynamicColumnsMap["User"], d.Table.Plugin),
	}

	var re = regexp.MustCompile(`\d+`)
	var substitution = ``
	salesforceTables := []string{}
	config := GetConfig(d.Connection)
	if config.Objects != nil && len(*config.Objects) > 0 {
		for _, tableName := range *config.Objects {
			pluginTableName := "salesforce_" + strcase.ToSnake(re.ReplaceAllString(tableName, substitution))
			if _, ok := tables[pluginTableName]; !ok {
				salesforceTables = append(salesforceTables, tableName)
			}
		}
	}

	if client == nil {
		plugin.Logger(ctx).Warn("salesforce.pluginTableDefinitions", "client_not_found: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
		return tables, nil
	}

	var wg sync.WaitGroup
	wg.Add(len(salesforceTables))
	for _, sfTable := range salesforceTables {
		tableName := "salesforce_" + strcase.ToSnake(re.ReplaceAllString(sfTable, substitution))
		if tables[tableName] != nil {
			wg.Done()
			continue
		}
		go func(name string) {
			defer wg.Done()
			tableName := "salesforce_" + strcase.ToSnake(re.ReplaceAllString(name, substitution))
			plugin.Logger(ctx).Debug("salesforce.pluginTableDefinitions", "object_name", name, "table_name", tableName)
			ctx = context.WithValue(ctx, contextKey("PluginTableName"), tableName)
			ctx = context.WithValue(ctx, contextKey("SalesforceTableName"), name)
			table := generateDynamicTables(ctx, d.Table.Plugin)
			// Ignore if the requested Salesforce object is not present.
			if table != nil {
				tables[tableName] = table
			}
		}(sfTable)
	}
	wg.Wait()
	return tables, nil
}

func generateDynamicTables(ctx context.Context, p *plugin.Plugin) *plugin.Table {

	client, err := connectRaw(ctx, d.ConnectionManager, d.Connection)
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
	salesforceCols := map[string]string{}
	// Key columns
	keyColumns := plugin.KeyColumnSlice{}

	salesforceObjectMetadata := *sObjectMeta
	salesforceObjectMetadataAsByte, err := json.Marshal(salesforceObjectMetadata["fields"])
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.generateDynamicTables", "json marshal error", err)
	}

	salesforceObjectFields := []map[string]interface{}{}
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

		if properties["soapType"] == nil {
			continue
		}
		soapType := strings.Split((properties["soapType"]).(string), ":")
		fieldType := soapType[len(soapType)-1]

		// Column dynamic generation
		// Don't convert to snake case since field names can have underscores in
		// them, so it's impossible to convert from snake case back to camel case
		// to match the original field name. Also, if we convert to snake case,
		// custom fields like "TestField" and "Test_Field" will result in duplicates
		var columnFieldName string
		if strings.HasSuffix(fieldName, "__c") {
			columnFieldName = strings.ToLower(fieldName)
		} else {
			columnFieldName = strcase.ToSnake(fieldName)
		}

		column := plugin.Column{
			Name:        columnFieldName,
			Description: fmt.Sprintf("%s.", properties["label"].(string)),
			Transform:   transform.FromP(getFieldFromSObjectMap, fieldName),
		}
		// Adding column type in the map to help in qual handling
		salesforceCols[columnFieldName] = fieldType

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

	Table := plugin.Table{
		Name:        tableName,
		Description: fmt.Sprintf("Represents Salesforce object %s.", salesforceObjectMetadata["name"]),
		List: &plugin.ListConfig{
			KeyColumns: keyColumns,
			Hydrate:    listSalesforceObjectsByTable(salesforceTableName, salesforceCols),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSalesforceObjectbyID(salesforceTableName),
		},
		Columns: cols,
	}

	return &Table
}
