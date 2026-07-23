---
title: "Steampipe Table: salesforce_named_credential - Query Salesforce Named Credentials using SQL"
description: "Allows users to query Salesforce Named Credentials, providing insights into the callout endpoints and authentication settings used by integrations."
---

# Table: salesforce_named_credential - Query Salesforce Named Credentials using SQL

A Salesforce Named Credential specifies the URL of a callout endpoint and its authentication settings, so Apex, flows, and external services can make authenticated callouts without hardcoding credentials. This table is queried through the Salesforce Tooling API.

## Table Usage Guide

The `salesforce_named_credential` table provides insights into how your Salesforce org authenticates to external systems. As a security or integration administrator, use this table to inventory outbound integration endpoints, review their authentication protocol and principal type, and spot endpoints using no authentication.

**Important Notes**
- This table is backed by the Salesforce Tooling API and requires a user with permission to view setup and configuration.
- Secret fields (password, OAuth tokens, AWS secrets, etc.) are intentionally not exposed as columns.
- This table does not support the `api_native` naming convention for columns. Because it cannot use the Describe-driven dynamic columns, its column names remain snake case even when `naming_convention` is set to `api_native`.

## Examples

### Basic info
Explore the named credentials configured in your organization.

```sql+postgres
select
  id,
  developer_name,
  endpoint,
  principal_type,
  protocol
from
  salesforce_named_credential;
```

```sql+sqlite
select
  id,
  developer_name,
  endpoint,
  principal_type,
  protocol
from
  salesforce_named_credential;
```

### Find named credentials that use no authentication
Identify callout endpoints configured without authentication, which may warrant a security review.

```sql+postgres
select
  developer_name,
  endpoint,
  protocol
from
  salesforce_named_credential
where
  protocol = 'NoAuthentication';
```

```sql+sqlite
select
  developer_name,
  endpoint,
  protocol
from
  salesforce_named_credential
where
  protocol = 'NoAuthentication';
```
