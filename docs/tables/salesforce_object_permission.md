---
title: "Steampipe Table: salesforce_object_permission - Query Salesforce Object Permissions using SQL"
description: "Allows users to query Salesforce Object Permissions, specifically the permissions that users have on Salesforce objects."
---

# Table: salesforce_object_permission - Query Salesforce Object Permissions using SQL

Salesforce Object Permissions is a feature within Salesforce that allows you to control the level of access that users have to Salesforce objects. It provides a way to set up and manage permissions for various Salesforce objects, including accounts, contacts, leads, and opportunities. Salesforce Object Permissions helps you maintain the security and integrity of your Salesforce data by ensuring that users only have the appropriate level of access to Salesforce objects.

## Table Usage Guide

The `salesforce_object_permission` table provides insights into the permissions that users have on Salesforce objects. As a Salesforce administrator, explore permission-specific details through this table, including the Salesforce object that the permission applies to, the type of permission, and the user or profile that the permission is associated with. Utilize it to uncover information about permissions, such as those that allow users to view, create, edit, or delete Salesforce objects.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Explore which Salesforce object permissions allow for modification or viewing of all records. This is beneficial for assessing user access rights and ensuring appropriate security measures are in place.

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
Determine the areas in which specific permissions are granted for transferring leads. This query is useful for assessing user permissions and ensuring appropriate access control within your Salesforce environment.

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
Determine areas in which users have comprehensive permissions, such as the ability to view or modify all records, to understand potential security risks and compliance issues in your system.

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
Determine the areas in which users have delete permissions to understand potential security risks or areas for access management improvements. This query is useful for administrators looking to optimize user roles and permissions.

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
Explore which Salesforce objects a user has read permissions for, allowing you to understand and manage access rights effectively. This can be particularly useful for auditing user permissions or troubleshooting access issues.

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