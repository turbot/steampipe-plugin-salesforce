package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/simpleforce/simpleforce"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// connect:: returns salesforce client after authentication
func connect(_ context.Context, d *plugin.QueryData) (*simpleforce.Client, error) {
	config := GetConfig(d.Connection)

	client := simpleforce.NewClient(*config.URL, simpleforce.DefaultClientID, simpleforce.DefaultAPIVersion)
	if client == nil {
		return nil, fmt.Errorf("couldn't get salesforce client. Client generation error.")
	}
	err := client.LoginPassword(*config.User, *config.Password, *config.Token)
	if err != nil {
		return nil, fmt.Errorf("client signing error %v", err)
	}

	return client, nil
}

// generateQuery:: returns sql query based on the column names, table name passed
func generateQuery(columns []string, tableName string) string {
	var queryColumns []string
	for _, column := range columns {
		queryColumns = append(queryColumns, strcase.ToCamel(column))
	}

	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(queryColumns, ", "), tableName)
}

// decodeQueryResult(ctx, apiResponse, responseStruct):: converts raw apiResponse to required output struct
func decodeQueryResult(ctx context.Context, response interface{}, respObject interface{}) error {
	resp, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// plugin.Logger(ctx).Info("decodeQueryResult", "Items", string(resp))
	err = json.Unmarshal(resp, respObject)
	if err != nil {
		return err
	}

	return nil
}
