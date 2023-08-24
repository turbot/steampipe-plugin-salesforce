# Table: salesforce_opportunity

Represents an opportunity, which is a sale or pending deal.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_opportunity#api_native_examples).

## Examples

### Basic info

```sql
select
  id,
  name,
  amount,
  type,
  stage_name,
  forecast_category,
  is_won
from
  salesforce_opportunity
order by
  stage_name;
```

### List only won opportunities

```sql
select
  id,
  name,
  amount,
  type,
  is_won
from
  salesforce_opportunity
where
  is_won
order by
  amount;
```

### List number of opportunities with respective stage

```sql
select
  count(*),
  stage_name
from
  salesforce_opportunity
group by
  stage_name;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
order by
  "StageName";
```

### List only won opportunities (with API Native naming convention)

```sql
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory",
  "IsWon"
from
  "Opportunity"
where
  "IsWon"
order by
  "Amount";
```

### List closed opportunities

```sql
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
where
  "CloseDate" <= now();
```

### List opportunities with open activity

```sql
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
where
  "HasOpenActivity";
```
