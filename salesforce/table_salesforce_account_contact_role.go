package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceAccountContactRole(ctx context.Context, p *plugin.Plugin) *plugin.Table {
	cols, keyColumns := dynamicColumns(ctx, "AccountContactRole", p)
	return &plugin.Table{
		Name:        "salesforce_account_contact_role",
		Description: "Represents the role that a Contact plays on an Account.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable("AccountContactRole", cols),
			KeyColumns: keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID("AccountContactRole"),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, p, cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the account contact role in Salesforce."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "ID of the Account."},
			{Name: "contact_id", Type: proto.ColumnType_STRING, Description: "ID of the Contact associated with this account."},
			{Name: "is_primary", Type: proto.ColumnType_BOOL, Description: "Specifies whether the Contact plays the primary role on the Account (true) or not (false). Note that each account has only one primary contact role. Label is Primary. Default value is false."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Name of the role played by the Contact on this Account, such as Decision Maker, Approver, Buyer, and so on."},

			// Other columns
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created contact role record."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of the creation of the contact role record."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who most recently changed the contact role record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent change in the contact role record."},
		}),
	}
}
