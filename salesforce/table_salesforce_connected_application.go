package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// ConnectedApplication is queryable through standard SOQL, but exposes only a
// limited field set (an inventory plus creation/modification audit) — its
// OAuth/session policy options are not available on the standard sObject, and
// it is not supported by the Tooling API query endpoint. This gives a connected-
// app inventory with change audit; policy detail must come from the Setup Audit
// Trail (Connected App Session Policy changes) instead.
func SalesforceConnectedApplication(ctx context.Context, config salesforceConfig) *plugin.Table {
	tableName := "ConnectedApplication"
	return &plugin.Table{
		Name:        "salesforce_connected_application",
		Description: "Represents a connected app — an external application that integrates with Salesforce via OAuth. Inventory and creation/modification audit, queried through standard SOQL.",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceObjectsByTable(tableName, map[string]string{}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the organization in Salesforce.", Hydrate: getOrganizationId, Transform: transform.FromValue()},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the connected application."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the connected application."},
		},
	}
}
