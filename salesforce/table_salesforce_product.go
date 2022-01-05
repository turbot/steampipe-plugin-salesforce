package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceProduct(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_product",
		Description: "salesforce Product",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceProduct,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceProduct,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "display_url", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "external_data_source_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "external_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "family", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "is_archived", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_modified_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_referenced_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_viewed_date", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "product_code", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "quantity_unit_of_measure", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "stock_keeping_unit", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "system_modstamp", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listSalesforceProduct(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Product2")
	plugin.Logger(ctx).Info("listSalesforceProduct", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			return nil, err
		}

		Products := new([]Product)
		err = decodeQueryResult(ctx, result.Records, Products)
		if err != nil {
			return nil, err
		}

		for _, product := range *Products {
			d.StreamListItem(ctx, product)
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

func getSalesforceProduct(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	userID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	obj := client.SObject("Product2").Get(userID)
	if obj == nil {
		// Object doesn't exist, handle the error
		return nil, nil
	}

	product := new(Product)
	err = decodeQueryResult(ctx, obj, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
