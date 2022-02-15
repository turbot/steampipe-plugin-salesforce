# Table: salesforce_order

Represents an order associated with a contract or an account.

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
