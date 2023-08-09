# Table: salesforce_account_contact_role

Represents the role that a Contact plays on an Account.

If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select id, account_id from salesforce_account_contact_role` would become `select "ID", "AccountID" from "AccountContactRole"`.

Additional examples can be found [below](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account_contact_role#show_approver_account_contact_roles).

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
