---
title: "Steampipe Table: salesforce_permission_set_assignment - Query Salesforce Permission Set Assignments using SQL"
description: "Allows users to query Permission Set Assignments in Salesforce, particularly user permissions and access settings, providing insights into user access and security configurations."
---

# Table: salesforce_permission_set_assignment - Query Salesforce Permission Set Assignments using SQL

A Salesforce Permission Set Assignment is a functionality within Salesforce that allows administrators to assign specific permissions to users. It is a flexible and granular way to provide access beyond what is defined in the user's profile. Permission Set Assignments help organizations to manage user access and security settings effectively, ensuring that users have only the access they need.

## Table Usage Guide

The `salesforce_permission_set_assignment` table provides insights into Permission Set Assignments within Salesforce. As a Salesforce administrator, explore assignment-specific details through this table, including user permissions and access settings. Utilize it to uncover information about assignments, such as those with specific permissions, the relationships between users and their assigned permissions, and the verification of access policies.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Discover the segments that detail the assignment of permission sets in Salesforce. This can help identify instances where specific permissions are assigned, aiding in understanding user access control and ensuring necessary security measures are in place.

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
Explore which users have the authority to modify all data in your Salesforce environment. This is particularly useful for maintaining security and control over data integrity.

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

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore which permission sets are assigned to various users. This helps in understanding user access and permissions within your system, thereby aiding in effective access management.

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
Discover the segments that consist of assignments not linked to any permission set groups. This can be useful to identify potential issues in access control and rectify them for enhanced security management.

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