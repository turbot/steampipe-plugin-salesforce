package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
	"github.com/simpleforce/simpleforce"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-salesforce"

type contextKey string

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.From(getFieldFromSObjectMapByColumnName).NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
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

func pluginTableDefinitions(ctx context.Context, td *plugin.TableMapData) (map[string]*plugin.Table, error) {
	// If unable to connect to salesforce instance, log warning and abort dynamic table creation

	client, err := connectRaw(ctx, td.ConnectionCache, td.Connection)
	if err != nil {
		// do not abort the plugin as static table needs to be generated
		plugin.Logger(ctx).Warn("salesforce.pluginTableDefinitions", "connection_error: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
	}

	staticTables := []string{"Account", "AccountContactRole", "Asset", "Contact", "Contract", "Lead", "Opportunity", "OpportunityContactRole", "Order", "Pricebook2", "Product2", "User", "PermissionSet", "PermissionSetAssignment", "ObjectPermissions"}

	dynamicColumnsMap := map[string]dynamicMap{}
	var mapLock sync.Mutex
	config := GetConfig(td.Connection)

	// If Salesforce client was obtained, don't generate dynamic columns for
	// defined static tables
	if client != nil {
		var wgd sync.WaitGroup
		wgd.Add(len(staticTables))
		for _, st := range staticTables {
			go func(staticTable string) {
				defer wgd.Done()
				dynamicCols, dynamicKeyColumns, salesforceCols := dynamicColumns(ctx, client, staticTable, config)
				mapLock.Lock()
				dynamicColumnsMap[staticTable] = dynamicMap{dynamicCols, dynamicKeyColumns, salesforceCols}
				defer mapLock.Unlock()
			}(st)
		}
		wgd.Wait()
	}

	// Initialize tables with static tables with static and dynamic columns(if credentials are set)
	var tables map[string]*plugin.Table

	// check the NamingConvention parameter value in config
	if config.NamingConvention != nil && *config.NamingConvention == "api_native" {
		tables = map[string]*plugin.Table{
			"Account":                 SalesforceAccount(ctx, dynamicColumnsMap["Account"], config),
			"AccountContactRole":      SalesforceAccountContactRole(ctx, dynamicColumnsMap["AccountContactRole"], config),
			"Asset":                   SalesforceAsset(ctx, dynamicColumnsMap["Asset"], config),
			"Contact":                 SalesforceContact(ctx, dynamicColumnsMap["Contact"], config),
			"Contract":                SalesforceContract(ctx, dynamicColumnsMap["Contract"], config),
			"Lead":                    SalesforceLead(ctx, dynamicColumnsMap["Lead"], config),
			"ObjectPermissions":       SalesforceObjectPermission(ctx, dynamicColumnsMap["ObjectPermissions"], config),
			"Opportunity":             SalesforceOpportunity(ctx, dynamicColumnsMap["Opportunity"], config),
			"OpportunityContactRole":  SalesforceOpportunityContactRole(ctx, dynamicColumnsMap["OpportunityContactRole"], config),
			"Order":                   SalesforceOrder(ctx, dynamicColumnsMap["Order"], config),
			"PermissionSet":           SalesforcePermissionSet(ctx, dynamicColumnsMap["PermissionSet"], config),
			"PermissionSetAssignment": SalesforcePermissionSetAssignment(ctx, dynamicColumnsMap["PermissionSetAssignment"], config),
			"Pricebook2":              SalesforcePricebook(ctx, dynamicColumnsMap["Pricebook2"], config),
			"Product2":                SalesforceProduct(ctx, dynamicColumnsMap["Product2"], config),
			"User":                    SalesforceUser(ctx, dynamicColumnsMap["User"], config),
		}
	} else {
		tables = map[string]*plugin.Table{
			"salesforce_account":                   SalesforceAccount(ctx, dynamicColumnsMap["Account"], config),
			"salesforce_account_contact_role":      SalesforceAccountContactRole(ctx, dynamicColumnsMap["AccountContactRole"], config),
			"salesforce_asset":                     SalesforceAsset(ctx, dynamicColumnsMap["Asset"], config),
			"salesforce_contact":                   SalesforceContact(ctx, dynamicColumnsMap["Contact"], config),
			"salesforce_contract":                  SalesforceContract(ctx, dynamicColumnsMap["Contract"], config),
			"salesforce_lead":                      SalesforceLead(ctx, dynamicColumnsMap["Lead"], config),
			"salesforce_object_permission":         SalesforceObjectPermission(ctx, dynamicColumnsMap["ObjectPermissions"], config),
			"salesforce_opportunity":               SalesforceOpportunity(ctx, dynamicColumnsMap["Opportunity"], config),
			"salesforce_opportunity_contact_role":  SalesforceOpportunityContactRole(ctx, dynamicColumnsMap["OpportunityContactRole"], config),
			"salesforce_order":                     SalesforceOrder(ctx, dynamicColumnsMap["Order"], config),
			"salesforce_permission_set":            SalesforcePermissionSet(ctx, dynamicColumnsMap["PermissionSet"], config),
			"salesforce_permission_set_assignment": SalesforcePermissionSetAssignment(ctx, dynamicColumnsMap["PermissionSetAssignment"], config),
			"salesforce_pricebook":                 SalesforcePricebook(ctx, dynamicColumnsMap["Pricebook2"], config),
			"salesforce_product":                   SalesforceProduct(ctx, dynamicColumnsMap["Product2"], config),
			"salesforce_user":                      SalesforceUser(ctx, dynamicColumnsMap["User"], config),
		}
	}

	var re = regexp.MustCompile(`\d+`)
	var substitution = ``
	salesforceTables := []string{}
	if config.Objects != nil && len(*config.Objects) > 0 {
		for _, tableName := range *config.Objects {
			var pluginTableName string
			if config.NamingConvention != nil && *config.NamingConvention == "api_native" {
				pluginTableName = strcase.ToSnake(re.ReplaceAllString(tableName, substitution))
			} else {
				pluginTableName = "salesforce_" + strcase.ToSnake(re.ReplaceAllString(tableName, substitution))
			}
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
		var tableName string
		if config.NamingConvention != nil && *config.NamingConvention == "api_native" {
			tableName = sfTable
		} else {
			tableName = "salesforce_" + strcase.ToSnake(re.ReplaceAllString(sfTable, substitution))
		}
		if tables[tableName] != nil {
			wg.Done()
			continue
		}
		go func(name string) {
			defer wg.Done()
			if config.NamingConvention != nil && *config.NamingConvention == "api_native" {
				tableName = name
			} else {
				tableName = "salesforce_" + strcase.ToSnake(re.ReplaceAllString(name, substitution))
			}
			plugin.Logger(ctx).Debug("salesforce.pluginTableDefinitions", "object_name", name, "table_name", tableName)
			ctx = context.WithValue(ctx, contextKey("PluginTableName"), tableName)
			ctx = context.WithValue(ctx, contextKey("SalesforceTableName"), name)
			table := generateDynamicTables(ctx, client, config)
			// Ignore if the requested Salesforce object is not present.
			if table != nil {
				tables[tableName] = table
			}
		}(sfTable)
	}
	wg.Wait()
	return tables, nil
}

func generateDynamicTables(ctx context.Context, client *simpleforce.Client, config salesforceConfig) *plugin.Table {
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

		// keep the field name as it is if NamingConvention is set to api_native
		if config.NamingConvention != nil && *config.NamingConvention == "api_native" {
			columnFieldName = fieldName
		} else if strings.HasSuffix(fieldName, "__c") {
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
			KeyColumns: plugin.SingleColumn(checkNameScheme(config, cols)),
			Hydrate:    getSalesforceObjectbyID(salesforceTableName),
		},
		Columns: cols,
	}
	return &Table
}

// set GetConfig parameter based on NamingConvention value
// if the object is unavailable then there will be no dynamic columns, so GetConfig parameter should be id to avoid failure of static table creation
func checkNameScheme(config salesforceConfig, dynamicColumns []*plugin.Column) string {
	if config.NamingConvention != nil && *config.NamingConvention == "api_native" && len(dynamicColumns) > 0 {
		return "Id"
	}
	return "id"
}
