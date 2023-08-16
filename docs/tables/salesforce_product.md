# Table: salesforce_product

Represents a product that org sells.

If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select id, name from salesforce_product` would become `select "ID", "Name" from "Product2"`.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_product#show_archived_products).

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

## API Native Examples

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "name",
  "Family",
  "IsActive",
  "StockKeepingUnit"
from
  "Product2";
```

### List inactive products (with API Native naming convention)

```sql
select
  "ID",
  "name",
  "Family",
  "IsActive",
  "StockKeepingUnit"
from
  "Product2"
where
  not "IsActive";
```

### Show archived products

```sql
select
  "ID",
  "name",
  "Family",
  "IsActive",
  "StockKeepingUnit"
from
  "Product2"
where
  "IsArchived";
```

### Show deleted products

```sql
select
  "ID",
  "name",
  "Family",
  "IsActive",
  "StockKeepingUnit"
from
  "Product2"
where
  "IsDeleted";
```
