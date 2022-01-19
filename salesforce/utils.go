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
		plugin.Logger(ctx).Warn("couldn't get salesforce client. Client setup error.")
		return nil, fmt.Errorf("couldn't get salesforce client. Client setup error.")
	}

	// login client
	err := client.LoginPassword(*config.User, *config.Password, *config.Token)
	if err != nil {
		plugin.Logger(ctx).Warn("client login error", err)
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
		// columnName := strcase.ToCamel(column)
		// // Salesforce custom fields are suffixed with '__c' after the custom field name.
		// if strings.HasSuffix(columnName, "C") {
		// 	columnName = columnName[0:len(columnName)-1] + "__c"
		// }
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

	return "s"
}

func getSalesforceColumnName(name string) string {
	columnName := strcase.ToCamel(name)
	// Salesforce custom fields are suffixed with '__c' after the custom field name.
	if strings.HasSuffix(columnName, "C") {
		columnName = columnName[0:len(columnName)-1] + "__c"
	}
	return columnName
}
