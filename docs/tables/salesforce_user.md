---
title: "Steampipe Table: salesforce_user - Query Salesforce Users using SQL"
description: "Allows users to query Users in Salesforce, providing detailed information about each user's profile, role, and status."
---

# Table: salesforce_user - Query Salesforce Users using SQL

Salesforce Users are the individuals who have log in access to the Salesforce organization. They are associated with a unique username and profile that determines the level of access they have within the organization. User information includes details about the user's name, email, profile, role, and status among other attributes.

## Table Usage Guide

The `salesforce_user` table provides insights into individual users within Salesforce. As a Salesforce administrator or auditor, explore user-specific details through this table, including their profile, role, and status. Utilize it to uncover information about users, such as their level of access, the roles they are assigned, and their activity status.

## Examples

### Basic info
Discover the segments of users in your Salesforce platform, gaining insights into their activity status and last login date. This can be beneficial for assessing user engagement and identifying inactive users.

```sql
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

```sql
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

```sql
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

```sql
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

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore which users are active and when they last logged in to effectively manage user access and understand usage patterns. This is useful for maintaining security, optimizing user experience, and making data-driven decisions.

```sql
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

```sql
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

### List guest users
Explore which users in your system are classified as 'Guests'. This is useful for identifying potential security risks or for auditing user access privileges.

```sql
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

```sql
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