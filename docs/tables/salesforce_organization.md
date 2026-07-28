---
title: "Steampipe Table: salesforce_organization - Query Salesforce Organization using SQL"
description: "Allows users to query the Salesforce organization record, providing insights into org-level configuration such as edition, instance, and locale settings."
---

# Table: salesforce_organization - Query Salesforce Organization using SQL

The Salesforce Organization record holds key configuration information about the org itself, including its edition, hosting instance, sandbox status, and default locale settings.

## Table Usage Guide

The `salesforce_organization` table provides insights into the top-level configuration of your Salesforce org. As an administrator, use this table to confirm the edition and instance, distinguish sandboxes from production, and review default locale and fiscal-year settings.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Explore the core configuration of your organization.

```sql+postgres
select
  id,
  name,
  organization_type,
  instance_name,
  is_sandbox
from
  salesforce_organization;
```

```sql+sqlite
select
  id,
  name,
  organization_type,
  instance_name,
  is_sandbox
from
  salesforce_organization;
```

### Check whether the org is a sandbox
Confirm the environment type before running changes.

```sql+postgres
select
  name,
  organization_type,
  is_sandbox
from
  salesforce_organization
where
  is_sandbox;
```

```sql+sqlite
select
  name,
  organization_type,
  is_sandbox
from
  salesforce_organization
where
  is_sandbox = 1;
```
