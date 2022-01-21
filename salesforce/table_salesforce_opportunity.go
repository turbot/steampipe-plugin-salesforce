package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceOpportunity(ctx context.Context, p *plugin.Plugin) *plugin.Table {
	tableName := "OpportunityContactRole"
	cols, keyColumns := dynamicColumns(ctx, tableName, p)
	return &plugin.Table{
		Name:        "salesforce_opportunity",
		Description: "Represents an opportunity, which is a sale or pending deal.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, cols),
			KeyColumns: keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, p, cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the opportunity in Salesforce."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "ID of the account associated with this opportunity."},
			{Name: "amount", Type: proto.ColumnType_DOUBLE, Description: "Estimated total sale amount. For opportunities with products, the amount is the sum of the related products."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "A name for this opportunity."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "ID of the User who has been assigned to work this opportunity."},

			// Other columns
			{Name: "campaign_id", Type: proto.ColumnType_STRING, Description: "ID of a related Campaign. This field is defined only for those organizations that have the campaign feature Campaigns enabled."},
			{Name: "close_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the opportunity is expected to close."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created the opportunity."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The creation date and time of the opportunity."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the opportunity."},
			{Name: "expected_revenue", Type: proto.ColumnType_DOUBLE, Description: "Calculated revenue based on the Amount and Probability fields."},
			{Name: "fiscal_quarter", Type: proto.ColumnType_INT, Description: "Represents the fiscal quarter. Valid values are 1, 2, 3, or 4."},
			{Name: "fiscal_year", Type: proto.ColumnType_INT, Description: "Represents the fiscal year, for example, 2006."},
			{Name: "forecast_category", Type: proto.ColumnType_STRING, Description: "Forecast category name displayed in reports, opportunity detail and edit pages, opportunity searches, and opportunity list views."},
			{Name: "forecast_category_name", Type: proto.ColumnType_STRING, Description: "Name of the forecast category."},
			{Name: "has_open_activity", Type: proto.ColumnType_BOOL, Description: "Indicates whether an opportunity has an open event or task (true) or not (false)."},
			{Name: "has_opportunity_line_item", Type: proto.ColumnType_BOOL, Description: "Indicates whether the opportunity has associated line items. A value of true means that Opportunity line items have been created for the opportunity."},
			{Name: "has_overdue_task", Type: proto.ColumnType_BOOL, Description: "Indicates whether an opportunity has an overdue task (true) or not (false)."},
			{Name: "is_closed", Type: proto.ColumnType_BOOL, Description: "Indicates that the opportunity is closed."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Indicates that the opportunity is deleted."},
			{Name: "is_private", Type: proto.ColumnType_BOOL, Description: "Indicates that the opportunity is private."},
			{Name: "is_won", Type: proto.ColumnType_BOOL, Description: "Indicates that the quote or proposal has been signed or electronically accepted."},
			{Name: "last_activity_date", Type: proto.ColumnType_TIMESTAMP, Description: "Value is one of the following, whichever is the most recent of a) Due date of the most recent event logged against the record or b) Due date of the most recently closed task associated with the record."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who last modified the oppurtinity record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The data and time of the last modification of the oppurtinity record."},
			{Name: "lead_source", Type: proto.ColumnType_STRING, Description: "Source of this opportunity, such as Advertisement or Trade Show."},
			{Name: "next_step", Type: proto.ColumnType_STRING, Description: "Description of next task in closing opportunity."},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: "ID of a related Pricebook2 object. The Pricebook2Id field indicates which Pricebook2 applies to this opportunity. The Pricebook2Id field is defined only for those organizations that have products enabled as a feature."},
			{Name: "probability", Type: proto.ColumnType_DOUBLE, Description: "Percentage of estimated confidence in closing the opportunity."},
			{Name: "stage_name", Type: proto.ColumnType_STRING, Description: "Current stage of opportunity."},
			{Name: "system_modstamp", Type: proto.ColumnType_STRING, Description: "The date and time when opportunity was last modified by a user or by an automated process."},
			{Name: "total_opportunity_quantity", Type: proto.ColumnType_STRING, Description: "Number of items included in this opportunity. Used in quantity-based forecasting."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of opportunity, such as Existing Business or New Business."},
		}),
	}
}
