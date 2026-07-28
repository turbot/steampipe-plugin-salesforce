package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

// NamedCredential is only queryable through the Tooling API. Columns are declared
// in snake case; toolingColumns renames them to the API field names in api_native mode.
func SalesforceNamedCredential(ctx context.Context, config salesforceConfig) *plugin.Table {
	tableName := "NamedCredential"
	return &plugin.Table{
		Name:        "salesforce_named_credential",
		Description: "Represents a named credential, which specifies the URL of a callout endpoint and its authentication settings for integrations. Queried through the Salesforce Tooling API.",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceToolingObjectsByTable(tableName),
		},
		Columns: toolingColumns(config, []*plugin.Column{
			{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the organization in Salesforce.", Hydrate: getOrganizationId, Transform: transform.FromValue()},

			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the named credential."},
			{Name: "developer_name", Type: proto.ColumnType_STRING, Description: "The unique name of the named credential in the API."},
			{Name: "master_label", Type: proto.ColumnType_STRING, Description: "The named credential label, displayed in the user interface."},
			{Name: "endpoint", Type: proto.ColumnType_STRING, Description: "The URL of the callout endpoint the named credential points to."},

			// Other columns
			// Note: secret fields (Password, OauthToken, OauthRefreshToken,
			// AwsAccessSecret, etc.) are intentionally not exposed as columns.
			{Name: "principal_type", Type: proto.ColumnType_STRING, Description: "The type of authentication principal (Anonymous, PerUser, or NamedUser)."},
			{Name: "named_credential_type", Type: proto.ColumnType_STRING, Description: "The type of named credential, for example, SecuredEndpoint or Legacy."},
			{Name: "protocol", Type: proto.ColumnType_STRING, Description: "The authentication protocol used for the callout, for example, Password, Oauth, or NoAuthentication."},
			{Name: "callout_status", Type: proto.ColumnType_STRING, Description: "The status of the named credential's callout configuration."},
			{Name: "namespace_prefix", Type: proto.ColumnType_STRING, Description: "The namespace prefix if the named credential is part of a managed package."},
		}),
	}
}
