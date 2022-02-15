# Table: salesforce_user

Represents a user in salesforce organization.

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
