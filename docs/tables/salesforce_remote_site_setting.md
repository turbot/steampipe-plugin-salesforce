---
title: "Steampipe Table: salesforce_remote_site_setting - Query Salesforce Remote Site Settings using SQL"
description: "Allows users to query Salesforce Remote Site Settings, providing insights into the external sites Salesforce is authorized to make callouts to."
---

# Table: salesforce_remote_site_setting - Query Salesforce Remote Site Settings using SQL

A Salesforce Remote Site Setting authorizes Salesforce to make callouts to a specific external site. Before Apex, flows, or other server-side code can call an external URL, that URL's origin must be registered as a remote site. This table is queried through the Salesforce Tooling API.

## Table Usage Guide

The `salesforce_remote_site_setting` table provides insights into the outbound callout allowlist of your Salesforce org. As a security or integration administrator, use this table to inventory the external sites Salesforce can reach, confirm which are active, and detect sites that permit insecure (HTTP) callouts.

**Important Notes**
- This table is backed by the Salesforce Tooling API and requires a user with permission to view setup and configuration.

## Examples

### Basic info
Explore the remote site settings configured in your organization.

```sql+postgres
select
  id,
  site_name,
  endpoint_url,
  is_active
from
  salesforce_remote_site_setting;
```

```sql+sqlite
select
  id,
  site_name,
  endpoint_url,
  is_active
from
  salesforce_remote_site_setting;
```

### Find remote sites that allow insecure callouts
Identify active remote sites where protocol security is relaxed, allowing callouts over unsecured HTTP.

```sql+postgres
select
  site_name,
  endpoint_url,
  protocol_mismatch
from
  salesforce_remote_site_setting
where
  is_active
  and protocol_mismatch;
```

```sql+sqlite
select
  site_name,
  endpoint_url,
  protocol_mismatch
from
  salesforce_remote_site_setting
where
  is_active = 1
  and protocol_mismatch = 1;
```
