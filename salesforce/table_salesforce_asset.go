package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforceAsset(ctx context.Context, dm dynamicMap) *plugin.Table {
	tableName := "Asset"
	return &plugin.Table{
		Name:        "salesforce_asset",
		Description: "Represents an item of commercial value, such as a product sold by your company or a competitor, that a customer has purchased and installed.",
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the product in asset."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the asset."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "ID of the Account associated with this asset."},
			{Name: "asset_level", Type: proto.ColumnType_INT, Description: "The asset's position in an asset hierarchy. If the asset has no parent or child assets, its level is 1."},
			{Name: "contact_id", Type: proto.ColumnType_STRING, Description: "ID of the Contact associated with this asset."},

			// Other columns
			{Name: "asset_provided_by_id", Type: proto.ColumnType_STRING, Description: "The account that provided the asset, typically a manufacturer."},
			{Name: "asset_serviced_by_id", Type: proto.ColumnType_STRING, Description: "The account in charge of servicing the asset."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who created the asset."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of the creation of the price book record."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the asset."},
			{Name: "install_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the asset was installed."},
			{Name: "is_competitor_product", Type: proto.ColumnType_BOOL, Description: "Indicates whether this Asset represents a product sold by a competitor (true) or not (false)."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "The Deleted."},
			{Name: "is_internal", Type: proto.ColumnType_BOOL, Description: "Indicates that the asset is produced or used internally (true) or not (false). Default value is false."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The Last Modified By ID."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The Last Modified Date."},
			{Name: "last_referenced_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time that the asset was last modified."},
			{Name: "last_viewed_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time that the asset was last viewed."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "The asset's owner. By default, the asset owner is the user who created the asset record."},
			{Name: "parent_id", Type: proto.ColumnType_STRING, Description: "The asset's parent asset."},
			{Name: "price", Type: proto.ColumnType_DOUBLE, Description: "Price paid for this asset."},
			{Name: "product_2_id", Type: proto.ColumnType_STRING, Description: "ID of the Product2 associated with this asset. Must be a valid Product2 ID."},
			{Name: "product_code", Type: proto.ColumnType_STRING, Description: "The product code of the related product."},
			{Name: "purchase_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date on which this asset was purchased."},
			{Name: "quantity", Type: proto.ColumnType_DOUBLE, Description: "Quantity purchased or installed."},
			{Name: "root_asset_id", Type: proto.ColumnType_STRING, Description: " The top-level asset in an asset hierarchy. Depending on where an asset lies in the hierarchy, its root could be the same as its parent."},
			{Name: "serial_number", Type: proto.ColumnType_STRING, Description: "Serial number for this asset."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Customizable picklist of values. The default picklist includes the following values: Purchased, Shipped, Installed, Registered, Obsolete."},
			{Name: "stock_keeping_unit", Type: proto.ColumnType_STRING, Description: "The SKU assigned to the related product."},
			{Name: "system_modstamp", Type: proto.ColumnType_TIMESTAMP, Description: "The System Modstamp."},
			{Name: "usage_end_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when usage for this asset ends or expires."},
		}),
	}
}
