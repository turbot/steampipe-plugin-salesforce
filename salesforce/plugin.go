package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-salesforce"

type contextKey string

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		SchemaMode:   plugin.SchemaModeDynamic,
		TableMapFunc: pluginTableDefinitions,
	}

	return p
}

func pluginTableDefinitions(ctx context.Context, p *plugin.Plugin) (map[string]*plugin.Table, error) {

	// Some fixed tables added to plugin
	salesforceTables := []string{
		"Account",
		"AccountContactRole",
		"Asset",
		"Campaign",
		"Case",
		"Contact",
		"Contract",
		"Lead",
		"Opportunity",
		"OpportunityContactRole",
		"Order",
		"Pricebook2",
		"PricebookEntry",
		"Product2",
		"User",
	}

	config := GetConfig(p.Connection)
	if config.Tables != nil && len(*config.Tables) > 0 {
		for _, tableName := range *config.Tables {
			if !helpers.StringSliceContains(salesforceTables, tableName) {
				salesforceTables = append(salesforceTables, tableName)
			}
		}
	}
	// Initialize tables
	tables := map[string]*plugin.Table{}

	var re = regexp.MustCompile(`\d+`)
	var substitution = ``

	for _, table := range salesforceTables {
		ctx = context.WithValue(ctx, contextKey("SalesforceTableName"), table)
		tableName := "salesforce_" + strcase.ToSnake(re.ReplaceAllString(table, substitution))
		ctx = context.WithValue(ctx, contextKey("PluginTableName"), tableName)
		plugin.Logger(ctx).Debug("pluginTableDefinitions Table Names", "SALESFORCE_NAME", table, "STEAMPIPE_NAME", tableName)
		if tables[tableName] == nil {
			tables[tableName] = generateDynamicTables(ctx, p)
		} else {
			plugin.Logger(ctx).Error("salesforce.pluginTableDefinitions", "table_already_exists", tableName)
		}
	}
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
		plugin.Logger(ctx).Error("salesforce.generateDynamicTables", fmt.Sprintf("Table %s not present in salesforce", salesforceTableName))
		return nil
	}

	// Top columns
	cols := []*plugin.Column{}

	data := *sObjectMeta
	data1, err := json.Marshal(data["fields"])
	if err != nil {
		plugin.Logger(ctx).Error("[simpleforce] json marshal error %v", err)
	}

	metadata := []map[string]interface{}{}
	var queryColumns []string
	err = json.Unmarshal(data1, &metadata)
	if err != nil {
		plugin.Logger(ctx).Error("[simpleforce] json unmarshal error %v", err)
	}

	for _, fields := range metadata {
		if fields["name"] == nil {
			continue
		}
		fieldName := fields["name"].(string)
		compoundFieldName := fields["compoundFieldName"]
		if compoundFieldName != nil && compoundFieldName.(string) != fieldName {
			continue
		}

		queryColumns = append(queryColumns, fieldName)
		if fields["soapType"] == nil {
			continue
		}
		soapType := strings.Split((fields["soapType"]).(string), ":")
		fieldType := soapType[len(soapType)-1]

		// Coloumn dynamic generation
		column := plugin.Column{
			Name:        strcase.ToSnake(fieldName),
			Description: fmt.Sprintf("The %s.", fields["label"].(string)),
			Transform:   transform.FromP(getFieldFromMap, fieldName),
		}

		// Set column type based on the `soapType` from salesforce schema
		switch fieldType {
		case "string", "ID":
			column.Type = proto.ColumnType_STRING
		case "date", "dateTime":
			column.Type = proto.ColumnType_TIMESTAMP
		case "boolean":
			column.Type = proto.ColumnType_BOOL
		case "double":
			column.Type = proto.ColumnType_DOUBLE
		case "int":
			column.Type = proto.ColumnType_INT
		default:
			column.Type = proto.ColumnType_JSON
		}
		cols = append(cols, &column)

	}

	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(queryColumns, ", "), salesforceTableName)

	return &plugin.Table{
		Name:        tableName,
		Description: fmt.Sprintf("Salesforce %s.", data["label"]),
		List: &plugin.ListConfig{
			// KeyColumns: keyColumns,
			Hydrate: listSalesforceResourceWithName(salesforceTableName, query),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSalesforceResourceWithName(salesforceTableName),
		},
		Columns: cols,
	}
}
