package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// LIST HYDRATE FUNCTION

func listSalesforceObjectsByTable(tableName string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceObjectsByTable", "connect error", err)
			return nil, err
		}

		query := generateQuery(d.QueryContext.Columns, tableName)
		plugin.Logger(ctx).Info("listSalesforceObjectsByTable", "query", query)

		for {
			result, err := client.Query(query)
			if err != nil {
				plugin.Logger(ctx).Error("listSalesforceObjectsByTable", "query error", err)
				return nil, err
			}

			AccountList := new([]map[string]interface{})
			err = decodeQueryResult(ctx, result.Records, AccountList)
			if err != nil {
				plugin.Logger(ctx).Error("listSalesforceObjectsByTable", "results decoding error", err)
				return nil, err
			}

			for _, account := range *AccountList {
				d.StreamListItem(ctx, account)
				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit
				if d.QueryStatus.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}

			// Paging
			if result.Done {
				break
			} else {
				query = result.NextRecordsURL
			}
		}

		return nil, nil
	}
}

//// GET HYDRATE FUNCTION

func getSalesforceObjectbyID(tableName string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		id := d.KeyColumnQualString("id")

		client, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("getSalesforceObjectbyID", "connect error", err)
			return nil, err
		}

		obj := client.SObject(tableName).Get(id)
		if obj == nil {
			// Object doesn't exist, handle the error
			plugin.Logger(ctx).Warn("getSalesforceObjectbyID", fmt.Sprintf("%s with Id \"%s\" not found", tableName, id))
			return nil, nil
		}

		opportunityContactRole := new(map[string]interface{})
		err = decodeQueryResult(ctx, obj, opportunityContactRole)
		if err != nil {
			plugin.Logger(ctx).Error("getSalesforceObjectbyID", "result decoding error", err)
			return nil, err
		}

		return *opportunityContactRole, nil
	}
}

//// TRANSFORM FUNCTION

func getFieldFromSObjectMap(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	param := d.Param.(string)
	ls := d.HydrateItem.(map[string]interface{})
	return ls[param], nil
}
