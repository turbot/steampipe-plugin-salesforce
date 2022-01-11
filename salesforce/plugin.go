package salesforce

// import (
// 	"context"

// 	"github.com/turbot/steampipe-plugin-sdk/plugin"
// 	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
// )

// const pluginName = "steampipe-plugin-salesforce"

// // Plugin creates this (salesforce) plugin
// func Plugin(ctx context.Context) *plugin.Plugin {
// 	p := &plugin.Plugin{
// 		Name:             pluginName,
// 		DefaultTransform: transform.FromCamel().NullIfZero(),
// 		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
// 			NewInstance: ConfigInstance,
// 			Schema:      ConfigSchema,
// 		},
// 		TableMap: map[string]*plugin.Table{
// 			"salesforce_account":                  SalesforceAccount(ctx),
// 			"salesforce_account_contact_role":     SalesforceAccountContactRole(ctx),
// 			"salesforce_contract":                 SalesforceContract(ctx),
// 			"salesforce_lead":                     SalesforceLead(ctx),
// 			"salesforce_opportunity":              SalesforceOpportunity(ctx),
// 			"salesforce_opportunity_contact_role": SalesforceOpportunityContactRole(ctx),
// 			"salesforce_order":                    SalesforceOrder(ctx),
// 			"salesforce_product":                  SalesforceProduct(ctx),
// 			"salesforce_user":                     SalesforceUser(ctx),
// 		},
// 	}

// 	return p
// }

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-salesforce"

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

	salesforceTables := []string{
		"Account",
		"AccountContactRole",
		"Contact",
		"Contract",
		"Lead",
		"Opportunity",
		"OpportunityContactRole",
		"Order",
		"Product2",
		"User",
	}

	// Initialize tables
	tables := map[string]*plugin.Table{}

	var re = regexp.MustCompile(`\d+`)
	var substitution = ``

	for _, table := range salesforceTables {
		ctx = context.WithValue(ctx, "SalesforceTableName", table)
		tableName := "salesforce_" + strcase.ToSnake(re.ReplaceAllString(table, substitution))
		ctx = context.WithValue(ctx, "PluginTableName", tableName)
		plugin.Logger(ctx).Info("pluginTableDefinitions Table Names", table, tableName)
		if tables[tableName] == nil {
			tables[tableName] = tableDynamicMetric(ctx, p)
		} else {
			plugin.Logger(ctx).Error("prometheus.pluginTableDefinitions", "table_already_exists", tableName)
		}
	}
	return tables, nil
}

func tableDynamicMetric(ctx context.Context, p *plugin.Plugin) *plugin.Table {

	client, err := connectRaw(ctx, p.ConnectionManager, p.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_dynamic_metric.tableDynamicMetric", "connection_error", err)
		return nil
	}

	// Get the query for the metric (required)
	salesforceTableName := ctx.Value("SalesforceTableName").(string)
	tableName := ctx.Value("PluginTableName").(string)
	plugin.Logger(ctx).Info("tableDynamicMetric Table Names", salesforceTableName, tableName)

	sObjectMeta := client.SObject(salesforceTableName).Describe()
	if sObjectMeta == nil {
		plugin.Logger(ctx).Error("prometheus_dynamic_metric.tableDynamicMetric", fmt.Sprintf("Table %s not present", salesforceTableName))
		return nil
	}

	// Top columns
	cols := []*plugin.Column{}
	// {Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Timestamp").Transform(transform.UnixMsToTimestamp), Description: "Timestamp of the value."},
	// {Name: "value", Type: proto.ColumnType_DOUBLE, Description: "Value of the metric."},
	// }

	// List of columns already seen
	// seenCols := map[string]bool{"__name__": true, "labels": true, "step_seconds": true, "timestamp": true, "value": true}

	// Key columns for query filtering
	// keyColumns := []*plugin.KeyColumn{
	// 	{Name: "step_seconds", Require: plugin.Optional},
	// 	{Name: "timestamp", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
	// }

	data := *sObjectMeta
	data1, err := json.Marshal(data["fields"])
	if err != nil {
		log.Fatalf("simpleforce json marshal error %v", err)
	}

	metadata := []map[string]interface{}{}
	var queryColumns []string
	err = json.Unmarshal(data1, &metadata)
	if err != nil {
		log.Fatalf("simpleforce json unmarshal error %v", err)
	}

	for _, fields := range metadata {
		if fields["compoundFieldName"] == nil {
			if fields["name"] == nil {
				continue
			}
			fieldName := fields["name"].(string)
			queryColumns = append(queryColumns, fieldName)
			if fields["soapType"] == nil {
				continue
			}
			soapType := strings.Split((fields["soapType"]).(string), ":")
			fieldType := soapType[len(soapType)-1]
			// column := plugin.Column{Name: strcase.ToSnake(fieldName), Description: fmt.Sprintf("The %s.", fields["label"].(string))}
			column := plugin.Column{Name: strcase.ToSnake(fieldName)}
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
	}

	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(queryColumns, ", "), salesforceTableName)
	context.WithValue(ctx, "Query", query)

	// Log the warnings
	// for _, i := range warnings {
	// 	plugin.Logger(ctx).Error("prometheus_dynamic_metric.tableDynamicMetric", "query_warning", i)
	// }

	// If there were no labels, then return the full set in each row
	// if len(cols) == 2 {
	// 	cols = append(cols, &plugin.Column{Name: "labels", Type: proto.ColumnType_JSON, Transform: transform.FromField("Metric"), Description: "Map of all labels in the metric."})
	// }

	// Other columns added at the end
	// cols = append(cols, &plugin.Column{Name: "step_seconds", Type: proto.ColumnType_INT, Transform: transform.FromQual("step_seconds").Transform(getStepSeconds), Description: "Interval in seconds between metric values. Default 60 seconds."})

	return &plugin.Table{
		Name: tableName,
		// Description: fmt.Sprintf("Salesforce %s.", data["label"]),
		List: &plugin.ListConfig{
			// KeyColumns: keyColumns,
			Hydrate: listMetricWithName(),
		},
		Columns: cols,
	}
}

func getMetricLabelFromMetric(_ context.Context, d *transform.TransformData) (interface{}, error) {
	// param := d.Param.(string)
	// ls := d.Value.(model.Metric)
	return "ls[model.LabelName(param)]", nil
}

func listMetricWithName() func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		return nil, nil
	}
}
