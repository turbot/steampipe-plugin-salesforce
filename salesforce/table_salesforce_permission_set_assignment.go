package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforcePermissionSetAssignment(ctx context.Context, dm dynamicMap, p *plugin.Plugin) *plugin.Table {
	tableName := "PermissionSetAssignment"
	return &plugin.Table{
		Name:        "salesforce_permission_set_assignment",
		Description: "Represents the association between a User and a PermissionSet.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, dm.salesforceColumns),
			KeyColumns: dm.keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, p, dm.cols, []*plugin.Column{
			{Name: "assignee_id", Type: proto.ColumnType_STRING, Description: "ID of the User to assign the permission set specified in PermissionSetId."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The Permission Set Assignment ID."},
			{Name: "permission_set_group_id", Type: proto.ColumnType_STRING, Description: "If associated with a permission set group, this is the ID of that group."},
			{Name: "permission_set_id", Type: proto.ColumnType_STRING, Description: "ID of the PermissionSet to assign to the user specified in AssigneeId."},
			{Name: "system_modstamp", Type: proto.ColumnType_TIMESTAMP, Description: "The Date Assigned."},
		}),
	}
}
