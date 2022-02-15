# Table: salesforce_account_contact_role

Represents the role that a Contact plays on an Account.

## Examples

### Basic info

```sql
select
  id,
  account_id,
  contact_id,
  is_primary,
  role
from
  salesforce_account_contact_role;
```

### List primary account contact role

```sql
select
  id,
  account_id,
  contact_id,
  is_primary
from
  salesforce_account_contact_role
where
  is_primary;
```
