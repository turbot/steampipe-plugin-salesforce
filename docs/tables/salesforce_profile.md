---
title: "Steampipe Table: salesforce_profile - Query Salesforce Profiles using SQL"
description: "Allows users to query Salesforce Profiles, providing insights into the permission and access-setting baselines assigned to users."
---

# Table: salesforce_profile - Query Salesforce Profiles using SQL

A Salesforce Profile defines a baseline set of permissions and access settings that control what a user can do. Every user is assigned exactly one profile, which is often extended with permission sets.

## Table Usage Guide

The `salesforce_profile` table provides insights into the profiles configured in your Salesforce organization. As a security administrator, use this table to review profiles, their associated user licenses, and user types when auditing access.

## Examples

### Basic info
Explore the profiles configured in your organization.

```sql+postgres
select
  id,
  name,
  user_type,
  user_license_id,
  created_date
from
  salesforce_profile;
```

```sql+sqlite
select
  id,
  name,
  user_type,
  user_license_id,
  created_date
from
  salesforce_profile;
```

### List profiles by user type
Group profiles by the category of user license they apply to.

```sql+postgres
select
  user_type,
  count(*) as profile_count
from
  salesforce_profile
group by
  user_type;
```

```sql+sqlite
select
  user_type,
  count(*) as profile_count
from
  salesforce_profile
group by
  user_type;
```
