package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforceLoginHistory(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "LoginHistory"
	return &plugin.Table{
		Name:        "salesforce_login_history",
		Description: "Represents the login history for all successful and failed login attempts for organizations and enabled portals.",
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the login history entry."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "The id of the user who attempted to log in."},
			{Name: "login_time", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time of the login attempt."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the login attempt, for example, Success or the reason for failure."},

			// Other columns
			{Name: "application", Type: proto.ColumnType_STRING, Description: "The application used to access the org during login."},
			{Name: "browser", Type: proto.ColumnType_STRING, Description: "The browser name and version, if known."},
			{Name: "platform", Type: proto.ColumnType_STRING, Description: "The operating system used during login, if known."},
			{Name: "source_ip", Type: proto.ColumnType_STRING, Description: "The source IP address of the login attempt."},
			{Name: "login_type", Type: proto.ColumnType_STRING, Description: "The type of login, for example, Application, OAuth, or Remote Access."},
			{Name: "login_url", Type: proto.ColumnType_STRING, Description: "The URL used to log in."},
			{Name: "api_type", Type: proto.ColumnType_STRING, Description: "The type of API request during login."},
			{Name: "api_version", Type: proto.ColumnType_STRING, Description: "The API version used during login."},
			{Name: "tls_protocol", Type: proto.ColumnType_STRING, Description: "The TLS protocol version used during login."},
			{Name: "cipher_suite", Type: proto.ColumnType_STRING, Description: "The TLS cipher suite used during login."},
		}),
	}
}
