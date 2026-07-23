package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// RemoteProxy (Remote Site Settings) is only queryable through the Tooling API.
func SalesforceRemoteSiteSetting(ctx context.Context, config salesforceConfig) *plugin.Table {
	tableName := "RemoteProxy"
	return &plugin.Table{
		Name:        "salesforce_remote_site_setting",
		Description: "Represents a remote site setting, which authorizes Salesforce to make callouts to an external site. Queried through the Salesforce Tooling API.",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceToolingObjectsByTable(tableName),
		},
		Columns: []*plugin.Column{
			{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the organization in Salesforce.", Hydrate: getOrganizationId, Transform: transform.FromValue()},

			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the remote site setting."},
			{Name: "site_name", Type: proto.ColumnType_STRING, Description: "The unique name of the remote site setting in the API."},
			{Name: "endpoint_url", Type: proto.ColumnType_STRING, Description: "The URL of the remote site that Salesforce is authorized to call."},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: "If true, the remote site setting is active."},

			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the remote site setting."},
			{Name: "protocol_mismatch", Type: proto.ColumnType_BOOL, Description: "If true, Salesforce allows callouts to this site over an unsecured protocol (HTTP)."},
			{Name: "namespace_prefix", Type: proto.ColumnType_STRING, Description: "The namespace prefix if the remote site setting is part of a managed package."},
		},
	}
}
