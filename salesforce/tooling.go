package salesforce

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/simpleforce/simpleforce"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Some Salesforce configuration objects (for example, NamedCredential,
// RemoteProxy, CspTrustedSite, CorsWhitelistEntry) are only queryable through
// the Tooling API, not the standard SOQL query endpoint. simpleforce exposes a
// shared, mutable Tooling() flag on the cached client, which is unsafe under the
// concurrent table access the plugin performs. Instead we reuse the existing
// authenticated session (SID + instance URL) and call the Tooling query
// endpoint directly.

type toolingQueryResult struct {
	TotalSize      int                      `json:"totalSize"`
	Done           bool                     `json:"done"`
	NextRecordsURL string                   `json:"nextRecordsUrl"`
	Records        []map[string]interface{} `json:"records"`
}

func toolingAPIVersion(d *plugin.QueryData) string {
	config := GetConfig(d.Connection)
	if config.APIVersion != nil && *config.APIVersion != "" {
		return strings.TrimPrefix(*config.APIVersion, "v")
	}
	return simpleforce.DefaultAPIVersion
}

// queryTooling runs a SOQL query against the Tooling API using the existing
// simpleforce session without mutating the shared client's Tooling flag.
func queryTooling(ctx context.Context, client *simpleforce.Client, apiVersion, q string) (*toolingQueryResult, error) {
	loc := client.GetLoc()

	var reqURL string
	if strings.HasPrefix(q, "/services/data") {
		// q is a nextRecordsUrl returned by a previous page.
		reqURL = loc + q
	} else {
		reqURL = fmt.Sprintf("%s/services/data/v%s/tooling/query?q=%s", loc, apiVersion, url.PathEscape(q))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+client.GetSid())
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("tooling query failed (HTTP %d): %s", resp.StatusCode, string(body))
	}

	var result toolingQueryResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

//// LIST HYDRATE FUNCTION (Tooling API)

func listSalesforceToolingObjectsByTable(tableName string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("salesforce.listSalesforceToolingObjectsByTable", "connection error", err)
			return nil, err
		}
		if client == nil {
			return nil, fmt.Errorf("salesforce.listSalesforceToolingObjectsByTable: client_not_found, unable to query table %s because of invalid steampipe salesforce configuration", d.Table.Name)
		}

		query := generateQuery(d.Table.Columns, tableName)
		apiVersion := toolingAPIVersion(d)

		for {
			result, err := queryTooling(ctx, client, apiVersion, query)
			if err != nil {
				plugin.Logger(ctx).Error("salesforce.listSalesforceToolingObjectsByTable", "query error", err)
				return nil, err
			}

			for _, record := range result.Records {
				d.StreamListItem(ctx, record)
			}

			// Paging
			if result.Done || result.NextRecordsURL == "" {
				break
			}
			query = result.NextRecordsURL
		}

		return nil, nil
	}
}
