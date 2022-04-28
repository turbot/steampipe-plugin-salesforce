package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func SalesforceOpportunityContactRole(ctx context.Context, dm dynamicMap, p *plugin.Plugin) *plugin.Table {
	tableName := "OpportunityContactRole"
	return &plugin.Table{
		Name:        "salesforce_opportunity_contact_role",
		Description: "Represents the role that a Contact plays on an Opportunity.",
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the opportunity contact role in Salesforce."},
			{Name: "contact_id", Type: proto.ColumnType_STRING, Description: "ID of an associated Contact."},
			{Name: "is_primary", Type: proto.ColumnType_BOOL, Description: "Indicates whether the associated Contact plays the primary role on the Opportunity (true) or not (false)."},
			{Name: "opportunity_id", Type: proto.ColumnType_STRING, Description: "Required. ID of an associated Opportunity."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Name of the role played by the associated Contact on the Opportunity, such as Business User or Decision Maker."},

			// Other columns
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created contact role record."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of the creation of the contact role record."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who most recently changed the contact role record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent change in the contact role record."},
		}),
	}
}
