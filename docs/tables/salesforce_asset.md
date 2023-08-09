# Table: salesforce_asset

Represents an item of commercial value, such as a product sold by your company or a competitor, that a customer has purchased and installed.

If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select id, name from salesforce_asset` would become `select "ID", "Name" from "Asset"`.

Additional examples can be found [below](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_asset#list_shipped_assets).

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
