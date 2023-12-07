---
title: "Steampipe Table: salesforce_user - Query Salesforce Users using SQL"
description: "Allows users to query Users in Salesforce, providing detailed information about each user's profile, role, and status."
---

# Table: salesforce_user - Query Salesforce Users using SQL

Salesforce Users are the individuals who have log in access to the Salesforce organization. They are associated with a unique username and profile that determines the level of access they have within the organization. User information includes details about the user's name, email, profile, role, and status among other attributes.

## Table Usage Guide

The `salesforce_user` table provides insights into individual users within Salesforce. As a Salesforce administrator or auditor, explore user-specific details through this table, including their profile, role, and status. Utilize it to uncover information about users, such as their level of access, the roles they are assigned, and their activity status.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).
- If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select username, alias from salesforce_user` would become `select "Username", "Alias" from "User"`.

## Examples

### Basic info
Discover the segments of users in your Salesforce platform, gaining insights into their activity status and last login date. This can be beneficial for assessing user engagement and identifying inactive users.

```sql+postgres
select
  username,
  alias,
  user_type,
  is_active,
  last_login_date
from
  salesforce_user;
```

```sql+sqlite
select
  username,
  alias,
  user_type,
  is_active,
  last_login_date
from
  salesforce_user;
```

### List active users
Explore which users are currently active in your Salesforce environment. This is useful for understanding user engagement and identifying any potential inactive accounts.

```sql+postgres
select
  username,
  alias,
  user_type,
  is_active,
  last_login_date
from
  salesforce_user
where
  is_active;
```

```sql+sqlite
select
  username,
  alias,
  user_type,
  is_active,
  last_login_date
from
  salesforce_user
where
  is_active;
```

### List standard users
Explore which standard users are active within your Salesforce platform. This query is particularly useful in understanding user activity and identifying any inactive accounts that may need attention.

```sql+postgres
select
  username,
  alias,
  user_type,
  is_active,
  last_login_date
from
  salesforce_user
where
  user_type = 'Standard';
```

```sql+sqlite
select
  username,
  alias,
  user_type,
  is_active,
  last_login_date
from
  salesforce_user
where
  user_type = 'Standard';
```

### List user designated as forecast managers
Gain insights into your Salesforce users who have been enabled to manage forecasts. This allows you to better understand who in your team has the authority to manipulate and oversee forecasting data.

```sql+postgres
select
  id,
  username,
  user_type,
  forecast_enabled
from
  salesforce_user
where
  forecast_enabled;
```

```sql+sqlite
select
  id,
  username,
  user_type,
  forecast_enabled
from
  salesforce_user
where
  forecast_enabled = 1;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore which users are active and when they last logged in to effectively manage user access and understand usage patterns. This is useful for maintaining security, optimizing user experience, and making data-driven decisions.

```sql+postgres
select
  "Username",
  "Alias",
  "UserType",
  "IsActive",
  "LastLoginDate"
from
  "User";
```

```sql+sqlite
select
  "Username",
  "Alias",
  "UserType",
  "IsActive",
  "LastLoginDate"
from
  "User";
```

### List active users (with API Native naming convention)
Explore which users are currently active and when they last logged in, helping to monitor user activity and understand usage patterns.

```sql+postgres
select
  "Username",
  "Alias",
  "UserType",
  "IsActive",
  "LastLoginDate"
from
  "User"
where
  "IsActive";
```

```sql+sqlite
select
  "Username",
  "Alias",
  "UserType",
  "IsActive",
  "LastLoginDate"
from
  "User"
where
  "IsActive" = 1;
```

### List guest users
Explore which users in your system are classified as 'Guests'. This is useful for identifying potential security risks or for auditing user access privileges.

```sql+postgres
select
  "Username",
  "Alias",
  "UserType",
  "IsActive",
  "LastLoginDate"
from
  "User"
where
  "UserType" = 'Guest';
```

```sql+sqlite
select
  "Username",
  "Alias",
  "UserType",
  "IsActive",
  "LastLoginDate"
from
  "User"
where
  "UserType" = 'Guest';
```

### List users who logged-in in last 30 days
Discover the segments of users who have been active within the past month. This is useful for understanding user engagement and activity patterns over time.

```sql+postgres
select
  "Username",
  "Alias",
  "UserType",
  "IsActive",
  "LastLoginDate"
from
  "User"
where
  "LastLoginDate" <= now() - interval '30' day;
```

```sql+sqlite
select
  "Username",
  "Alias",
  "UserType",
  "IsActive",
  "LastLoginDate"
from
  "User"
where
  "LastLoginDate" <= datetime('now', '-30 day');
```