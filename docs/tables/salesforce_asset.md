# Table: salesforce_asset

Represents an item of commercial value, such as a product sold by your company or a competitor, that a customer has purchased and installed.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_asset#api_native_examples).

## Examples

### Basic info

```sql
select
  id,
  name,
  price,
  status,
  asset_level,
  contact_id,
  quantity
from
  salesforce_asset;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "Name",
  "Price",
  "Status",
  "Quantity"
from
  "Asset";
```

### List shipped assets

```sql
select
  "ID",
  "Name",
  "Price",
  "Status",
  "Quantity"
from
  "Asset"
where
  "Status" = 'Shipped';
```

### List internal assets

```sql
select
  "ID",
  "Name",
  "Price",
  "Status",
  "Quantity"
from
  "Asset"
where
  "IsInternal";
```
