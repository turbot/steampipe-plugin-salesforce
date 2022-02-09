---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/salesforce.svg"
brand_color: "#06ac38"
display_name: "Salesforce"
short_name: "salesforce"
description: "Steampipe plugin to query accounts, users, oppurtinities and more from your Salesforce instance."
og_description: "Query Salesforce with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/salesforce-social-graphic.png"
---

# Salesforce + Steampipe

[Salesforce](https://www.salesforce.com/) is a customer relationship management (CRM) platform.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List won opportunities:

```sql
select
  id,
  name,
  amount,
  close_date,
  fiscal_quarter,
  fiscal_year
from
  salesforce_opportunity
where
  is_won;
```

```
+--------------------+-------------------------------------+--------+---------------------------+----------------+-------------+
| id                 | name                                | amount | close_date                | fiscal_quarter | fiscal_year |
+--------------------+-------------------------------------+--------+---------------------------+----------------+-------------+
| 0065j00000Oi1SpAAJ | GenePoint Standby Generator         | 85000  | 2021-10-23T05:30:00+05:30 | 1              | 2015        |
| 0065j00000Oi1SzAAJ | GenePoint SLA                       | 30000  | 2021-12-16T05:30:00+05:30 | 2              | 2015        |
| 0065j00000Oi1SoAAJ | Express Logistics Standby Generator | 220000 | 2021-09-15T05:30:00+05:30 | 1              | 2015        |
+--------------------+-------------------------------------+--------+---------------------------+----------------+-------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/salesforce/tables)**

## Get started

### Install

Download and install the latest salesforce plugin:

```bash
steampipe plugin install salesforce
```

### Credentials

- [Create your connected application](https://trailhead.salesforce.com/en/content/learn/projects/build-a-connected-app-for-api-integration/create-a-connected-app)

- Configure basic [connected application settings](https://help.salesforce.com/s/articleView?id=sf.connected_app_create_basics.htm&type=5)

- Reset your [security token](https://help.salesforce.com/articleView?id=user_security_token.htm&type=5)

### Configuration

Installing the latest salesforce plugin will create a config file (`~/.steampipe/config/salesforce.spc`) with a single connection named `salesforce`:

```hcl
connection "salesforce" {
  plugin = "salesforce"

  # Your salesforce instance URL, for example, 'https://na01.salesforce.com/'
  # url       = "https://na01.salesforce.com/"

  # Username of the Salesforce account
  # username  = "test@example.com"

  # Password of the Salesforce account
  # password  = "Dummy@~Password"

  # The Salesforce security token is only required If the client's IP address is not added to the organization's list of trusted IPs
  # https://help.salesforce.com/s/articleView?id=sf.security_networkaccess.htm&type=5
  # https://migration.trujay.com/help/how-to-add-an-ip-address-to-whitelist-on-salesforce/
  # token     = "ABO5C3PNqOP0BHsPGwldmssv1"

  # Client Id of Connected App from Salesforce
  # client_id = "YOUR_SALESFORCE_CLIENT_ID"

  # List of salesforce tables to be generated. Salesforce Object API Name (refer https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm) should be used here."
  # Currently plugin supports a standard set of Salesforce tables by default in the plugin. If there is a need to query an standard object from, you can add standard object in the objects list. Steampipe will automatically generate respective table for the connection.
  # objects = ["Case", "Campaign", "CustomApp__c"]

  # Version of the salesforce API
  # api_version = "43.0"
}
```

## Custom Fields

Salesforce support addition of the [custom fields](https://help.salesforce.com/s/articleView?id=sf.adding_fields.htm&type=5) to standard objects.

**Note**: If you have set up Salesforce credentials correctly in Steampipe configuration. Steampipe will generate the tables schema with all the custom fields along with standard object fields dynamically.

<!-- https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm -->

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-salesforce
- Community: [Slack Channel](https://steampipe.io/community/join)
