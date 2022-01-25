# Table: {metric_name}

Query data from the object called `{object_name}`. A table is automatically created to represent each Object.

For the connection configuration:

```hcl
connection "salesforce" {
  plugin = "salesforce"
  url      = "https://abc-dev-ed.my.salesforce.com"
  user     = "abcde@xyz.com"
  password = "dummy@9ATbuu"
  token    = "Y3NPBGO8qRV0RYlGPwldsmxv7"
  tables   = ["Case", "Campaign", "CustomApp__c"]
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
  tables   = ["Case", "Campaign", "CustomApp__c"]
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
> .inspect salesforce.salesforce_campaign
+--------------------------------+--------------------------+------------------------------------------+
| column                         | type                     | description                              |
+--------------------------------+--------------------------+------------------------------------------+
| actual_cost                    | double precision         | The Actual Cost in Campaign.             |
| amount_all_opportunities       | double precision         | The Value Opportunities in Campaign.     |
| amount_won_opportunities       | double precision         | The Value Won Opportunities in Campaign. |
| budgeted_cost                  | double precision         | The Budgeted Cost in Campaign.           |
| campaign_member_record_type_id | text                     | The Record Type ID.                      |
| created_by_id                  | text                     | The Created By ID.                       |
+--------------------------------+--------------------------+------------------------------------------+
```

### Get all the values for salesforce_campaign

```sql
select
  *
from
  salesforce_campaign
```

### Get current values for a object with specific id

```sql
select
  *
from
  salesforce_campaign
where
  id = '7015j0000019GVgAAM'
```
