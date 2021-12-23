package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

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
