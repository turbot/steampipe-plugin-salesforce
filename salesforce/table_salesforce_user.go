package salesforce

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceUser(ctx context.Context, p *plugin.Plugin) *plugin.Table {
	cols, keyColumns := dynamicColumns(ctx, "User", p)
	return &plugin.Table{
		Name:        "salesforce_user",
		Description: "Represents a user in organization.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceUser,
			KeyColumns: keyColumns,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceUser,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: mergeTableColumns(ctx, p, cols, []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the user in Salesforce."},
			{Name: "alias", Type: proto.ColumnType_STRING, Description: "The user's alias. For example, jsmith."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Login name of the user."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Display name of the user."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The user's email address."},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: "Indicates whether the user has access to log in (true) or not (false)."},

			// Other columns
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "ID of the Account associated with a Customer Portal user. This field is null for Salesforce users."},
			{Name: "created_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who created the user including creation date and time."},
			{Name: "department", Type: proto.ColumnType_STRING, Description: "The company department associated with the user."},
			{Name: "employee_number", Type: proto.ColumnType_STRING, Description: "The user's employee number."},
			{Name: "forecast_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether the user is enabled as a forecast manager (true) or not (false)."},
			{Name: "last_login_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when the user last successfully logged in. This value is updated if 60 seconds elapses since the user's last login."},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING, Description: "Id of the user who last changed the user fields, including modification date and time."},
			{Name: "profile_id", Type: proto.ColumnType_STRING, Description: "ID of the user's Profile."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The state associated with the User."},
			{Name: "user_type", Type: proto.ColumnType_STRING, Description: "The category of user license. Can be one of Standard, PowerPartner, CSPLitePortal, CustomerSuccess, PowerCustomerSuccess, CsnOnly, and Guest."},
		}),
	}
}

func listSalesforceUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSalesforceUser", "connect error", err)
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "User")
	plugin.Logger(ctx).Info("listSalesforceUser", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceUser", "query error", err)
			return nil, err
		}

		Users := new([]User)
		err = decodeQueryResult(ctx, result.Records, Users)
		if err != nil {
			plugin.Logger(ctx).Error("listSalesforceUser", "decode results error", err)
			return nil, err
		}

		for _, user := range *Users {
			d.StreamListItem(ctx, user)
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

func getSalesforceUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	userID := d.KeyColumnQualString("id")

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceUser", "connect error", err)
		return nil, err
	}

	obj := client.SObject("User").Get(userID)
	if obj == nil {
		// Object doesn't exist, handle the error
		plugin.Logger(ctx).Warn("getSalesforceUser", fmt.Sprintf("user \"%s\" not found", userID))
		return nil, nil
	}

	user := new(User)
	err = decodeQueryResult(ctx, obj, user)
	if err != nil {
		plugin.Logger(ctx).Error("getSalesforceUser", "decode results error", err)
		return nil, err
	}

	return user, nil
}
