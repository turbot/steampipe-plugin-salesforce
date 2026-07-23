package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforceProfile(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "Profile"
	return &plugin.Table{
		Name:        "salesforce_profile",
		Description: "Represents a profile, which defines a set of permissions and access settings that control what users can do in the organization.",
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the profile."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the profile."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the profile."},
			{Name: "user_type", Type: proto.ColumnType_STRING, Description: "The category of user license associated with the profile, for example, Standard or PowerPartner."},

			// Other columns
			{Name: "user_license_id", Type: proto.ColumnType_STRING, Description: "The id of the user license associated with the profile."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who created the profile."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the profile was created."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who last modified the profile."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the profile was last modified."},
		}),
	}
}
