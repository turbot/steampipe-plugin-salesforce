# Table: salesforce_account_contact_role

Represents the role that a Contact plays on an Account.

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
  salesforce_account_contact_role;
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
  salesforce_account_contact_role
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
  salesforce_account_contact_role
where
  user_type = 'Standard';
```
