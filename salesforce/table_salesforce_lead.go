package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func SalesforceLead(ctx context.Context, dm dynamicMap, p *plugin.Plugin) *plugin.Table {
	tableName := "Lead"
	return &plugin.Table{
		Name:        "salesforce_lead",
		Description: "Represents a prospect or lead.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, dm.salesforceColumns),
			KeyColumns: dm.keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, p, dm.cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the lead in Salesforce."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The lead's email address."},
			{Name: "is_converted", Type: proto.ColumnType_BOOL, Description: "Indicates whether the lead has been converted (true) or not (false)."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the lead, as displayed on lead detail page."},
			{Name: "phone", Type: proto.ColumnType_STRING, Description: "Lead's primary phone number."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the lead, for example, Open, Contacted, Qualified."},

			// Other columns
			{Name: "address", Type: proto.ColumnType_JSON, Description: "Street address for the lead."},
			{Name: "annual_revenue", Type: proto.ColumnType_DOUBLE, Description: "Annual revenue for the lead's company."},
			{Name: "company", Type: proto.ColumnType_STRING, Description: "The lead's company."},
			{Name: "converted_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date on which this lead was converted."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created the lead."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date and time of the lead."},
			{Name: "industry", Type: proto.ColumnType_STRING, Description: "Primary business of lead's company."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who last changed the lead record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of the last changes to lead record."},
			{Name: "lead_source", Type: proto.ColumnType_STRING, Description: "Source of lead, for example, Advertisement, Partner, or Web."},
			{Name: "number_of_employees", Type: proto.ColumnType_INT, Description: "Number of employees at the lead's company."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "Id of the assigned owner of the lead."},
			{Name: "rating", Type: proto.ColumnType_STRING, Description: "Indicates value or prospect of the lead, for example, Hot, Warm, Cold."},
			{Name: "website", Type: proto.ColumnType_STRING, Description: "URL of the lead's company's website."},
		}),
	}
}
