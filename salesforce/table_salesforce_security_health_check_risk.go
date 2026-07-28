package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

// SecurityHealthCheckRisks (Tooling API) reports each org security setting, its
// current value vs the Salesforce baseline, and a risk rating. It is the read-
// only path to session/password/login-IP posture — reachable with the "View
// Health Check" permission, without the write-level Metadata API.
func SalesforceSecurityHealthCheckRisk(ctx context.Context, config salesforceConfig) *plugin.Table {
	tableName := "SecurityHealthCheckRisks"
	return &plugin.Table{
		Name:        "salesforce_security_health_check_risk",
		Description: "Represents a single Security Health Check finding — an org security setting, its current value versus the Salesforce baseline, and its risk rating. Queried through the Salesforce Tooling API.",
		List: &plugin.ListConfig{
			Hydrate: listSalesforceToolingObjectsByTable(tableName),
		},
		Columns: toolingColumns(config, []*plugin.Column{
			{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the organization in Salesforce.", Hydrate: getOrganizationId, Transform: transform.FromValue()},

			{Name: "durable_id", Type: proto.ColumnType_STRING, Description: "Stable identifier of the health-check risk row."},
			{Name: "setting_group", Type: proto.ColumnType_STRING, Description: "The group the setting belongs to, for example, PasswordPolicies or SessionSettings."},
			{Name: "setting", Type: proto.ColumnType_STRING, Description: "The name of the security setting."},
			{Name: "org_value", Type: proto.ColumnType_STRING, Description: "The setting's current value in this org (display form)."},
			{Name: "standard_value", Type: proto.ColumnType_STRING, Description: "The Salesforce baseline (recommended) value for the setting (display form)."},
			{Name: "risk_type", Type: proto.ColumnType_STRING, Description: "Risk rating of the org value versus the baseline: HIGH_RISK, MEDIUM_RISK, or MEETS_STANDARD."},
			{Name: "setting_risk_category", Type: proto.ColumnType_STRING, Description: "Risk category of the setting: HIGH_RISK, MEDIUM_RISK, LOW_RISK, or INFORMATIONAL."},
			{Name: "org_value_raw", Type: proto.ColumnType_STRING, Description: "The setting's current value in raw form."},
			{Name: "standard_value_raw", Type: proto.ColumnType_STRING, Description: "The Salesforce baseline value in raw form."},
		}),
	}
}
