package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceOpportunity(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_opportunity",
		Description: "In Salesforce, an opportunity is a sale or pending deal.",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceOpportunity,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceOpportunity,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the opportunity in Salesforce."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the opportunity, for example, Acme.com - Office Equipment Order."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "Id of the account that the opportunity is linked to."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "Assigned owner of opportunity."},
			{Name: "amount", Type: proto.ColumnType_DOUBLE, Description: "Estimated total sale amount."},

			{Name: "campaign_id", Type: proto.ColumnType_STRING, Description: "Id of the campaign responsible for generating the opportunity."},
			{Name: "close_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when plan to close the opportunity."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created the opportunity."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The creation date and time of the opportunity."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the opportunity."},
			{Name: "expected_revenue", Type: proto.ColumnType_DOUBLE, Description: "Calculated revenue based on the Amount and Probability fields."},
			{Name: "fiscal_quarter", Type: proto.ColumnType_INT, Description: "Quarter of the year when the oppurtiny was creadted."},
			{Name: "fiscal_year", Type: proto.ColumnType_INT, Description: "Year when the oppurtiny was creadted."},
			{Name: "forecast_category", Type: proto.ColumnType_STRING, Description: "Forecast category name displayed in reports, opportunity detail and edit pages, opportunity searches, and opportunity list views."},
			{Name: "forecast_category_name", Type: proto.ColumnType_STRING, Description: "Name of the forecast category."},
			{Name: "has_open_activity", Type: proto.ColumnType_BOOL, Description: "True, if opportunity has open any activity."},
			{Name: "has_opportunity_line_item", Type: proto.ColumnType_BOOL, Description: "True, if has any opportunity line item."},
			{Name: "has_overdue_task", Type: proto.ColumnType_BOOL, Description: "True, if any overdue task is left."},
			{Name: "is_closed", Type: proto.ColumnType_BOOL, Description: "Indicates that the opportunity is closed."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Indicates that the opportunity is deleted."},
			{Name: "is_private", Type: proto.ColumnType_BOOL, Description: "Indicates that the opportunity is private."},
			{Name: "is_won", Type: proto.ColumnType_BOOL, Description: "Indicates that the quote or proposal has been signed or electronically accepted."},
			{Name: "last_activity_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time of the last activity on the opportunity."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who last modified the oppurtinity record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The data and time of the last modification of the oppurtinity record."},
			{Name: "lead_source", Type: proto.ColumnType_STRING, Description: "Source of the opportunity, such as Advertisement, Partner, or Web."},
			{Name: "next_step", Type: proto.ColumnType_STRING, Description: "Description of next task in closing opportunity."},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: "A PriceBookEntry record is a unique combination of a PriceBookID, a Product2Id, a Currency, and an Active status."},
			{Name: "probability", Type: proto.ColumnType_DOUBLE, Description: "Likelihood that opportunity will close, stated as a percentage."},
			{Name: "stage_name", Type: proto.ColumnType_STRING, Description: "Current stage of opportunity."},
			{Name: "system_modstamp", Type: proto.ColumnType_STRING, Description: "The date and time when opportunity was last modified by a user or by an automated process."},
			{Name: "total_opportunity_quantity", Type: proto.ColumnType_STRING, Description: "Total of all Quantity field values for all products in the Products related list if the opportunity has products."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of opportunity, such as Existing Business or New Business."},
		},
	}
}

func listSalesforceOpportunity(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSalesforceOpportunity", "connect error", err)
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Opportunity")
	plugin.Logger(ctx).Info("listSalesforceOpportunity", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceOpportunity", "query error", err)
			return nil, err
		}

		Opportunities := new([]Opportunity)
		err = decodeQueryResult(ctx, result.Records, Opportunities)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceOpportunity", "decode results error", err)
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
		plugin.Logger(ctx).Error("getSalesforceOpportunity", "connect error", err)
		return nil, err
	}

	obj := client.SObject("Opportunity").Get(opportunityID)
	if obj == nil {
		// Object doesn't exist, handle the error
		plugin.Logger(ctx).Warn("getSalesforceOpportunity", fmt.Sprintf("opportunity \"%s\" not found", opportunityID))
		return nil, nil
	}

	opportunity := new(Opportunity)
	err = decodeQueryResult(ctx, obj, opportunity)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceOpportunity", "decode results error", err)
		return nil, err
	}

	return opportunity, nil
}