# Table: salesforce_lead

Represents a prospect or lead.

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
