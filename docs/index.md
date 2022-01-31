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

#### Setup your connected application for `Client Id & Client Secret`:

- [Create your connected application](https://trailhead.salesforce.com/en/content/learn/projects/build-a-connected-app-for-api-integration/create-a-connected-app)
- When you create a connected app, make sure that you understand how it’s going to be used so you can configure the appropriate settings.
  - **For example**, if you’re creating a connected app to integrate an external application with your Salesforce API, configure the connected app with OAuth authorization settings.

|                                            | `REQUIRED EDITIONS`                                                                           |
| ------------------------------------------ | --------------------------------------------------------------------------------------------- |
| Available in                               | both `Salesforce Classic` (not available in all orgs) and `Lightning Experience`              |
| Connected applications can be created in   | `Group, Professional, Enterprise, Essentials, Performance, Unlimited, and Developer Editions` |
| Connected applications can be installed in | `All Editions`                                                                                |

#### Configure basic connected application settings

1. From Setup, enter Apps in the Quick Find box, and select `App Manager.`
2. Click New `Connected App.`
3. Enter the connected app’s name, which displays in the App Manager and on its App Launcher tile.
4. The connected app name must be unique within your org. If the connected app was created in the Spring ‘14 release or later, you can reuse the name of a deleted connected app.
5. If you have a web page with more information about your app, provide an info URL.

#### Reset Your Security Token

- [Reset your security token](https://help.salesforce.com/articleView?id=user_security_token.htm&type=5)
- From your personal settings, enter `Reset` in the `Quick Find` box, then select `Reset My Security Token`
- Click Reset Security Token. The new security token is sent to the email address in your Salesforce personal settings.

  - A new security token is emailed to you when you reset your password. Or you can reset your token separately.
  - After generating the `ClientId and Security Token` with the above steps. Update the value in `~/steampipe/config/salesforce.spc`

### Configuration

Installing the latest salesforce plugin will create a config file (`~/.steampipe/config/salesforce.spc`) with a single connection named `salesforce`:

```hcl
connection "salesforce" {
  plugin = "salesforce"

  # Your salesforce instance URL, for example, 'https://na01.salesforce.com/'
  # url       = "YOUR_SALESFORCE_URL"

  # Username of the Salesforce account
  # user      = "YOUR_SALESFORCE_USERNAME"

  # Password of the Salesforce account
  # password  = "YOUR_SALESFORCE_PASSWORD"

  # Security token, could be omitted if Trusted IP is configured
  # token     = "YOUR_SALESFORCE_TOKEN"

  # Client Id of Salesforce from Connected App"
  # client_id = "YOUR_SALESFORCE_CLIENT_ID"

  # List of salesforce tables to be generated. Salesforce Object API Name should be used here""
  # tables = ["Case", "Campaign"]

  # Version of the salesforce API
  # api_version = "43.0"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-salesforce
- Community: [Slack Channel](https://steampipe.io/community/join)
