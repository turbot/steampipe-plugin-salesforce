# Table: salesforce_{object_name}

Query data from the object called `salesforce_{object_name}`, e.g., `salesforce_campaign`, `salesforce_custom_app__c`. A table is automatically created to represent each object in the `objects` argument.

For instance, if the connection configuration is:

```hcl
connection "salesforce" {
  plugin    = "salesforce"
  url       = "https://my-dev-env.my.salesforce.com"
  user      = "user@example.com"
  password  = "MyPassword"
  token     = "MyToken"
  client_id = "MyClientID"
  objects   = ["Case", "Campaign", "CustomApp__c"]
}
```

This plugin will automatically create tables called `salesforce_case`, `salesforce_campaign` and `salesforce_custom_app__c`:

```
select id, owner_id, is_deleted from salesforce_custom_app__c
+--------------------+--------------------+------------+
| id                 | owner_id           | is_deleted |
+--------------------+--------------------+------------+
| a005j000006vluAAAQ | 0055j000003GU6IAAW | false      |
| a005j000006vluBAAQ | 0055j000003GU6IAAW | false      |
| a005j000006vluFAAQ | 0055j000003GU6IAAW | false      |
+--------------------+--------------------+------------+
```

Please note that plugin initialization time will increase depending on the number of objects included in the `objects` argument.

**Note:** Salesforce custom object API names are always suffixed with `__c`, which is reflected in the table names as well.

## Examples

### Inspect the table structure

Assuming your connection configuration is:

```hcl
connection "salesforce" {
  plugin    = "salesforce"
  url       = "https://my-dev-env.my.salesforce.com"
  user      = "user@example.com"
  password  = "MyPassword"
  token     = "MyToken"
  client_id = "MyClientID"
  objects   = ["Case", "Campaign", "CustomApp__c"]
}
```

List all tables with:

```sql
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

```sql
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

```sql
select
  *
from
  salesforce_custom_app__c;
```

### List custom apps added in the last 24 hours

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

```sql
select
  *
from
  salesforce_custom_app__c
where
  id = '7015j0000019GVgAAM';
```
