package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforceSetupAuditTrail(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "SetupAuditTrail"
	return &plugin.Table{
		Name:        "salesforce_setup_audit_trail",
		Description: "Represents changes users made to the organization's configuration in the Setup area of Salesforce.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, dm.salesforceColumns),
			KeyColumns: dm.keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn(checkNameScheme(config, dm.cols)),
		},
		Columns: mergeTableColumns(ctx, config, dm.cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the setup audit trail entry."},
			{Name: "action", Type: proto.ColumnType_STRING, Description: "The action performed, describing the configuration change made in Setup."},
			{Name: "section", Type: proto.ColumnType_STRING, Description: "The Setup section in which the change was made."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the change was made."},

			// Other columns
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who made the change."},
			{Name: "display", Type: proto.ColumnType_STRING, Description: "The full description of the change that was made."},
			{Name: "delegate_user", Type: proto.ColumnType_STRING, Description: "The login of the delegate user who made the change, if the change was made on behalf of another user."},
		}),
	}
}
