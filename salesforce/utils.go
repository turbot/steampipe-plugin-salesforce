package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/simpleforce/simpleforce"
	"github.com/turbot/steampipe-plugin-sdk/connection"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*simpleforce.Client, error) {
	return connectRaw(ctx, d.ConnectionManager, d.Connection)
}

// connect:: returns salesforce client after authentication
func connectRaw(_ context.Context, cm *connection.Manager, c *plugin.Connection) (*simpleforce.Client, error) {
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
		return nil, fmt.Errorf("'password' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if config.Token == nil {
		return nil, fmt.Errorf("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if config.User == nil {
		return nil, fmt.Errorf("'user' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// setup client
	client := simpleforce.NewClient(*config.URL, clientID, apiVersion)
	if client == nil {
		return nil, fmt.Errorf("couldn't get salesforce client. Client setup error.")
	}

	// login client
	err := client.LoginPassword(*config.User, *config.Password, *config.Token)
	if err != nil {
		return nil, fmt.Errorf("client logging error %v", err)
	}

	// Save to cache
	cm.Cache.Set(cacheKey, client)

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

	plugin.Logger(ctx).Info("decodeQueryResult", "Items", string(resp))
	err = json.Unmarshal(resp, respObject)
	if err != nil {
		return err
	}

	return nil
}
