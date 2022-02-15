# Table: salesforce_permission_set

Represents a set of permissions that's used to grant more access to one or more users without changing their profile or reassigning profiles.
PermissionSet has a read-only child relationship with PermissionSetGroup. Permission set contains the aggregated permissions for the group.

**Note**: This table has one field for each permission with the pattern `permissions_permission_name`, e.g., `permissions_edit_task`. If true, users assigned to this permission set have the named permission. The number of fields varies depending on the permissions for the organization and license type.

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
