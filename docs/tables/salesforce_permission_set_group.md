---
title: "Steampipe Table: salesforce_permission_set_group - Query Salesforce Permission Set Groups using SQL"
description: "Allows users to query Salesforce Permission Set Groups, providing insights into bundles of permission sets assigned to users together."
---

# Table: salesforce_permission_set_group - Query Salesforce Permission Set Groups using SQL

A Salesforce Permission Set Group bundles multiple permission sets so they can be assigned to users together. A user assigned the group receives the combined permissions of all its permission sets.

## Table Usage Guide

The `salesforce_permission_set_group` table provides insights into the permission set groups configured in your Salesforce organization. As a security administrator, use this table to review group definitions, their calculation status, and creation metadata.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Explore the permission set groups configured in your organization.

```sql+postgres
select
  id,
  developer_name,
  master_label,
  status,
  created_date
from
  salesforce_permission_set_group;
```

```sql+sqlite
select
  id,
  developer_name,
  master_label,
  status,
  created_date
from
  salesforce_permission_set_group;
```

### List permission set groups that need recalculation
Identify groups whose combined permissions are outdated and need recalculation.

```sql+postgres
select
  developer_name,
  master_label,
  status
from
  salesforce_permission_set_group
where
  status != 'Updated';
```

```sql+sqlite
select
  developer_name,
  master_label,
  status
from
  salesforce_permission_set_group
where
  status != 'Updated';
```
