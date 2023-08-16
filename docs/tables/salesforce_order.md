# Table: salesforce_order

Represents an order associated with a contract or an account.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_order#list_draft_orders).

## Examples

### Basic info

```sql
select
  id,
  name,
  account_id,
  order_number,
  status,
  total_amount,
  type
from
  salesforce_order;
```

### List number of orders by status

```sql
select
  count(*),
  status
from
  salesforce_order
group by
  status;
```

## API Native Examples

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "Name",
  "AccountID",
  "OrderNumber",
  "Status",
  "TotalAmount",
  "Type"
from
  "Order";
```

### List number of orders by status (with API Native naming convention)

```sql
select
  count(*),
  "Status"
from
  "Order"
group by
  "Status";
```

### List draft orders

```sql
select
  "ID",
  "Name",
  "AccountID",
  "OrderNumber",
  "Status",
  "TotalAmount",
  "Type"
from
  "Order"
where
  "Status" = 'Draft';
```

### List deleted orders

```sql
select
  "ID",
  "Name",
  "AccountID",
  "OrderNumber",
  "Status",
  "TotalAmount",
  "Type"
from
  "Order"
where
  "IsDeleted";
```
