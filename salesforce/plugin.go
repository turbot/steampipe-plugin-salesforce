package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-salesforce"

// Plugin creates this (salesforce) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"salesforce_account":              SalesforceAccount(ctx),
			"salesforce_account_contact_role": SalesforceAccountContactRole(ctx),
			"salesforce_contract":             SalesforceContract(ctx),
			"salesforce_opportunity":          SalesforceOpportunity(ctx),
			"salesforce_order":                SalesforceOrder(ctx),
			"salesforce_user":                 SalesforceUser(ctx),
		},
	}

	return p
}
