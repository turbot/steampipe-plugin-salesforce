package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforceContract(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "Contract"
	return &plugin.Table{
		Name:        "salesforce_contract",
		Description: "Represents a contract (a business agreement) associated with an Account.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, dm.salesforceColumns),
			KeyColumns: dm.keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName),
			KeyColumns: plugin.SingleColumn(checkNameScheme(config, dm.cols)),
		},
		Columns: mergeTableColumns(ctx, config, dm.cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the contract in Salesforce."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "ID of the Account associated with this contract."},
			{Name: "contract_number", Type: proto.ColumnType_STRING, Description: "Number of the contract."},
			{Name: "contract_term", Type: proto.ColumnType_INT, Description: "Number of months that the contract is valid."},
			{Name: "end_date", Type: proto.ColumnType_TIMESTAMP, Description: "Calculated end date of the contract. This value is calculated by adding the ContractTerm to the start_date."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "ID of the user who owns the contract."},
			{Name: "start_date", Type: proto.ColumnType_TIMESTAMP, Description: "Start date for this contract."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The picklist of values that indicate order status. A contract can be Draft, InApproval, or Activated."},

			// Other columns
			{Name: "activated_by_id", Type: proto.ColumnType_STRING, Description: "ID of the User who activated this contract."},
			{Name: "activated_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when this contract was activated."},
			{Name: "billing_address", Type: proto.ColumnType_JSON, Description: "Billing address of the account."},
			{Name: "company_signed_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date on which the contract was signed by organization."},
			{Name: "company_signed_id", Type: proto.ColumnType_STRING, Description: "ID of the User who signed the contract."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created the contract record."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when contract record was created."},
			{Name: "customer_signed_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date on which the customer signed the contract."},
			{Name: "customer_signed_id", Type: proto.ColumnType_STRING, Description: "ID of the Contact who signed this contract."},
			{Name: "customer_signed_title", Type: proto.ColumnType_STRING, Description: "Title of the contact who signed the contract."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Statement describing the contract."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Indicates whether the object has been moved to the Recycle Bin (true) or not (false)."},
			{Name: "last_activity_date", Type: proto.ColumnType_TIMESTAMP, Description: "Value is one of the following, whichever is the most recent. a) Due date of the most recent event logged against the record. b) Due date of the most recently closed task associated with the record."},
			{Name: "last_approved_date", Type: proto.ColumnType_TIMESTAMP, Description: "Last date the contract was approved."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The id of user who most recently changed the contract record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time of the last change to contract record."},
			{Name: "last_referenced_date", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the current user last accessed this record, a record related to this record, or a list view."},
			{Name: "last_viewed_date", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the current user last viewed this record or list view. If this value is null, the user might have only accessed this record or list view (last_referenced_date) but not viewed it."},
			{Name: "owner_expiration_notice", Type: proto.ColumnType_STRING, Description: "Number of days ahead of the contract end date (15, 30, 45, 60, 90, and 120). Used to notify the owner in advance that the contract is ending."},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: "ID of the pricebook, if any, associated with this contract."},
			{Name: "special_terms", Type: proto.ColumnType_STRING, Description: "Special terms that apply to the contract."},
			{Name: "system_modstamp", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when contract was last modified by a user or by an automated process."},
		}),
	}
}
