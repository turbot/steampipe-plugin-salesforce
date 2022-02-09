# Table: {object_name}

Query data from the object called `salesforce_{object_name}`, e.g., `salesforce_campaign`, `salesforce_custom_app__c`. A table is automatically created to represent each object in the `objects` argument.

For instance, if the the connection configuration is:

```hcl
connection "salesforce" {
  plugin = "salesforce"
  url      = "https://abc-dev-ed.my.salesforce.com"
  user     = "abcde@xyz.com"
  password = "dummy@9ATbuu"
  token    = "Y3NPBGO8qRV0RYlGPwldsmxv7"
  # List of salesforce tables to be generated. Salesforce Object API Name (refer https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm) should be used here."
  # Currently plugin supports a standard set of Salesforce tables by default in the plugin. If there is a need to query an standard object from, you can add standard object in the tables list. Steampipe will automatically generate respective table for the connection.
  objects   = ["Case", "Campaign", "CustomApp__c"]
}
```

This plugin will automatically create tables called `salesforce_case`, `salesforce_campaign` and `salesforce_custom_app__c`:

```
select * from salesforce_custom_app__c
+--------------------+--------------------+------------+------+---------------------------+--------------------+---------------------------+---------------------+---------------------------+
| id                 | owner_id           | is_deleted | name | created_date              | created_by_id      | last_modified_date        | last_modified_by_id | system_modstamp           |
+--------------------+--------------------+------------+------+---------------------------+--------------------+---------------------------+---------------------+---------------------------+
| a005j000006vluAAAQ | 0055j000003GU6IAAW | false      | AAP1 | 2022-01-13T14:18:28+05:30 | 0055j000003GU6IAAW | 2022-01-13T14:18:28+05:30 | 0055j000003GU6IAAW  | 2022-01-13T14:18:28+05:30 |
| a005j000006vluBAAQ | 0055j000003GU6IAAW | false      | APP2 | 2022-01-13T14:19:17+05:30 | 0055j000003GU6IAAW | 2022-01-13T14:19:17+05:30 | 0055j000003GU6IAAW  | 2022-01-13T14:19:17+05:30 |
| a005j000006vluFAAQ | 0055j000003GU6IAAW | false      | APP1 | 2022-01-13T14:18:58+05:30 | 0055j000003GU6IAAW | 2022-01-13T14:18:58+05:30 | 0055j000003GU6IAAW  | 2022-01-13T14:18:58+05:30 |
+--------------------+--------------------+------------+------+---------------------------+--------------------+---------------------------+---------------------+---------------------------+
```

However, please note that this could be slow depending on how many objects are in your environment.
**NOTE:** Salesforce custom objects and fields are always suffixed with `__c`

## Examples

### Inspect the table structure

Assuming your connection configuration is:

```hcl
connection "salesforce" {
  plugin = "salesforce"
  url      = "https://abc-dev-ed.my.salesforce.com"
  user     = "abcde@xyz.com"
  password = "dummy@9ATbuu"
  token    = "Y3NPBGO8qRV0RYlGPwldsmxv7"
  objects   = ["Case", "Campaign", "CustomApp__c"]
}
```

List all tables with:

```sql
.inspect salesforce
+-------------------------------------+--------------------------------------------------------+
| table                               | description                                            |
+-------------------------------------+--------------------------------------------------------+
| ...                                 | ...                                                    |
| salesforce_account_contact_role     | Represents the role that a Contact plays on an Account.|
| ...                                 | ...                                                    |
| salesforce_campaign                 | Salesforce Campaign.                                   |
| salesforce_case                     | Salesforce Case.                                       |
| salesforce_custom_app__c            | Salesforce Custom APP.                                 |
| ...                                 | ...                                                    |
+-------------------------------------+--------------------------------------------------------+
```

To get details of a specific object table, inspect it by name:

```sql
> .inspect salesforce_custom_app__c
+---------------------+--------------------------+--------------------------+
| column              | type                     | description              |
+---------------------+--------------------------+--------------------------+
| created_by_id       | text                     | The Created By ID.       |
| created_date        | timestamp with time zone | The Created Date.        |
| id                  | text                     | The Record ID.           |
| is_deleted          | boolean                  | The Deleted.             |
| last_modified_by_id | text                     | The Last Modified By ID. |
| last_modified_date  | timestamp with time zone | The Last Modified Date.  |
| name                | text                     | The Custom APP Name.     |
| owner_id            | text                     | The Owner ID.            |
| system_modstamp     | timestamp with time zone | The System Modstamp.     |
+---------------------+--------------------------+--------------------------+
```

### Get all the values from salesforce_custom_app\_\_c

```sql
select
  *
from
  salesforce_custom_app__c;
```

### List custom apps added in last 24 hrs

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
