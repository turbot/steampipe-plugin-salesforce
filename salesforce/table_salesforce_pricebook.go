package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforcePricebook(ctx context.Context, dm dynamicMap) *plugin.Table {
	tableName := "Pricebook2"
	return &plugin.Table{
		Name:        "salesforce_pricebook",
		Description: "Represents a price book that contains the list of products that your org sells.",
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the product in pricebook."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The Price Book Name."},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: "Indicates whether the price book is active (true) or not (false)."},

			// Other columns
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who created the price book."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of the creation of the price book record."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the price book."},
			{Name: "is_archived", Type: proto.ColumnType_BOOL, Description: "Describes whether the price book is archived. The default value is false."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Indicates whether the price book has been archived (true) or not (false)."},
			{Name: "is_standard", Type: proto.ColumnType_BOOL, Description: "Indicates whether the price book is the standard price book for the org (true) or not (false). Every org has one standard price book—all other price books are custom price books."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who most recently changed the product record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of most recent change in the product record."},
			{Name: "last_referenced_date", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp for when the current user last viewed a record related to this record."},
			{Name: "last_viewed_date", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp for when the current user last viewed this record. If this value is null, it's possible that this record was referenced (LastReferencedDate) and not viewed."},
			{Name: "system_modstamp", Type: proto.ColumnType_STRING, Description: "The date and time when order record was last modified by a user or by an automated process."},
		}),
	}
}
