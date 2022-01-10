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

### List Active users

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

### List Standard users

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
