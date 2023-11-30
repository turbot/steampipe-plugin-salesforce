---
title: "Steampipe Table: salesforce_permission_set - Query Salesforce Permission Sets using SQL"
description: "Allows users to query Salesforce Permission Sets, providing insights into the permissions and access settings within a Salesforce organization."
---

# Table: salesforce_permission_set - Query Salesforce Permission Sets using SQL

Salesforce Permission Sets are a flexible and granular means of assigning permissions and access settings within a Salesforce organization. They allow you to extend user's functional access without changing their roles or profiles. Permission Sets can be used to grant additional permissions and access settings to users, on top of what their profile provides.

## Table Usage Guide

The `salesforce_permission_set` table provides insights into Salesforce Permission Sets within a Salesforce organization. As a Salesforce administrator, explore permission set-specific details through this table, including assigned permissions, access settings, and associated metadata. Utilize it to uncover information about permission sets, such as those with specific user access, the permissions associated with each set, and the verification of access settings.

## Examples

### Basic info
Explore which permissions are custom-made within your Salesforce environment. This can help you better manage user access and understand the creation timeline of these permissions.

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
Explore which permission sets in your Salesforce environment are not custom-made. This helps to understand the default permissions given and aids in maintaining security standards.

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
Explore which permission sets in Salesforce have been granted the capability to modify all data. This is useful to identify potential security risks and ensure only appropriate roles have such extensive permissions.

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

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore the basic details of your permission sets to understand their custom status and creation date. This can help you manage and organize your permission sets effectively.

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
Discover the segments that consist of non-custom permission sets. This can be useful in understanding the default sets provided by the platform and to ensure they align with your organization's security guidelines.

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
Discover the segments that have recently updated their access rights by focusing on those that have made changes within the past month. This is useful for maintaining security and ensuring that permissions are up-to-date.

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
Discover the segments that require activation within your permission sets. This can help you identify areas where additional steps may be needed before the permission set can be used, improving your system's security and compliance.

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