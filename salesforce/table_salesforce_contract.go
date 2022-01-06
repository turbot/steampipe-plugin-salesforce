package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceContract(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_contract",
		Description: "Salesforce Contract",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceContract,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceContract,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "activated_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "activated_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "billing_address", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "company_signed_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "company_signed_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "contract_number", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "contract_term", Type: proto.ColumnType_INT, Description: ""},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customer_signed_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customer_signed_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customer_signed_title", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "end_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "last_activity_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_approved_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_referenced_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_viewed_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "owner_expiration_notice", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "special_terms", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "start_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "status", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "system_modstamp", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listSalesforceContract(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSalesforceContract", "connect error", err)
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Contract")
	plugin.Logger(ctx).Info("listSalesforceContract", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceContract", "query error", err)
			return nil, err
		}

		ContractList := new([]Contract)
		err = decodeQueryResult(ctx, result.Records, ContractList)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceContract", "decode results error", err)
			return nil, err
		}

		for _, contract := range *ContractList {
			d.StreamListItem(ctx, contract)
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

func getSalesforceContract(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	contractID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceContract", "connect error", err)
		return nil, err
	}

	obj := client.SObject("Contract").Get(contractID)
	if obj == nil {
		// Object doesn't exist, handle the error
		plugin.Logger(ctx).Warn("getSalesforceContract", fmt.Sprintf("contract \"%s\" not found", contractID))
		return nil, nil
	}

	account := new(Account)
	err = decodeQueryResult(ctx, obj, account)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceContract", "decode results error", err)
		return nil, err
	}

	return nil, nil
}
