package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func SalesforceOrganization(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "Organization"
	return &plugin.Table{
		Name:        "salesforce_organization",
		Description: "Represents key configuration information for a Salesforce organization.",
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique id of the organization."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the organization."},
			{Name: "organization_type", Type: proto.ColumnType_STRING, Description: "The edition of the organization, for example, Enterprise or Developer Edition."},
			{Name: "instance_name", Type: proto.ColumnType_STRING, Description: "The instance where the organization is hosted."},
			{Name: "is_sandbox", Type: proto.ColumnType_BOOL, Description: "If true, the organization is a sandbox."},

			// Other columns
			{Name: "division", Type: proto.ColumnType_STRING, Description: "The division of the organization."},
			{Name: "street", Type: proto.ColumnType_STRING, Description: "The street of the organization's primary address."},
			{Name: "city", Type: proto.ColumnType_STRING, Description: "The city of the organization's primary address."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The state of the organization's primary address."},
			{Name: "postal_code", Type: proto.ColumnType_STRING, Description: "The postal code of the organization's primary address."},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "The country of the organization's primary address."},
			{Name: "primary_contact", Type: proto.ColumnType_STRING, Description: "The primary contact for the organization."},
			{Name: "language_locale_key", Type: proto.ColumnType_STRING, Description: "The default language locale of the organization."},
			{Name: "default_locale_sid_key", Type: proto.ColumnType_STRING, Description: "The default locale of the organization."},
			{Name: "fiscal_year_start_month", Type: proto.ColumnType_INT, Description: "The month in which the fiscal year starts."},
			{Name: "trial_expiration_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time a trial organization expires, if applicable."},
			{Name: "namespace_prefix", Type: proto.ColumnType_STRING, Description: "The namespace prefix registered for the organization, if any."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the organization was created."},
		}),
	}
}
