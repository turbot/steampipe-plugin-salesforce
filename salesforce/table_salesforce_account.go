package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_account",
		Description: "Salesforce Account",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceAccount,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceAccount,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "owner_id", Type: proto.ColumnType_STRING},
			{Name: "created_by_id", Type: proto.ColumnType_STRING},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP},
			{Name: "industry", Type: proto.ColumnType_STRING},
			{Name: "annual_revenue", Type: proto.ColumnType_DOUBLE},
			{Name: "number_of_employees", Type: proto.ColumnType_DOUBLE},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP},
			{Name: "website", Type: proto.ColumnType_STRING},
			{Name: "phone", Type: proto.ColumnType_STRING},
			{Name: "account_source", Type: proto.ColumnType_STRING},
			{Name: "type", Type: proto.ColumnType_STRING},
			{Name: "billing_address", Type: proto.ColumnType_JSON},
			{Name: "shipping_address", Type: proto.ColumnType_JSON},
		},
	}
}

func listSalesforceAccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSalesforceAccount", "connect error", err)
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Account")
	plugin.Logger(ctx).Info("listSalesforceAccount", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceAccount", "query error", err)
			return nil, err
		}

		AccountList := new([]Account)
		err = decodeQueryResult(ctx, result.Records, AccountList)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceAccount", "decode results error", err)
			return nil, err
		}

		for _, account := range *AccountList {
			d.StreamListItem(ctx, account)
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

func getSalesforceAccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	accountID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceAccount", "connect error", err)
		return nil, err
	}

	obj := client.SObject("Account").Get(accountID)
	if obj == nil {
		// Object doesn't exist, handle the error
		plugin.Logger(ctx).Warn("getSalesforceAccount", fmt.Sprintf("account \"%s\" not found", accountID))
		return nil, nil
	}

	account := new(Account)
	err = decodeQueryResult(ctx, obj, account)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceAccount", "decode results error", err)
		return nil, err
	}

	return nil, nil
}
