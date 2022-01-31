package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforcePermissionSet(ctx context.Context, dm dynamicMap, p *plugin.Plugin) *plugin.Table {
	tableName := "PermissionSet"
	return &plugin.Table{
		Name:        "salesforce_permission_set",
		Description: "Represents a set of permissions that's used to grant more access to one or more users without changing their profile or reassigning profiles.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName),
			KeyColumns: dm.keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, p, dm.cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the permission set."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The permission set unique name in the API."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the permission set."},
			{Name: "is_custom", Type: proto.ColumnType_BOOL, Description: "If true, the permission set is custom (created by an admin); if false, the permission set is standard and related to a specific permission set license."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The Created Date."},

			// Other columns
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The contact id of the user who created this permission set."},
			{Name: "has_activation_required", Type: proto.ColumnType_BOOL, Description: "Indicates whether the permission set requires an associated active session (true) or not (false)."},
			{Name: "is_owned_by_profile", Type: proto.ColumnType_BOOL, Description: "If true, the permission set is owned by a profile."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "The permission set label, which corresponds to Label in the user interface."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The Last Modified By ID."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The Last Modified Date."},
			{Name: "license_id", Type: proto.ColumnType_STRING, Description: "The ID of either the related PermissionSetLicense or UserLicense associated with this permission set."},
			{Name: "namespace_prefix", Type: proto.ColumnType_STRING, Description: "The namespace prefix for a permission set that's been installed as part of a managed package. If the permission set isn't packaged or is part of an unmanaged package, this value is empty."},
			{Name: "permission_set_group_id", Type: proto.ColumnType_STRING, Description: "If the permission set is owned by a permission set group, this field returns the ID of the permission set group."},
			{Name: "profile_id", Type: proto.ColumnType_STRING, Description: "If the permission set is owned by a profile, this field contains the ID of the Profile."},
			{Name: "system_modstamp", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when order record was last modified by a user or by an automated process."},
		}),
	}
}
