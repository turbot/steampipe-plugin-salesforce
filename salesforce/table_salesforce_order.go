package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceOrder(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_order",
		Description: "Salesforce Order",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceOrder,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceOrder,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the order in Salesforce."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Title for the order that distinguishes it from other orders."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "Id of the account associated with the order."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Indication of the stage that the order has reached in the order business process."},
			{Name: "total_amount", Type: proto.ColumnType_DOUBLE, Description: "Total amount of the order."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of order."},

			{Name: "activated_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who activated the order."},
			{Name: "activated_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and timestamp when the order was activated."},
			{Name: "bill_to_contact_id", Type: proto.ColumnType_STRING, Description: "Id of the contact to whom the order is billed."},
			{Name: "billing_address", Type: proto.ColumnType_JSON, Description: "The billing adress for the order."},
			{Name: "company_authorized_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user at your company who authorized the order."},
			{Name: "company_authorized_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the order was authorized."},
			{Name: "contract_id", Type: proto.ColumnType_STRING, Description: "Id of the parent contract that distinguishes it from other contracts."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created the order record."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date and time of the order record."},
			{Name: "customer_authorized_by_id", Type: proto.ColumnType_STRING, Description: "Id of the contact on the order's account who authorized the order."},
			{Name: "customer_authorized_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the contact authorized the order."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the order."},
			{Name: "effective_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the order becomes effective."},
			{Name: "end_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the order ends."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Indicates that the order is deleted."},
			{Name: "is_reduction_order", Type: proto.ColumnType_BOOL, Description: "Indicates that the order is a reduction order."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who most recently changed the order record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent change in the order record."},
			{Name: "last_referenced_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent reference of the order record."},
			{Name: "last_viewed_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent viewing of the order record."},
			{Name: "order_number", Type: proto.ColumnType_STRING, Description: "Unique number automatically assigned to the order."},
			{Name: "order_reference_number", Type: proto.ColumnType_STRING, Description: "Reference number assigned to the order."},
			{Name: "original_order_id", Type: proto.ColumnType_STRING, Description: "For reduction orders only. Prepopulated with the ID of the parent order when you create a reduction order by clicking Reduce Order."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "Id of the Owner of the order. Can be a user or queue."},
			{Name: "po_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when a purchase order was entered."},
			{Name: "po_number", Type: proto.ColumnType_STRING, Description: "Number of the purchase order."},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ship_to_contact_id", Type: proto.ColumnType_STRING, Description: "Id of the contact to whom the order is shipped."},
			{Name: "shipping_address", Type: proto.ColumnType_JSON, Description: "The shipping adress for the order."},
			{Name: "status_code", Type: proto.ColumnType_STRING, Description: "Status code of the stage that the order has reached in the order business process."},
			{Name: "system_modstamp", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when order record was last modified by a user or by an automated process."},
		},
	}
}

func listSalesforceOrder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSalesforceOrder", "connect error", err)
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Order")
	plugin.Logger(ctx).Info("listSalesforceOrder", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceOrder", "query error", err)
			return nil, err
		}

		Orders := new([]Order)
		err = decodeQueryResult(ctx, result.Records, Orders)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceOrder", "decode results error", err)
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
		plugin.Logger(ctx).Error("getSalesforceOrder", "connect error", err)
		return nil, err
	}

	obj := client.SObject("Order").Get(orderID)
	if obj == nil {
		// Object doesn't exist, handle the error
		plugin.Logger(ctx).Warn("getSalesforceOrder", fmt.Sprintf("order \"%s\" not found", orderID))
		return nil, nil
	}

	order := new(Order)
	err = decodeQueryResult(ctx, obj, order)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceOrder", "decode results error", err)
		return nil, err
	}

	return order, nil
}
