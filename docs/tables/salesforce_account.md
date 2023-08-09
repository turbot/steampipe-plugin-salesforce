# Table: salesforce_account

Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).

If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select name, annual_revenue from salesforce_account` would become `select "Name", "AnnualRevenue" from "Account"`.

Additional examples can be found [below](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#show_details_about_the_turbot_account).

## Examples

### Basic info

```sql
select
  id,
  name,
  description,
  annual_revenue,
  ownership,
  industry,
  created_date,
  rating,
  website
from
  salesforce_account
```

### List number of accounts by industry type

```sql
select
  count(*),
  industry
from
  salesforce_account
group by
  industry;
```

### List number of accounts by ownership

```sql
select
  count(*),
  ownership
from
  salesforce_account
group by
  ownership;
```

### List accounts with hot rating

```sql
select
  id,
  name,
  description,
  annual_revenue,
  ownership,
  industry,
  created_date,
  rating,
  website
from
  salesforce_account
where
  rating = 'Hot'
```

### Show details about the turbot account

```sql
select
  "ID",
  "Name",
  "Description",
  "AnnualRevenue",
  "Ownership",
  "Industry",
  "Rating"
from
  "Account"
where
  "Name" = 'turbot';
```

### Show customer accounts

```sql
select
  "ID",
  "Name",
  "Description",
  "AnnualRevenue",
  "Ownership",
  "Industry",
  "Rating"
from
  "Account"
where
  "Type" = 'Customer';
```
