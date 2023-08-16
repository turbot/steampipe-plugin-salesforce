# Table: salesforce_user

Represents a user in salesforce organization.

If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select username, alias from salesforce_user` would become `select "Username", "Alias" from "User"`.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_user#api_native_examples).

## Examples

### Basic info

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
