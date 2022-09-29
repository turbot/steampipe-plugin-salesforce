package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func SalesforceAccountContactRole(ctx context.Context, dm dynamicMap) *plugin.Table {
	tableName := "AccountContactRole"
	return &plugin.Table{
		Name:        "salesforce_account_contact_role",
		Description: "Represents the role that a Contact plays on an Account.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, dm.salesforceColumns),
			KeyColumns: dm.keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, dm.cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the account contact role in Salesforce."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "ID of the Account."},
			{Name: "contact_id", Type: proto.ColumnType_STRING, Description: "ID of the Contact associated with this account."},
			{Name: "is_primary", Type: proto.ColumnType_BOOL, Description: "Specifies whether the Contact plays the primary role on the Account (true) or not (false). Note that each account has only one primary contact role."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Name of the role played by the Contact on this Account, such as Decision Maker, Approver, Buyer, and so on."},

			// Other columns
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created contact role record."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of the creation of the contact role record."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who most recently changed the contact role record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent change in the contact role record."},
		}),
	}
}
