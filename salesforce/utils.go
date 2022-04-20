package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/simpleforce/simpleforce"
	"github.com/turbot/steampipe-plugin-sdk/v3/connection"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
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
	apiVersion := simpleforce.DefaultAPIVersion
	clientID := "steampipe"
	securityToken := ""

	if config.ClientId != nil {
		clientID = *config.ClientId
	}

	if config.APIVersion != nil {
		apiVersion = *config.APIVersion
	}

	if config.Username == nil {
		plugin.Logger(ctx).Warn("salesforce.connectRaw", "'username' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
		return nil, nil
	}

	if config.Password == nil {
		plugin.Logger(ctx).Warn("salesforce.connectRaw", "'password' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
		return nil, nil
	}

	// The Salesforce security token is only required If the client's IP address is not added to the organization's list of trusted IPs
	// https://help.salesforce.com/s/articleView?id=sf.security_networkaccess.htm&type=5
	// https://migration.trujay.com/help/how-to-add-an-ip-address-to-whitelist-on-salesforce/
	if config.Token != nil {
		securityToken = *config.Token
	}

	// setup client
	client := simpleforce.NewClient(*config.URL, clientID, apiVersion)
	if client == nil {
		plugin.Logger(ctx).Error("salesforce.connectRaw", "couldn't get salesforce client. Client setup error.")
		return nil, fmt.Errorf("salesforce.connectRaw couldn't get salesforce client. Client setup error.")
	}

	// LoginPassword signs into salesforce using password. token is optional if trusted IP is configured.
	// Ref: https://developer.salesforce.com/docs/atlas.en-us.214.0.api_rest.meta/api_rest/intro_understanding_username_password_oauth_flow.htm
	// Ref: https://developer.salesforce.com/docs/atlas.en-us.214.0.api.meta/api/sforce_api_calls_login.htm
	err := client.LoginPassword(*config.Username, *config.Password, securityToken)
	if err != nil {
		plugin.Logger(ctx).Error("salesforce.connectRaw", "client login error", err)
		return nil, fmt.Errorf("client login error %v", err)
	}

	// Save to cache
	cm.Cache.Set(cacheKey, client)

	return client, nil
}

// generateQuery:: returns sql query based on the column names, table name passed
func generateQuery(columns []*plugin.Column, tableName string) string {
	var queryColumns []string
	for _, column := range columns {
		queryColumns = append(queryColumns, getSalesforceColumnName(column.Name))
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
func buildQueryFromQuals(equalQuals plugin.KeyColumnQualMap, tableColumns []*plugin.Column, salesforceCols map[string]string) string {
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
						default:
							filters = append(filters, fmt.Sprintf("%s %s %d", getSalesforceColumnName(filterQualItem.Name), qual.Operator, value.GetInt64Value()))
						}
					case proto.ColumnType_DOUBLE:
						switch qual.Operator {
						case "<>":
							filters = append(filters, fmt.Sprintf("%s != %f", getSalesforceColumnName(filterQualItem.Name), value.GetDoubleValue()))
						default:
							filters = append(filters, fmt.Sprintf("%s %s %f", getSalesforceColumnName(filterQualItem.Name), qual.Operator, value.GetDoubleValue()))
						}
					// Need a way to distinguish b/w date and dateTime fields
					case proto.ColumnType_TIMESTAMP:
						// https://developer.salesforce.com/docs/atlas.en-us.234.0.soql_sosl.meta/soql_sosl/sforce_api_calls_soql_select_dateformats.htm
						if salesforceCols[filterQual.Name] == "date" {
							switch qual.Operator {
							case "=", ">=", ">", "<=", "<":
								filters = append(filters, fmt.Sprintf("%s %s %s", getSalesforceColumnName(filterQualItem.Name), qual.Operator, value.GetTimestampValue().AsTime().Format("2006-01-02")))
							}
						} else {
							switch qual.Operator {
							case "=", ">=", ">", "<=", "<":
								filters = append(filters, fmt.Sprintf("%s %s %s", getSalesforceColumnName(filterQualItem.Name), qual.Operator, value.GetTimestampValue().AsTime().Format("2006-01-02T15:04:05Z")))
							}
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
	var columnName string
	// Salesforce custom fields are suffixed with '__c' and are not converted to
	// snake case in the table schema, so use the column name as is
	if strings.HasSuffix(name, "__c") {
		columnName = name
	} else {
		columnName = strcase.ToCamel(name)
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
func dynamicColumns(ctx context.Context, client *simpleforce.Client, salesforceTableName string, p *plugin.Plugin) ([]*plugin.Column, plugin.KeyColumnSlice, map[string]string) {
	sObjectMeta := client.SObject(salesforceTableName).Describe()
	if sObjectMeta == nil {
		plugin.Logger(ctx).Error("salesforce.dynamicColumns", fmt.Sprintf("Table %s not present in salesforce", salesforceTableName))
		return []*plugin.Column{}, plugin.KeyColumnSlice{}, map[string]string{}
	}

	// Top columns
	cols := []*plugin.Column{}
	salesforceCols := map[string]string{}
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
			Description: fmt.Sprintf("%s.", fields["label"].(string)),
			Transform:   transform.FromP(getFieldFromSObjectMap, fieldName),
		}
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

	return cols, keyColumns, salesforceCols
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
