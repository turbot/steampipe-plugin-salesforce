# Table: salesforce_pricebook

Represents a price book that contains the list of products that your org sells.

## Examples

### Basic info

```sql
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook;
```

### List standard price books

```sql
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook
where
  is_standard;
```

### List active price books

```sql
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook
where
  is_active;
```
