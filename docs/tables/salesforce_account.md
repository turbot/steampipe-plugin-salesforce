# Table: salesforce_account

Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).

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
