package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforceAccount(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "Account"
	return &plugin.Table{
		Name:        "salesforce_account",
		Description: "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, dm.salesforceColumns),
			KeyColumns: dm.keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableAccountColumns(ctx, config, dm.cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the account in Salesforce."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the account."},
			{Name: "annual_revenue", Type: proto.ColumnType_DOUBLE, Description: "Estimated annual revenue of the account."},
			{Name: "industry", Type: proto.ColumnType_STRING, Description: "Primary business of account."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "The ID of the user who currently owns this account. Default value is the user logged in to the API to perform the create."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of account, for example, Customer, Competitor, or Partner."},

			// Other columns
			{Name: "account_source", Type: proto.ColumnType_STRING, Description: "The source of the account record. For example, Advertisement, Data.com, or Trade Show."},
			// {Name: "clean_status", Type: proto.ColumnType_STRING, Description: "Indicates the record's clean status as compared with Data.com. Values are: Matched, Different,Acknowledged,NotFound,Inactive,Pending, SelectMatch, or Skipped."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who created the account."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "The creation date and time of the account."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Text description of the account."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Indicates whether the object has been moved to the Recycle Bin (true) or not (false)."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The id of the user who last changed the contact fields, including modification date and time."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time of last modification to account."},
			{Name: "number_of_employees", Type: proto.ColumnType_DOUBLE, Description: "Number of employees working at the company represented by this account."},
			// {Name: "ownership", Type: proto.ColumnType_STRING, Description: "Ownership type for the account, for example Private, Public, or Subsidiary."},
			{Name: "phone", Type: proto.ColumnType_STRING, Description: "The contact's primary phone number."},
			// {Name: "rating", Type: proto.ColumnType_STRING, Description: "The account's prospect rating, for example Hot, Warm, or Cold."},
			// {Name: "sic", Type: proto.ColumnType_STRING, Description: "Standard Industrial Classification code of the company's main business categorization, for example, 57340 for Electronics."},
			// {Name: "ticker_symbol", Type: proto.ColumnType_STRING, Description: "The stock market symbol for this account."},
			// {Name: "tradestyle", Type: proto.ColumnType_STRING, Description: "A name, different from its legal name, that an org may use for conducting business. Similar to “Doing business as” or \"DBA\"."},
			// {Name: "website", Type: proto.ColumnType_STRING, Description: "The website of this account, for example, www.acme.com."},

			// JSON columns
			{Name: "billing_address", Type: proto.ColumnType_JSON, Description: "The billing adress of the account."},
			{Name: "shipping_address", Type: proto.ColumnType_JSON, Description: "The shipping adress of the account."},
		}),
	}
}
