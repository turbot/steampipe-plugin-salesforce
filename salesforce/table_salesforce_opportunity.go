package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceOpportunity(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_opportunity",
		Description: "salesforce Opportunity",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceOpportunity,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceOpportunity,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "amount", Type: proto.ColumnType_DOUBLE, Description: ""},
			{Name: "campaign_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "close_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "expected_revenue", Type: proto.ColumnType_DOUBLE, Description: ""},
			{Name: "fiscal_quarter", Type: proto.ColumnType_DOUBLE, Description: ""},
			{Name: "fiscal_year", Type: proto.ColumnType_INT, Description: ""},
			{Name: "forecast_category", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "forecast_category_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "has_open_activity", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "has_opportunity_line_item", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "has_overdue_task", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "is_closed", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "is_private", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "is_won", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "last_activity_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_referenced_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_viewed_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "lead_source", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "next_step", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "probability", Type: proto.ColumnType_DOUBLE, Description: ""},
			{Name: "stage_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "system_modstamp", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "total_opportunity_quantity", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			// {Name: "current_generators_c", Type: proto.ColumnType_STRING, Description: ""},
			// {Name: "delivery_installation_status_c", Type: proto.ColumnType_STRING, Description: ""},
			// {Name: "main_competitors_c", Type: proto.ColumnType_STRING, Description: ""},
			// {Name: "order_number_c", Type: proto.ColumnType_STRING, Description: ""},
			// {Name: "tracking_number_c", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listSalesforceOpportunity(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Opportunity")
	plugin.Logger(ctx).Info("listSalesforceOpportunity", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			return nil, err
		}

		Opportunities := new([]Opportunity)
		err = decodeQueryResult(ctx, result.Records, Opportunities)
		if err != nil {
			return nil, err
		}

		for _, opportunity := range *Opportunities {
			d.StreamListItem(ctx, opportunity)
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

func getSalesforceOpportunity(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	opportunityID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	obj := client.SObject("Opportunity").Get(opportunityID)
	if obj == nil {
		// Object doesn't exist, handle the error
		return nil, nil
	}

	opportunity := new(Opportunity)
	err = decodeQueryResult(ctx, obj, opportunity)
	if err != nil {
		return nil, err
	}

	return opportunity, nil
}
