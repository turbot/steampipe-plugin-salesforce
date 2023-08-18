# Table: salesforce_object_permission

Represents the enabled object permissions for the parent PermissionSet.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_object_permission#api_native_examples).

## Examples

### Basic info

```sql
select
  id,
  parent_id,
  sobject_type,
  permissions_modify_all_records,
  permissions_view_all_records
from
  salesforce_object_permission
order by
  sobject_type;
```

### List permission sets with "Transfer Leads" permission on "Lead" object

```sql
select
  sop.id,
  sop.parent_id,
  sps.name,
  sps.permissions_transfer_any_lead,
  sop.sobject_type,
  sop.permissions_read,
  sop.permissions_create
from
  salesforce_object_permission sop,
  salesforce_permission_set sps
where
  sobject_type = 'Lead' and
  sps.id = sop.parent_id;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "ParentID",
  "SobjectType",
  "PermissionsModifyAllRecords",
  "PermissionsViewAllRecords"
from
  "ObjectPermission"
order by
  "SobjectType";
```

### Show delete permissions

```sql
select
  "ID",
  "ParentID",
  "SobjectType",
  "PermissionsModifyAllRecords",
  "PermissionsViewAllRecords"
from
  "ObjectPermission"
where
  "PermissionsDelete";
```

### Show read permissions

```sql
select
  "ID",
  "ParentID",
  "SobjectType",
  "PermissionsModifyAllRecords",
  "PermissionsViewAllRecords"
from
  "ObjectPermission"
where
  "PermissionsRead";
```
