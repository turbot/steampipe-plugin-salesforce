# Table: salesforce_object_permission

Represents the enabled object permissions for the parent PermissionSet.

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

### List permission set with "Transfer Leads" permission, on "Lead" object

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
