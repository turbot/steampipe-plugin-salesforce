# Table: salesforce_asset

Represents an item of commercial value, such as a product sold by your company or a competitor, that a customer has purchased and installed.

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
