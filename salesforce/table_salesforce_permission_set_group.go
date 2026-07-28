package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func SalesforcePermissionSetGroup(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "PermissionSetGroup"
	return &plugin.Table{
		Name:        "salesforce_permission_set_group",
		Description: "Represents a group of permission sets. Assigning a group to a user grants the combined permissions of all permission sets in the group.",
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the permission set group."},
			{Name: "developer_name", Type: proto.ColumnType_STRING, Description: "The unique name of the permission set group in the API."},
			{Name: "master_label", Type: proto.ColumnType_STRING, Description: "The permission set group label, displayed in the user interface."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The calculation status of the permission set group, for example, Updated or Outdated."},

			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the permission set group."},
			{Name: "has_activation_required", Type: proto.ColumnType_BOOL, Description: "If true, the permission set group requires an associated active session."},
			{Name: "namespace_prefix", Type: proto.ColumnType_STRING, Description: "The namespace prefix if the permission set group is part of a managed package."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who created the permission set group."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the permission set group was created."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who last modified the permission set group."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the permission set group was last modified."},
		}),
	}
}
