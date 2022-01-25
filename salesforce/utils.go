package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/simpleforce/simpleforce"
	"github.com/turbot/steampipe-plugin-sdk/connection"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*simpleforce.Client, error) {
	return connectRaw(ctx, d.ConnectionManager, d.Connection)
}

// connect:: returns salesforce client after authentication
func connectRaw(ctx context.Context, cm *connection.Manager, c *plugin.Connection) (*simpleforce.Client, error) {
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "simpleforce"
	if cachedData, ok := cm.Cache.Get(cacheKey); ok {
		return cachedData.(*simpleforce.Client), nil
	}

	config := GetConfig(c)
	clientID := simpleforce.DefaultClientID
	apiVersion := simpleforce.DefaultAPIVersion

	if config.ClientId != nil {
		clientID = *config.ClientId
	}

	if config.APIVersion != nil {
		apiVersion = *config.APIVersion
	}

	if config.Password == nil {
		plugin.Logger(ctx).Warn("'password' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
		return nil, nil
	}

	if config.Token == nil {
		plugin.Logger(ctx).Warn("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
		return nil, nil
	}

	if config.User == nil {
		plugin.Logger(ctx).Warn("'user' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
		return nil, nil
	}

	// setup client
	client := simpleforce.NewClient(*config.URL, clientID, apiVersion)
	if client == nil {
		plugin.Logger(ctx).Warn("salesforce.connectRaw", "couldn't get salesforce client. Client setup error.")
		return nil, fmt.Errorf("salesforce.connectRaw couldn't get salesforce client. Client setup error.")
	}

	// login client
	err := client.LoginPassword(*config.User, *config.Password, *config.Token)
	if err != nil {
		plugin.Logger(ctx).Warn("salesforce.connectRaw", "client login error", err)
		return nil, fmt.Errorf("client login error %v", err)
	}

	// Save to cache
	cm.Cache.Set(cacheKey, client)

	return client, nil
}

// generateQuery:: returns sql query based on the column names, table name passed
func generateQuery(columns []string, tableName string) string {
	var queryColumns []string
	for _, column := range columns {
		queryColumns = append(queryColumns, getSalesforceColumnName(column))
	}

	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(queryColumns, ", "), tableName)
}

// decodeQueryResult(ctx, apiResponse, responseStruct):: converts raw apiResponse to required output struct
func decodeQueryResult(ctx context.Context, response interface{}, respObject interface{}) error {
	resp, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// For debugging purpose - commenting out to avoid unnecessary logs
	// plugin.Logger(ctx).Info("decodeQueryResult", "Items", string(resp))
	err = json.Unmarshal(resp, respObject)
	if err != nil {
		return err
	}

	return nil
}

// buildQueryFromQuals :: generate SOQL based on the contions specified in sql query
// refrences
// - https://developer.salesforce.com/docs/atlas.en-us.234.0.soql_sosl.meta/soql_sosl/sforce_api_calls_soql_select_comparisonoperators.htm
func buildQueryFromQuals(equalQuals plugin.KeyColumnQualMap, tableColumns []*plugin.Column) string {
	filters := []string{}

	for _, filterQualItem := range tableColumns {
		filterQual := equalQuals[filterQualItem.Name]
		if filterQual == nil {
			continue
		}

		// Check only if filter qual map matches with optional column name
		if filterQual.Name == filterQualItem.Name {
			if filterQual.Quals == nil {
				continue
			}

			for _, qual := range filterQual.Quals {
				if qual.Value != nil {
					value := qual.Value
					switch filterQualItem.Type {
					case proto.ColumnType_STRING:
						// In case of IN caluse
						if value.GetListValue() != nil {
							// continue
							switch qual.Operator {
							case "=":
								stringValueSlice := []string{}
								for _, q := range value.GetListValue().Values {
									stringValueSlice = append(stringValueSlice, q.GetStringValue())
								}
								if len(stringValueSlice) > 0 {
									filters = append(filters, fmt.Sprintf("%s IN ('%s')", getSalesforceColumnName(filterQualItem.Name), strings.Join(stringValueSlice, "','")))
								}
							case "<>":
								stringValueSlice := []string{}
								for _, q := range value.GetListValue().Values {
									stringValueSlice = append(stringValueSlice, q.GetStringValue())
								}
								if len(stringValueSlice) > 0 {
									filters = append(filters, fmt.Sprintf("%s NOT IN ('%s')", getSalesforceColumnName(filterQualItem.Name), strings.Join(stringValueSlice, "','")))
								}
							}
						} else {
							switch qual.Operator {
							case "=":
								filters = append(filters, fmt.Sprintf("%s = '%s'", getSalesforceColumnName(filterQualItem.Name), value.GetStringValue()))
							case "<>":
								filters = append(filters, fmt.Sprintf("%s != '%s'", getSalesforceColumnName(filterQualItem.Name), value.GetStringValue()))
							}
						}
					case proto.ColumnType_BOOL:
						switch qual.Operator {
						case "<>":
							filters = append(filters, fmt.Sprintf("%s = %s", getSalesforceColumnName(filterQualItem.Name), "FALSE"))
						case "=":
							filters = append(filters, fmt.Sprintf("%s = %s", getSalesforceColumnName(filterQualItem.Name), "TRUE"))
						}
					case proto.ColumnType_INT:
						switch qual.Operator {
						case "<>":
							filters = append(filters, fmt.Sprintf("%s != %d", getSalesforceColumnName(filterQualItem.Name), value.GetInt64Value()))
						case "=":
							filters = append(filters, fmt.Sprintf("%s %s %d", getSalesforceColumnName(filterQualItem.Name), qual.Operator, value.GetInt64Value()))
						}
					case proto.ColumnType_DOUBLE:
						switch qual.Operator {
						case "<>":
							filters = append(filters, fmt.Sprintf("%s != %f", getSalesforceColumnName(filterQualItem.Name), value.GetDoubleValue()))
						case "=":
							filters = append(filters, fmt.Sprintf("%s %s %f", getSalesforceColumnName(filterQualItem.Name), qual.Operator, value.GetDoubleValue()))
						}
					}
				}
			}

		}
	}

	if len(filters) > 0 {
		return strings.Join(filters, " AND ")
	}

	return ""
}

func getSalesforceColumnName(name string) string {
	columnName := strcase.ToCamel(name)
	// Salesforce custom fields are suffixed with '__c' after the custom field name.
	if strings.HasSuffix(columnName, "C") {
		columnName = columnName[0:len(columnName)-1] + "__c"
	}
	return columnName
}

// append the dynamic columns with static columns for the table
func mergeTableColumns(ctx context.Context, p *plugin.Plugin, dynamicColumns []*plugin.Column, staticColumns []*plugin.Column) []*plugin.Column {
	var columns []*plugin.Column
	columns = append(columns, staticColumns...)
	for _, col := range dynamicColumns {
		if isColumnAvailable(col.Name, staticColumns) {
			continue
		}
		columns = append(columns, col)
	}
	return columns
}

// dynamicColumns:: Returns list coulms for a salesforce object
func dynamicColumns(ctx context.Context, salesforceTableName string, p *plugin.Plugin) ([]*plugin.Column, plugin.KeyColumnSlice) {
	// If unable to connect to salesforce instance, log warning and abort dynamic table creation
	client, err := connectRaw(ctx, p.ConnectionManager, p.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.dynamicColumns", "connection_error: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
		return []*plugin.Column{}, plugin.KeyColumnSlice{}
	}
	if client == nil {
		plugin.Logger(ctx).Error("salesforce.dynamicColumns", "connection_error: unable to generate dynamic tables because of invalid steampipe salesforce configuration", err)
		return []*plugin.Column{}, plugin.KeyColumnSlice{}
	}

	sObjectMeta := client.SObject(salesforceTableName).Describe()
	if sObjectMeta == nil {
		plugin.Logger(ctx).Error("salesforce.dynamicColumns", fmt.Sprintf("Table %s not present in salesforce", salesforceTableName))
		return []*plugin.Column{}, plugin.KeyColumnSlice{}
	}

	// Top columns
	cols := []*plugin.Column{}
	// Key columns
	keyColumns := plugin.KeyColumnSlice{}

	salesforceObjectMetadata := *sObjectMeta
	salesforceObjectMetadataAsByte, err := json.Marshal(salesforceObjectMetadata["fields"])
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.dynamicColumns", "json marshal error %v", err)
	}

	salesforceObjectFields := []map[string]interface{}{}
	// var queryColumns []string
	err = json.Unmarshal(salesforceObjectMetadataAsByte, &salesforceObjectFields)
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.dynamicColumns", "json unmarshal error %v", err)
	}
	for _, fields := range salesforceObjectFields {
		if fields["name"] == nil {
			continue
		}
		fieldName := fields["name"].(string)
		compoundFieldName := fields["compoundFieldName"]
		if compoundFieldName != nil && compoundFieldName.(string) != fieldName {
			continue
		}

		if fields["soapType"] == nil {
			continue
		}
		soapType := strings.Split((fields["soapType"]).(string), ":")
		fieldType := soapType[len(soapType)-1]

		// Coloumn dynamic generation
		columnFieldName := strcase.ToSnake(fieldName)
		column := plugin.Column{
			Name:        columnFieldName,
			Description: fmt.Sprintf("The %s.", fields["label"].(string)),
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

	return cols, keyColumns
}

// isColumnAvailable:: Checks if the column is not present in the existing columns slice
func isColumnAvailable(columnName string, columns []*plugin.Column) bool {
	for _, col := range columns {
		if col.Name == columnName {
			return true
		}
	}
	return false
}
