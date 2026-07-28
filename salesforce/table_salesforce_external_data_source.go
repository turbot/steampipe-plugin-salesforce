package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func SalesforceExternalDataSource(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "ExternalDataSource"
	return &plugin.Table{
		Name:        "salesforce_external_data_source",
		Description: "Represents an external data source, a connection to data stored outside the Salesforce org (for example, via Salesforce Connect).",
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the external data source."},
			{Name: "developer_name", Type: proto.ColumnType_STRING, Description: "The unique name of the external data source in the API."},
			{Name: "master_label", Type: proto.ColumnType_STRING, Description: "The external data source label, displayed in the user interface."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the external data source (identifies the adapter used to connect)."},

			// Other columns
			{Name: "endpoint", Type: proto.ColumnType_STRING, Description: "The URL or connection string for the external system."},
			{Name: "principal_type", Type: proto.ColumnType_STRING, Description: "The type of authentication principal (Anonymous, PerUser, or NamedUser)."},
			{Name: "is_writable", Type: proto.ColumnType_BOOL, Description: "If true, the external data source supports write operations."},
			{Name: "repository", Type: proto.ColumnType_STRING, Description: "The repository of the external data source."},
			{Name: "namespace_prefix", Type: proto.ColumnType_STRING, Description: "The namespace prefix if the external data source is part of a managed package."},
			{Name: "language", Type: proto.ColumnType_STRING, Description: "The language of the external data source label."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who created the external data source."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the external data source was created."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who last modified the external data source."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the external data source was last modified."},
		}),
	}
}
