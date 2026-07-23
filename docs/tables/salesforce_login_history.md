---
title: "Steampipe Table: salesforce_login_history - Query Salesforce Login History using SQL"
description: "Allows users to query Salesforce login history, providing insights into successful and failed login attempts across the organization."
---

# Table: salesforce_login_history - Query Salesforce Login History using SQL

Salesforce login history records successful and failed login attempts for an organization and its enabled portals, including the user, source IP, application, and login type. Login history is retained for the last six months.

## Table Usage Guide

The `salesforce_login_history` table provides insights into authentication activity in your Salesforce organization. As a security administrator, use this table to investigate suspicious logins, track integration logins by application, and identify failed login attempts by source IP.

## Examples

### Basic info
Explore recent login attempts and their outcomes.

```sql+postgres
select
  id,
  user_id,
  login_time,
  status,
  source_ip,
  application
from
  salesforce_login_history
order by
  login_time desc;
```

```sql+sqlite
select
  id,
  user_id,
  login_time,
  status,
  source_ip,
  application
from
  salesforce_login_history
order by
  login_time desc;
```

### List failed login attempts
Identify login attempts that did not succeed, useful for detecting brute-force activity.

```sql+postgres
select
  user_id,
  login_time,
  status,
  source_ip
from
  salesforce_login_history
where
  status != 'Success';
```

```sql+sqlite
select
  user_id,
  login_time,
  status,
  source_ip
from
  salesforce_login_history
where
  status != 'Success';
```
