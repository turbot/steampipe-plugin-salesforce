# Table: salesforce_account

Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).

## Examples

### Basic info

```sql
select
  name,
  industry,
  type,
  ownership,
  rating
from
  salesforce_account;
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