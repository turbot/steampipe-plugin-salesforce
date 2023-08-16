# Table: salesforce_contact

Represents a contact, which is a person associated with an account.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_contact#list_deleted_contacts).

## Examples

### Basic info

```sql
select
  id,
  name,
  account_id,
  email,
  lead_source,
  title,
  department
from
  salesforce_contact;
```

## API Native Examples

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "Name",
  "AccountID",
  "Email",
  "Department"
from
  "Contact";
```

### List deleted contacts

```sql
select
  "ID",
  "Name",
  "AccountID",
  "Email",
  "Department"
from
  "Contact"
where
  "IsDeleted";
```

### Show contacts created in last 30 days

```sql
select
  "ID",
  "Name",
  "AccountID",
  "Email",
  "Department"
from
  "Contact"
where
  "CreatedDate" <= now() - interval '30' day;
```
