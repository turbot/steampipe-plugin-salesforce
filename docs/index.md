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

### Configuration

Installing the latest salesforce plugin will create a config file (`~/.steampipe/config/salesforce.spc`) with a single connection named `salesforce`:

```hcl
connection "salesforce" {
  plugin = "salesforce"

  # Salesforce instance URL, e.g., "https://na01.salesforce.com/"
  # url = "YOUR_SALESFORCE_URL"

  # Salesforce account name
  # user = "YOUR_SALESFORCE_USERNAME"

  # Salesforce account password
  # password = "YOUR_SALESFORCE_PASSWORD"

  # The Salesforce security token is only required If the client's IP address is not added to the organization's list of trusted IPs
  # https://help.salesforce.com/s/articleView?id=sf.security_networkaccess.htm&type=5
  # https://migration.trujay.com/help/how-to-add-an-ip-address-to-whitelist-on-salesforce/
  # token     = "YOUR_SALESFORCE_TOKEN"

  # Client Id of Connected App from Salesforce
  # client_id = "YOUR_SALESFORCE_CLIENT_ID"

  # List of Salesforce object names to generate additional tables for
  # This argument only accepts exact Salesforce standard and custom object names, e.g., AccountBrand, OpportunityStage, CustomApp__c
  # For a full list of standard object names, please see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm)
  # All custom object names should end in "__c", following Salesforce object naming
  # objects = ["AccountBrand", "OpportunityStage", "CustomApp__c"]

  # Salesforce API version to connect to.
  # api_version = "43.0"
}
```

### Credentials

- [Create your connected application](https://trailhead.salesforce.com/en/content/learn/projects/build-a-connected-app-for-api-integration/create-a-connected-app)

- Configure basic [connected application settings](https://help.salesforce.com/s/articleView?id=sf.connected_app_create_basics.htm&type=5)

- Reset your [security token](https://help.salesforce.com/articleView?id=user_security_token.htm&type=5)

### Credentials

#### Setup your connected application for `Client ID` & `Client Secret`:

- [Create your connected application](https://trailhead.salesforce.com/en/content/learn/projects/build-a-connected-app-for-api-integration/create-a-connected-app)
- When you create a connected app, make sure that you understand how it’s going to be used so you can configure the appropriate settings.
  - **For example**, if you’re creating a connected app to integrate an external application with your Salesforce API, configure the connected app with OAuth authorization settings.
- Note: Connected Apps can only be created in: Group, Professional, Enterprise, Essentials, Performance, Unlimited, and Developer Editions, but installed in all editions.

#### Configure basic connected application settings

1. From Setup, enter Apps in the Quick Find box, and select `App Manager.`
2. Click New `Connected App.`
3. Enter the connected app’s name, which displays in the App Manager and on its App Launcher tile.
4. The connected app name must be unique within your org. If the connected app was created in the Spring ‘14 release or later, you can reuse the name of a deleted connected app.
5. If you have a web page with more information about your app, provide an info URL.

#### Reset your security token

- [Reset your security token](https://help.salesforce.com/articleView?id=user_security_token.htm&type=5)
- After generating the `Client ID` and `Security Token` with the above steps. Update the value in `~/steampipe/config/salesforce.spc`

## Custom Fields

Salesforce support addition of the [custom fields](https://help.salesforce.com/s/articleView?id=sf.adding_fields.htm&type=5) to standard objects.

If you have set up Salesforce credentials correctly in the Steampipe configuration, Steampipe will generate the tables schema with all the custom fields along with standard object fields dynamically.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-salesforce
- Community: [Slack Channel](https://steampipe.io/community/join)
