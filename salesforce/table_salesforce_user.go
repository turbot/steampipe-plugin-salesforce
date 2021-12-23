package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func SalesforceUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "salesforce_user",
		Description: "salesforce_user",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceUser,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceUser,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "username", Type: proto.ColumnType_STRING},
			{Name: "last_modified_by_id", Type: proto.ColumnType_STRING},
			{Name: "last_login_date", Type: proto.ColumnType_STRING},
			{Name: "profile_id", Type: proto.ColumnType_STRING},
			{Name: "created_by_id", Type: proto.ColumnType_STRING},
			{Name: "is_active", Type: proto.ColumnType_BOOL},
		},
	}
}

func listSalesforceUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	query := generateQuery(d.QueryContext.Columns, "User")
	plugin.Logger(ctx).Info("listSalesforceUser", "Query", query)

	for {
		result, err := client.Query(query)
		if err != nil {
			return nil, err
		}

		Users := new([]User)
		err = decodeQueryResult(ctx, result.Records, Users)
		if err != nil {
			return nil, err
		}

		for _, account := range *Users {
			d.StreamListItem(ctx, account)
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
		return nil, err
	}

	obj := client.SObject("User").Get(userID)
	if obj == nil {
		// Object doesn't exist, handle the error
		return nil, nil
	}

	user := new(User)
	err = decodeQueryResult(ctx, obj, user)
	if err != nil {
		return nil, err
	}

	return user, nil

}
