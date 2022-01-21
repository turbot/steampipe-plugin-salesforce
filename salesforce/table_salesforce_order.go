package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceOrder(ctx context.Context, p *plugin.Plugin) *plugin.Table {
	cols, keyColumns := dynamicColumns(ctx, "Order", p)
	return &plugin.Table{
		Name:        "salesforce_order",
		Description: "Represents an order associated with a contract or an account.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceOrder,
			KeyColumns: keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceOrder,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, p, cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the order in Salesforce."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Title for the order that distinguishes it from other orders."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "ID of the Account associated with this order."},
			{Name: "order_number", Type: proto.ColumnType_STRING, Description: "Order number assigned to this order."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: " ID of the User or queue that owns this order."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The Status field specifies the current state of an order. Status strings represent its current state (Draft or Activated)."},
			{Name: "total_amount", Type: proto.ColumnType_DOUBLE, Description: "Total amount of the order."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of order."},

			// Other columns
			{Name: "activated_by_id", Type: proto.ColumnType_STRING, Description: "ID of the User who activated this order."},
			{Name: "activated_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the order was activated."},
			{Name: "bill_to_contact_id", Type: proto.ColumnType_STRING, Description: "ID of the contact that the order is billed to."},
			{Name: "billing_address", Type: proto.ColumnType_JSON, Description: "The billing adress for the order."},
			{Name: "company_authorized_by_id", Type: proto.ColumnType_STRING, Description: "ID of the user who authorized the account associated with the order."},
			{Name: "company_authorized_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date on which your organization authorized the order."},
			{Name: "contract_id", Type: proto.ColumnType_STRING, Description: "ID of the contract associated with this order. Can only be updated when the order's StatusCode value is Draft."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created the order record."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date and time of the order record."},
			{Name: "customer_authorized_by_id", Type: proto.ColumnType_STRING, Description: "ID of the contact who authorized the order."},
			{Name: "customer_authorized_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date on which the contact authorized the order."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the order."},
			{Name: "effective_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date at which the order becomes effective."},
			{Name: "end_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date at which the order ends."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Indicates that the order is deleted."},
			{Name: "is_reduction_order", Type: proto.ColumnType_BOOL, Description: "Determines whether an order is a reduction order."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who most recently changed the order record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent change in the order record."},
			{Name: "last_referenced_date", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the current user last accessed this record, a record related to this record, or a list view."},
			{Name: "last_viewed_date", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the current user last viewed this record or list view. If this value is null, the user might have only accessed this record or list view (LastReferencedDate) but not viewed it."},
			{Name: "order_reference_number", Type: proto.ColumnType_STRING, Description: "Reference number assigned to the order."},
			{Name: "original_order_id", Type: proto.ColumnType_STRING, Description: "Optional. ID of the original order that a reduction order is reducing, if the reduction order is reducing a single order."},
			{Name: "po_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of the purchase order."},
			{Name: "po_number", Type: proto.ColumnType_STRING, Description: "Number identifying the purchase order."},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: "ID of the price book associated with this order."},
			{Name: "ship_to_contact_id", Type: proto.ColumnType_STRING, Description: "ID of the contact that the order is shipped to."},
			{Name: "status_code", Type: proto.ColumnType_STRING, Description: "Status code of the stage that the order has reached in the order business process."},
			{Name: "system_modstamp", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when order record was last modified by a user or by an automated process."},
			{Name: "shipping_address", Type: proto.ColumnType_JSON, Description: "The shipping adress for the order."},
		}),
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
