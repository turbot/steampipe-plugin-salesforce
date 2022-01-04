package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceOrder(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_order",
		Description: "salesforce_order",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceOrder,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceOrder,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "activated_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "activated_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "bill_to_contact_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "company_authorized_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "company_authorized_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "contract_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customer_authorized_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customer_authorized_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "effective_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "end_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "is_reduction_order", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_referenced_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_viewed_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "order_number", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "order_reference_number", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "original_order_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "po_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "po_number", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ship_to_contact_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "status", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "status_code", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "system_modstamp", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "total_amount", Type: proto.ColumnType_DOUBLE, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "shipping_address", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "billing_address", Type: proto.ColumnType_JSON, Description: ""},
		},
	}
}

func listSalesforceOrder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Order")
	plugin.Logger(ctx).Info("listSalesforceUser", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			return nil, err
		}

		Orders := new([]Order)
		err = decodeQueryResult(ctx, result.Records, Orders)
		if err != nil {
			return nil, err
		}

		for _, order := range *Orders {
			d.StreamListItem(ctx, order)
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

func getSalesforceOrder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	orderID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	obj := client.SObject("Order").Get(orderID)
	if obj == nil {
		// Object doesn't exist, handle the error
		return nil, nil
	}

	order := new(Order)
	err = decodeQueryResult(ctx, obj, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
