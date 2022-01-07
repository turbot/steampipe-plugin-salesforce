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

# PagerDuty + Steampipe

[Salesforce](https://www.salesforce.com/) is a customer relationship management (CRM) platform.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List won oppurtinities:

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

Download and install the latest Salesforce plugin:

```bash
steampipe plugin install salesforce
```

### Credentials

| Item        | Description                                                                                                                             |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | [Get your user token](https://onlinehelp.coveo.com/en/ces/7.0/administrator/getting_the_security_token_for_your_salesforce_account.htm) |

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
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-salesforce
- Community: [Slack Channel](https://steampipe.io/community/join)
