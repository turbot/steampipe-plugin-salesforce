# Table: salesforce_permission_set_assignment

Represents the association between a user and permission set.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_permission_set_assignment#api_native_examples).

## Examples

### Basic info

```sql
select
  id,
  assignee_id,
  permission_set_group_id,
  permission_set_id
from
  salesforce_permission_set_assignment;
```

### List users with permission to "Modify All Data"

```sql
with sps as (
  select
    id,
    name,
    description,
    label
  from
    salesforce_permission_set
  where
    permissions_modify_all_data
),
spsa as (
  select
    *
  from
    salesforce_permission_set_assignment
)
select
  su.name as user_name,
  su.email as user_email,
  su.username as user_username,
  spsa.assignee_id as user_id,
  sps.id as permission_set_id,
  sps.name as permission_set_name,
  sps.description as permission_set_description,
  sps.label as permission_set_label
from
  sps,
  spsa,
  salesforce_user as su
where
  sps.id = spsa.permission_set_id
  and spsa.assignee_id = su.id;
```

## API Native Examples

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "AssigneeID",
  "PermissionSetGroupID",
  "PermissionSetID"
from
  "PermissionSetAssignment";
```

### List assignments which are not associated with any permission set groups

```sql
select
  "ID",
  "AssigneeID",
  "PermissionSetGroupID",
  "PermissionSetID"
from
  "PermissionSetAssignment"
where
  "PermissionSetGroupID" is null;
```
