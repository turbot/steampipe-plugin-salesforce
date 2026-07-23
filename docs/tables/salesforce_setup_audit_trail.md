---
title: "Steampipe Table: salesforce_setup_audit_trail - Query Salesforce Setup Audit Trail using SQL"
description: "Allows users to query the Salesforce setup audit trail, providing insights into configuration changes made in the Setup area of the organization."
---

# Table: salesforce_setup_audit_trail - Query Salesforce Setup Audit Trail using SQL

The Salesforce setup audit trail records administrative and configuration changes made in the Setup area, including who made each change and when. It is a key source for change auditing and compliance.

## Table Usage Guide

The `salesforce_setup_audit_trail` table provides insights into configuration changes in your Salesforce organization. As a security or compliance administrator, use this table to review recent Setup changes, attribute them to users, and detect unexpected modifications.

## Examples

### Basic info
Explore recent configuration changes made in Setup.

```sql+postgres
select
  id,
  action,
  section,
  created_by_id,
  created_date
from
  salesforce_setup_audit_trail
order by
  created_date desc;
```

```sql+sqlite
select
  id,
  action,
  section,
  created_by_id,
  created_date
from
  salesforce_setup_audit_trail
order by
  created_date desc;
```

### List changes made in the last 7 days
Focus on the most recent configuration activity for review.

```sql+postgres
select
  action,
  section,
  display,
  created_by_id,
  created_date
from
  salesforce_setup_audit_trail
where
  created_date >= now() - interval '7' day;
```

```sql+sqlite
select
  action,
  section,
  display,
  created_by_id,
  created_date
from
  salesforce_setup_audit_trail
where
  created_date >= datetime('now', '-7 day');
```
