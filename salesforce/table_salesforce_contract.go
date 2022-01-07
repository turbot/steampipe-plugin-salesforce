package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceContract(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_contract",
		Description: "Salesforce Contract",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceContract,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceContract,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the contract in Salesforce."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "Id of the account associated with the contract."},
			{Name: "contract_number", Type: proto.ColumnType_STRING, Description: "Unique number assigned to the contract."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "The user id of the account owner for the contract."},

			{Name: "activated_by_id", Type: proto.ColumnType_STRING, Description: "Id of the person who activated the contract."},
			{Name: "activated_date", Type: proto.ColumnType_TIMESTAMP, Description: "Activation date of the contract."},
			{Name: "billing_address", Type: proto.ColumnType_JSON, Description: "Billing address of the account."},
			{Name: "company_signed_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the contract was authorized."},
			{Name: "company_signed_id", Type: proto.ColumnType_STRING, Description: "User at company who authorized the contract."},
			{Name: "contract_term", Type: proto.ColumnType_INT, Description: "Number of months that the contract is in effect."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created the contract record."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when contract record was created."},
			{Name: "customer_signed_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the contact signed the contract."},
			{Name: "customer_signed_id", Type: proto.ColumnType_STRING, Description: "Contact on the account that authorized the contract."},
			{Name: "customer_signed_title", Type: proto.ColumnType_STRING, Description: "Title of the contact who signed the contract."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Statement describing the contract."},
			{Name: "end_date", Type: proto.ColumnType_TIMESTAMP, Description: "Last day when the contract is in effect."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Deletion status of the contract."},
			{Name: "last_activity_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time of the last change to contract record."},
			{Name: "last_approved_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time of the last approval."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "The id of user who most recently changed the contract record."},
			{Name: "last_modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time of the last change to contract record."},
			{Name: "last_referenced_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time of the last reference of contract."},
			{Name: "last_viewed_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time contract was last viewed."},
			{Name: "owner_expiration_notice", Type: proto.ColumnType_STRING, Description: "Number of days before the contract end date to send the notification, if the contract owner and account owner wants to be notified of an upcoming contract expiration."},
			{Name: "pricebook_2_id", Type: proto.ColumnType_STRING, Description: "Id of the pricebook to track the contract."},
			{Name: "special_terms", Type: proto.ColumnType_STRING, Description: "Any terms that are agreed to in the contract."},
			{Name: "start_date", Type: proto.ColumnType_TIMESTAMP, Description: "The first day date when the contract is in effect."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Stage that the contract has reached in the contract business process."},
			{Name: "system_modstamp", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when contract was last modified by a user or by an automated process."},
		},
	}
}

func listSalesforceContract(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSalesforceContract", "connect error", err)
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "Contract")
	plugin.Logger(ctx).Info("listSalesforceContract", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceContract", "query error", err)
			return nil, err
		}

		ContractList := new([]Contract)
		err = decodeQueryResult(ctx, result.Records, ContractList)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceContract", "decode results error", err)
			return nil, err
		}

		for _, contract := range *ContractList {
			d.StreamListItem(ctx, contract)
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

func getSalesforceContract(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	contractID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceContract", "connect error", err)
		return nil, err
	}

	obj := client.SObject("Contract").Get(contractID)
	if obj == nil {
		// Object doesn't exist, handle the error
		plugin.Logger(ctx).Warn("getSalesforceContract", fmt.Sprintf("contract \"%s\" not found", contractID))
		return nil, nil
	}

	account := new(Account)
	err = decodeQueryResult(ctx, obj, account)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceContract", "decode results error", err)
		return nil, err
	}

	return nil, nil
}
