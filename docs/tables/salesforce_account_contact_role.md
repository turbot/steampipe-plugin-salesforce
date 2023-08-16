# Table: salesforce_account_contact_role

Represents the role that a Contact plays on an Account.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account_contact_role#api_native_examples).

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

## API Native Examples

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "AccountID",
  "ContactID",
  "IsPrimary"
from
  "AccountContactRole";
```

### List primary account contact role (with API Native naming convention)

```sql
select
  "ID",
  "AccountID",
  "ContactID",
  "IsPrimary"
from
  "AccountContactRole"
where
  "IsPrimary";
```

### Show approver account contact roles

```sql
select
  "ID",
  "AccountID",
  "ContactID",
  "IsPrimary",
  "Role"
from
  "AccountContactRole"
where
  "Role" = 'Approver';
```

### Show account contact roles created in last 30 days

```sql
select
  "ID",
  "AccountID",
  "ContactID",
  "IsPrimary",
  "Role"
from
  "AccountContactRole"
where
  "CreatedDate" <= now() - interval '30' day;
```
