---
title: "Steampipe Table: salesforce_order - Query Salesforce Orders using SQL"
description: "Allows users to query Salesforce Orders, specifically the details of all orders, providing insights into order data and related patterns."
---

# Table: salesforce_order - Query Salesforce Orders using SQL

Salesforce Orders is a feature within Salesforce that allows businesses to manage and track customer orders. It provides a comprehensive view of all orders, including their details, status, and associated accounts. Salesforce Orders helps businesses streamline their order management process, ensuring efficient and effective customer service.

## Table Usage Guide

The `salesforce_order` table provides insights into Order data within Salesforce. As a Sales or Customer Service representative, explore order-specific details through this table, including order number, status, account details, and associated metadata. Utilize it to uncover information about orders, such as those pending, completed, or associated with specific accounts, aiding in efficient order management and customer service.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Explore the sales orders in Salesforce to analyze their status and total amount. This can be useful for assessing the overall sales performance and identifying any unusual order patterns.

```sql+postgres
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

```sql+sqlite
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
Determine the distribution of order statuses to gain insights into business performance. This can help identify areas that require attention, such as unfulfilled orders or returns.

```sql+postgres
select
  count(*),
  status
from
  salesforce_order
group by
  status;
```

```sql+sqlite
select
  count(*),
  status
from
  salesforce_order
group by
  status;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore which orders have a specific status or type to gain insights into overall sales performance and identify potential areas for improvement.

```sql+postgres
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

```sql+sqlite
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
Analyze the distribution of order statuses to better understand your business operations and customer behavior patterns.

```sql+postgres
select
  count(*),
  "Status"
from
  "Order"
group by
  "Status";
```

```sql+sqlite
select
  count(*),
  "Status"
from
  "Order"
group by
  "Status";
```

### List draft orders
Explore which orders are still in the draft stage to understand their status and manage them effectively. This is useful for tracking uncompleted transactions and identifying potential areas for follow-up or cancellation.

```sql+postgres
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

```sql+sqlite
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
Explore which orders have been deleted to maintain accurate records and ensure proper account management. This aids in tracking potential errors or fraudulent activities.

```sql+postgres
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

```sql+sqlite
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