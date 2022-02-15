# Table: salesforce_opportunity

Represents an opportunity, which is a sale or pending deal.

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
