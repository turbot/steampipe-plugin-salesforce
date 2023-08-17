# Table: salesforce_lead

Represents a prospect or lead.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_lead#api_native_examples).

## Examples

### Basic info

```sql
select
  id,
  name,
  company,
  email,
  industry,
  is_converted,
  rating,
  status,
  website
from
  salesforce_lead;
```

### List number of leads by industry type

```sql
select
  count(*),
  industry
from
  salesforce_lead
group by
  industry;
```

### List number of leads by status

```sql
select
  count(*),
  status
from
  salesforce_lead
group by
  status;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead";
```

### List number of leads by industry type (with API Native naming convention)

```sql
select
  count(*),
  "Industry"
from
  "Lead"
group by
  "Industry";
```

### List cold rated leads

```sql
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead"
where
  "Rating" = 'Cold';
```

### List qualified leads

```sql
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead"
where
  "Status" = 'Qualified';
```
