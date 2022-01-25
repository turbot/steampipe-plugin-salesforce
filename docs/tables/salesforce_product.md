# Table: salesforce_product

Represents a product that org sells.

## Examples

### Basic info

```sql
select
  id,
  name,
  family,
  is_active,
  created_by_id,
  quantity_unit_of_measure,
  stock_keeping_unit
from
  salesforce_product;
```

### List inactive products

```sql
select
  id,
  name,
  family,
  is_active,
  created_by_id,
  quantity_unit_of_measure,
  stock_keeping_unit
from
  salesforce_product
where
  not is_active;
```
