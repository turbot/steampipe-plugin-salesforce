package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceLead(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_lead",
		Description: "Represents a prospect or lead.",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceLead,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceLead,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
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
		},
	}
}

func listSalesforceLead(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSalesforceLead", "connect error", err)
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Lead")
	plugin.Logger(ctx).Info("listSalesforceLead", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceLead", "query error", err)
			return nil, err
		}

		Leads := new([]Lead)
		err = decodeQueryResult(ctx, result.Records, Leads)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceLead", "decode results error", err)
			return nil, err
		}

		for _, account := range *Leads {
			d.StreamListItem(ctx, account)
		}

		// Paging
		if result.Done {
			break
		} else {
			query = result.NextRecordsURL
		}
	}

	return nil, nil
}

func getSalesforceLead(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	leadID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceLead", "connect error", err)
		return nil, err
	}

	obj := client.SObject("Lead").Get(leadID)
	if obj == nil {
		// Object doesn't exist, handle the error
		plugin.Logger(ctx).Warn("getSalesforceLead", fmt.Sprintf("lead \"%s\" not found", leadID))
		return nil, nil
	}

	lead := new(User)
	err = decodeQueryResult(ctx, obj, lead)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceLead", "decode results error", err)
		return nil, err
	}

	return lead, nil
}
