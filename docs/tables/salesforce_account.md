# Table: salesforce_account

Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

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

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)

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
  "Account";
```

### List number of accounts by industry type (with API Native naming convention)

```sql
select
  count(*),
  "Industry"
from
  "Account"
group by
  "Industry";
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
