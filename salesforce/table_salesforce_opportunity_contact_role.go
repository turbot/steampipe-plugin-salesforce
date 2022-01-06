package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceOpportunityContactRole(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_opportunity_contact_role",
		Description: "salesforce Opportunity Contact Role",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceOpportunityContactRole,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceOpportunityContactRole,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "contact_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "opportunity_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "role", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_primary", Type: proto.ColumnType_BOOL, Description: ""},
		},
	}
}

func listSalesforceOpportunityContactRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSalesforceOpportunityContactRole", "connect error", err)
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "OpportunityContactRole")
	plugin.Logger(ctx).Info("listSalesforceOpportunityContactRole", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceOpportunityContactRole", "query error", err)
			return nil, err
		}

		OpportunityContactRoles := new([]OpportunityContactRole)
		err = decodeQueryResult(ctx, result.Records, OpportunityContactRoles)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceOpportunityContactRole", "decode results error", err)
			return nil, err
		}

		for _, opportunity := range *OpportunityContactRoles {
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

func getSalesforceOpportunityContactRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	opportunityContactRoleID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceOpportunityContactRole", "connect error", err)
		return nil, err
	}

	obj := client.SObject("OpportunityContactRole").Get(opportunityContactRoleID)
	if obj == nil {
		// Object doesn't exist, handle the error
		plugin.Logger(ctx).Warn("getSalesforceOpportunityContactRole", fmt.Sprintf("opportunity contact role \"%s\" not found", opportunityContactRoleID))
		return nil, nil
	}

	opportunityContactRole := new(OpportunityContactRole)
	err = decodeQueryResult(ctx, obj, opportunityContactRole)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceOpportunityContactRole", "decode results error", err)
		return nil, err
	}

	return opportunityContactRole, nil
}
