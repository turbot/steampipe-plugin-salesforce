---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/salesforce.svg"
brand_color: "#00A1E0"
display_name: "Salesforce"
short_name: "salesforce"
description: "Steampipe plugin to query accounts, opportunities, users and more from your Salesforce instance."
og_description: "Query Salesforce with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/salesforce-social-graphic.png"
---

# Salesforce + Steampipe

[Salesforce](https://www.salesforce.com/) is a customer relationship management (CRM) platform.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List won opportunities:

```sql
select
  name,
  amount,
  close_date
from
  salesforce_opportunity
where
  is_won;
```

```
+-------------------------------------+--------+---------------------------+
| name                                | amount | close_date                |
+-------------------------------------+--------+---------------------------+
| GenePoint Standby Generator         | 85000  | 2021-10-23T05:30:00+05:30 |
| GenePoint SLA                       | 30000  | 2021-12-16T05:30:00+05:30 |
| Express Logistics Standby Generator | 220000 | 2021-09-15T05:30:00+05:30 |
+-------------------------------------+------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/salesforce/tables)**

## Get started

### Install

Download and install the latest salesforce plugin:

```bash
steampipe plugin install salesforce
```

### Configuration

Installing the latest salesforce plugin will create a config file (`~/.steampipe/config/salesforce.spc`) with a single connection named `salesforce`:

```hcl
connection "salesforce" {
  plugin = "salesforce"

  # Salesforce instance URL, e.g., "https://na01.salesforce.com/"
  # url = "https://na01.salesforce.com/"

  # Salesforce account name
  # username = "user@example.com"

  # Salesforce account password
  # password = "Dummy@~Password"

  # The Salesforce security token is only required If the client's IP address is not added to the organization's list of trusted IPs
  # https://help.salesforce.com/s/articleView?id=sf.security_networkaccess.htm&type=5
  # token = "ABO5C3PNqOP0BHsPFakeToken"

  # Salesforce client ID of the connected app
  # client_id = "3MVG99E3Ry5mh4z_FakeID"

  # List of Salesforce object names to generate additional tables for
  # This argument only accepts exact Salesforce standard and custom object names, e.g., AccountBrand, OpportunityStage, CustomApp__c
  # For a full list of standard object names, please see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm)
  # All custom object names should end in "__c", following Salesforce object naming standards
  # objects = ["AccountBrand", "OpportunityStage", "CustomApp__c"]

  # Salesforce API version to connect to
  # api_version = "43.0"

  # The naming_convention allows users to control the naming format for tables and columns in the plugin. Below are the supported values:
  # api_native: If set to this value, the plugin will use the native format for table names, meaning there will be no salesforce_ prefix, and the table and column names will remain as they are in Salesforce.
  # snake_case: If the user does not specify any value or sets it to this value, the plugin will use snake case for table and column names and table names will have a salesforce_ prefix.
  # naming_convention = "snake_case"
}
```

### Credentials

- [Create your connected application](https://trailhead.salesforce.com/en/content/learn/projects/build-a-connected-app-for-api-integration/create-a-connected-app)
- Configure basic [connected application settings](https://help.salesforce.com/s/articleView?id=sf.connected_app_create_basics.htm&type=5)
- Reset your [security token](https://help.salesforce.com/articleView?id=user_security_token.htm&type=5), which is required if you are connecting from an IP address outside your company's trusted IP range

## Custom Fields

Salesforce supports the addition of [custom fields](https://help.salesforce.com/s/articleView?id=sf.adding_fields.htm&type=5) to standard objects.

If you have set up Salesforce credentials correctly in the Steampipe configuration, Steampipe will generate the tables schema with all the custom fields along with standard object fields dynamically.

For instance, if the `Account` object in my Salesforce account has a custom field with the label `Priority` and the API name `Priority__c`, the table schema will be generated as:

```sh
.inspect salesforce_account
+-----------------------+--------------------------+-------------------------------------------------------------+
| column         | type | description                                                                            |
+-----------------------+---------------------------+------------------------------------------------------------+
| account_number | text | The Account Number.                                                                    |
| account_source | text | The source of the account record. For example, Advertisement, Data.com, or Trade Show. |
| priority__c    | text | The account's priority.                                                                |
+----------------+------+----------------------------------------------------------------------------------------+
```

The custom field `priority__c` column can then be queried like other columns:

```sql
select
  account_number,
  priority__c
from
  salesforce_account;
```

**Note:** Salesforce custom field names are always suffixed with `__c`, which is reflected in the column names as well.

## Custom Objects

Salesforce also supports creating [custom objects](https://help.salesforce.com/s/articleView?id=sf.dev_objectcreate_task_lex.htm&type=5) to track and store data that's unique to your organization.

Steampipe will create table schemas for all custom objects set in the `objects` argument.

For instance, if my connection configuration is:

```hcl
connection "salesforce" {
  plugin    = "salesforce"
  url       = "https://my-dev-env.my.salesforce.com"
  username  = "user@example.com"
  password  = "MyPassword"
  token     = "MyToken"
  client_id = "MyClientID"
  objects   = ["CustomApp__c", "OtherCustomApp__c"]
}
```

Steampile will automatically create two tables, `salesforce_custom_app__c` and `salesforce_other_custom_app__c`, which can then be inspected and queried like other tables:

```sh
.inspect salesforce
+---------------------------------+---------------------------------------------------------+
| table                           | description                                             |
+---------------------------------+---------------------------------------------------------+
| salesforce_account_contact_role | Represents the role that a Contact plays on an Account. |
| salesforce_custom_app__c        | Represents Salesforce object CustomApp__c.              |
| salesforce_other_custom_app__c  | Represents Salesforce object OtherCustomApp__c.         |
+---------------------------------+---------------------------------------------------------+
```

To get details of a specific custom object table, inspect it by name:

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

This table can also be queried like other tables:

```sql
select
  *
from
  salesforce_custom_app__c;
```

**Note:** Salesforce custom object names are always suffixed with `__c`, which is reflected in the table names as well.

## Naming Convention

The `naming_convention` configuration argument allows you to control the naming format for tables and columns in the plugin. Below are the supported values:

### API Native

If `naming_convention` is set to `api_native`, the plugin will use the Salesforce naming conventions. Table and column names will have mixed case and table names will not start with `salesforce_`.

For example:

```sql
select
  "Id",
  "WhoCount",
  "WhatCount",
  "Subject",
  "IsAllDayEvent"
from
  "Event";
```

```
+---------------------+----------+-------------+---------+---------------+
| ID                  | WhoCount |  WhatCount  | Subject | IsAllDayEvent |
+----------------------------------------------+-------------------------+
| 00U2t0000000Mw3dEAD | 0        |  0          | test    | false         |
+---------------------+----------+-----------------------+---------------+
```

### Snake Case

If you do not specify a value for `naming_convention` or set it to `snake_case`, the plugin will use snake case for table and column names, and table names will have a `salesforce*` prefix.

For example:

```sql
select
  id,
  who_count,
  what_count,
  subject,
  is_all_day_event
from
  salesforce_event;
```

```
+---------------------+-----------+------------+---------+------------------+
| id                  | who_count | what_count | subject | is_all_day_event |
+----------------------------------------------+----------------------------+
| 00U2t0000000Mw3dEAD | 0         |  0         | test    | false            |
+---------------------+-----------+------------+---------+------------------+
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-salesforce
- Community: [Slack Channel](https://steampipe.io/community/join)
