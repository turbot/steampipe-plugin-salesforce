---
title: "Steampipe Table: salesforce_{object_name} - Query Salesforce {object_name} using SQL"
description: "Allows users to query {object_name} in Salesforce, providing insights into Salesforce objects and their related data."
---

# Table: salesforce_{object_name} - Query Salesforce {object_name} using SQL

Salesforce is a customer relationship management solution that offers applications for small, midsize, and enterprise organizations, with a focus on sales and support. The Salesforce platform allows users to manage all interactions with their customers and prospects, including marketing automation, sales and customer service, analytics, and application development. {object_name} in Salesforce represents a specific type of object that can be created, manipulated, and deleted.

## Table Usage Guide

The `salesforce_{object_name}` table provides insights into {object_name} within Salesforce. As a Salesforce administrator or developer, you can explore specific details through this table, including object properties, relationships, and associated metadata. Use it to uncover information about {object_name}, such as their properties, relationships with other objects, and the manipulation of these objects.

## Examples

### Inspect the table structure

Assuming your connection configuration is:

```hcl
connection "salesforce" {
  plugin    = "salesforce"
  url       = "https://my-dev-env.my.salesforce.com"
  username  = "user@example.com"
  password  = "MyPassword"
  token     = "MyToken"
  client_id = "MyClientID"
  objects   = ["Case", "Campaign", "CustomApp__c"]
}
```

List all tables with:

```sh
.inspect salesforce
+---------------------------------+---------------------------------------------------------+
| table                           | description                                             |
+---------------------------------+---------------------------------------------------------+
| salesforce_account_contact_role | Represents the role that a Contact plays on an Account. |
| salesforce_campaign             | Represents Saleforce object Campaign.                   |
| salesforce_case                 | Represents Salesforce object Case.                      |
| salesforce_custom_app__c        | Represents Salesforce object CustomApp__c.              |
+---------------------------------+---------------------------------------------------------+
```

To get details of a specific object table, inspect it by name:

```sh
.inspect salesforce_custom_app__c
+---------------------+--------------------------+-------------------------+
| column              | type                     | description             |
+---------------------+--------------------------+-------------------------+
| created_by_id       | text                     | ID of app creator.      |
| created_date        | timestamp with time zone | Created date.           |
| id                  | text                     | App record ID.          |
| is_deleted          | boolean                  | True if app is deleted. |
| last_modified_by_id | text                     | ID of last modifier.    |
| last_modified_date  | timestamp with time zone | Last modified date.     |
| name                | text                     | App name.               |
| owner_id            | text                     | Owner ID.               |
| system_modstamp     | timestamp with time zone | System Modstamp.        |
+---------------------+--------------------------+-------------------------+
```

### Get all values from salesforce_custom_app\_\_c
Explore all the custom applications in your Salesforce environment to gain insights into their configurations and usage. This can be particularly useful for auditing purposes or when planning system upgrades or changes.

```sql
select
  *
from
  salesforce_custom_app__c;
```

### List custom apps added in the last 24 hours
Explore which custom apps have been added recently. This is useful for maintaining an up-to-date inventory and tracking the evolution of your application landscape.

```sql
select
  id,
  name,
  owner_id
from
  salesforce_custom_app__c
where
  created_date = now() - interval '24 hrs';
```

### Get details for a custom app by ID
Determine the details of a specific custom application in your Salesforce environment by using its unique identifier. This can be useful for obtaining a comprehensive view of an application's settings or status.

```sql
select
  *
from
  salesforce_custom_app__c
where
  id = '7015j0000019GVgAAM';
```

If `naming_convention` is set to `api_native`, the plugin will use the Salesforce naming conventions. Table and column names will have mixed case and table names will not start with `salesforce_`.

For instance, if the connection configuration is:

```hcl
connection "salesforce" {
  plugin            = "salesforce"
  url               = "https://my-dev-env.my.salesforce.com"
  username          = "user@example.com"
  password          = "MyPassword"
  token             = "MyToken"
  client_id         = "MyClientID"
  objects           = ["Case", "Campaign", "CustomApp__c"]
  naming_convention = "api_native"
}
```

This plugin will automatically create tables called `Case`, `Campaign` and `CustomApp__c`:

```sh
select "ID", "OwnerID", "IsDeleted" from "CustomApp__c"
+--------------------+--------------------+------------+
| ID                 | OwnerID            | IsDeleted  |
+--------------------+--------------------+------------+
| a005j000006vluAAAQ | 0055j000003GU6IAAW | false      |
| a005j000006vluBAAQ | 0055j000003GU6IAAW | false      |
| a005j000006vluFAAQ | 0055j000003GU6IAAW | false      |
+--------------------+--------------------+------------+
```

## API Native Examples

List all tables with:

```sh
.inspect salesforce
+---------------------------------+---------------------------------------------------------+
| table                           | description                                             |
+---------------------------------+---------------------------------------------------------+
| AccountContactRole              | Represents the role that a Contact plays on an Account. |
| Campaign                        | Represents Saleforce object Campaign.                   |
| Case                            | Represents Salesforce object Case.                      |
| CustomApp__c                    | Represents Salesforce object CustomApp__c.              |
+---------------------------------+---------------------------------------------------------+
```

To get details of a specific object table, inspect it by name:

```sh
.inspect CustomApp__c
+---------------------+--------------------------+-------------------------+
| column              | type                     | description             |
+---------------------+--------------------------+-------------------------+
| CreatedByID         | text                     | ID of app creator.      |
| CreatedDate         | timestamp with time zone | Created date.           |
| ID                  | text                     | App record ID.          |
| IsDeleted           | boolean                  | True if app is deleted. |
| LastModifiedByID    | text                     | ID of last modifier.    |
| LastModifiedDate    | timestamp with time zone | Last modified date.     |
| Name                | text                     | App name.               |
| OwnerID             | text                     | Owner ID.               |
| SystemModstamp      | timestamp with time zone | System Modstamp.        |
+---------------------+--------------------------+-------------------------+
```

### List custom apps added in the last 24 hours (with API Native naming convention)
Discover the custom applications added within the past day. This can be useful for monitoring recent additions and their owners for management or security purposes.

```sql
select
  "ID",
  "Name",
  "OwnerID"
from
  "CustomApp__c"
where
  "CreatedDate" = now() - interval '24 hrs';
```