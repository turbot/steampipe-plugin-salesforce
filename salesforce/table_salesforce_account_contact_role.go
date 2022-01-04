package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceAccountContactRole(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_account_contact_role",
		Description: "Salesforce Account Contact Role",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceAccountContactRole,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceAccountContactRole,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "contact_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "role", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_primary", Type: proto.ColumnType_BOOL, Description: ""},
		},
	}
}

func listSalesforceAccountContactRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "AccountContactRole")
	plugin.Logger(ctx).Info("listSalesforceAccount", "Query", query)
	query = "SELECT Id, ContactId, AccountId, CreatedById, CreatedDate, Role, LastModifiedById, LastModifiedDate, IsPrimary FROM AccountContactRole"

	for {
		result, err := client.Query(query)
		if err != nil {
			return nil, err
		}

		AccountContactRoleList := new([]AccountContactRole)
		err = decodeQueryResult(ctx, result.Records, AccountContactRoleList)
		if err != nil {
			return nil, err
		}

		for _, account := range *AccountContactRoleList {
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

func getSalesforceAccountContactRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	accountContactRoleID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	obj := client.SObject("AccountContactRole").Get(accountContactRoleID)
	if obj == nil {
		// Object doesn't exist, handle the error
		return nil, nil
	}

	accountContactRole := new(AccountContactRole)
	err = decodeQueryResult(ctx, obj, accountContactRole)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
