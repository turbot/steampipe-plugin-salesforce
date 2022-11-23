package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforceProduct(ctx context.Context, dm dynamicMap) *plugin.Table {
	tableName := "Product2"
	return &plugin.Table{
		Name:        "salesforce_product",
		Description: "Represents a product that org sells.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, dm.salesforceColumns),
			KeyColumns: dm.keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, dm.cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the product in Salesforce."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The product's name."},
			{Name: "family", Type: proto.ColumnType_STRING, Description: "Name of the product family associated with this record."},
			{Name: "product_code", Type: proto.ColumnType_STRING, Description: "The internal code or product number that you use to identify the product."},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: "Indicates that the product is ready for use in a price book, opportunity, or quote, and whether you can see the product in views."},

			// Other columns
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who created the product record, with the date and time of creation."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of the creation of the product role record."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the product."},
			{Name: "display_url", Type: proto.ColumnType_STRING, Description: "URL leading to a specific version of a record in the linked external data source."},
			{Name: "external_data_source_id", Type: proto.ColumnType_STRING, Description: "The id of the related external data source."},
			{Name: "external_id", Type: proto.ColumnType_STRING, Description: "The unique identifier of a record in the linked external data source."},
			{Name: "is_archived", Type: proto.ColumnType_BOOL, Description: "Describes whether the product is archived. The default value is false."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Indicates whether the object has been moved to the Recycle Bin (true) or not (false)."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who most recently changed the product record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent change in the product record."},
			{Name: "last_referenced_date", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the current user last viewed product record."},
			{Name: "last_viewed_date", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the current user last viewed this record. If this value is null, this record might only have been referenced (last_referenced_date) and not viewed by the current user."},
			{Name: "quantity_unit_of_measure", Type: proto.ColumnType_STRING, Description: "Unit of the product—for example, kilograms, liters, or cases."},
			{Name: "stock_keeping_unit", Type: proto.ColumnType_STRING, Description: "The product's SKU, which can be used with or in place of the Product Code field."},
			{Name: "system_modstamp", Type: proto.ColumnType_STRING, Description: "The date and time when order record was last modified by a user or by an automated process."},
		}),
	}
}
