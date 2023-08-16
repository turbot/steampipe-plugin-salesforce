# Table: salesforce_permission_set

Represents a set of permissions that's used to grant more access to one or more users without changing their profile or reassigning profiles.
PermissionSet has a read-only child relationship with PermissionSetGroup. Permission set contains the aggregated permissions for the group.

**Note**: This table has one field for each permission with the pattern `permissions_permission_name`, e.g., `permissions_edit_task`. If true, users assigned to this permission set have the named permission. The number of fields varies depending on the permissions for the organization and license type.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_permission_set#show_permission_sets_created_in_last_30_days).

## Examples

### Basic info

```sql
select
  id,
  name,
  label,
  description,
  is_custom,
  created_date,
from
  salesforce_permission_set
```

### List non-custom permission sets

```sql
select
  id,
  name,
  label,
  description,
  is_custom,
  created_date,
from
  salesforce_permission_set
where
  not is_custom;
```

### List permission sets that contain the "Modify All Data" permission

```sql
select
  id,
  name,
  label,
  description,
  is_custom,
  created_date,
  permissions_modify_all_data
from
  salesforce_permission_set
where
  permissions_modify_all_data;
```

## API Native Examples

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "Name",
  "Label",
  "Description",
  "IsCustom",
  "CreatedDate"
from
  "PermissionSet";
```

### List non-custom permission sets (with API Native naming convention)

```sql
select
  "ID",
  "Name",
  "Label",
  "Description",
  "IsCustom",
  "CreatedDate"
from
  "PermissionSet"
where
  not "IsCustom";
```

### Show permission sets created in last 30 days

```sql
select
  "ID",
  "Name",
  "Label",
  "Description",
  "IsCustom",
  "CreatedDate"
from
  "PermissionSet"
where
  "CreatedDate" <= now() - interval '30' day;
```

### List permission sets where activation required

```sql
select
  "ID",
  "Name",
  "Label",
  "Description",
  "IsCustom",
  "CreatedDate"
from
  "PermissionSet"
where
  "HasActivationRequired";
```
