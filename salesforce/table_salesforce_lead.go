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
		Description: "Salesforce Lead",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceLead,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceLead,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "email", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "annual_revenue", Type: proto.ColumnType_DOUBLE, Description: ""},
			{Name: "company", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "converted_date", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "industry", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_converted", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "lead_source", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "number_of_employees", Type: proto.ColumnType_INT, Description: ""},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "phone", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "rating", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "status", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "website", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "address", Type: proto.ColumnType_JSON, Description: ""},
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
