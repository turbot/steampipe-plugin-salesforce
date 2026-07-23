---
title: "Steampipe Table: salesforce_external_data_source - Query Salesforce External Data Sources using SQL"
description: "Allows users to query Salesforce External Data Sources, providing insights into connections to data stored outside the Salesforce organization."
---

# Table: salesforce_external_data_source - Query Salesforce External Data Sources using SQL

A Salesforce External Data Source defines a connection to data stored outside the Salesforce org, for example, via Salesforce Connect. It specifies the endpoint, adapter type, and authentication used to reach the external system.

## Table Usage Guide

The `salesforce_external_data_source` table provides insights into the external systems your Salesforce organization integrates with. As an integration or platform administrator, use this table to review configured endpoints, authentication principal types, and whether each source is writable.

## Examples

### Basic info
Explore the external data sources configured in your organization.

```sql+postgres
select
  id,
  developer_name,
  master_label,
  type,
  endpoint
from
  salesforce_external_data_source;
```

```sql+sqlite
select
  id,
  developer_name,
  master_label,
  type,
  endpoint
from
  salesforce_external_data_source;
```

### List writable external data sources
Identify external data sources that allow write operations, which may need tighter access controls.

```sql+postgres
select
  id,
  developer_name,
  endpoint,
  is_writable
from
  salesforce_external_data_source
where
  is_writable;
```

```sql+sqlite
select
  id,
  developer_name,
  endpoint,
  is_writable
from
  salesforce_external_data_source
where
  is_writable = 1;
```
