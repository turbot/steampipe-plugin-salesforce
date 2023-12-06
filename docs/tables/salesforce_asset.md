---
title: "Steampipe Table: salesforce_asset - Query Salesforce Assets using SQL"
description: "Allows users to query Assets in Salesforce, providing detailed information about assets associated with a particular account, such as purchase date, asset value, and asset status."
---

# Table: salesforce_asset - Query Salesforce Assets using SQL

Salesforce Assets are items that a company has sold, such as products or services. These assets are associated with a customer's account and can be used to track items like warranty or service information. Assets in Salesforce provide a comprehensive view of customers' purchased products, helping businesses understand their customers' needs and provide better service.

## Table Usage Guide

The `salesforce_asset` table enables users to gain insights into the assets associated with specific accounts in Salesforce. As a sales or customer service professional, you can utilize this table to track detailed information about sold products or services, including their purchase dates, values, and statuses. This information can be beneficial in understanding customer trends, managing warranties, and improving overall customer service.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
This example helps you understand the basic details of your assets in Salesforce, such as their ID, name, price, status, and quantity. It's useful for getting a quick overview of your assets and their associated contacts, which can aid in tracking asset performance and managing resources.

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
Determine the areas in which your assets are allocated by assessing their price, quantity, and status. This can help you manage your resources more effectively and identify areas for potential improvement.

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
Explore which assets have been shipped to manage inventory and ensure accurate tracking of asset movement. This is particularly useful for businesses to keep track of their resources and improve their asset management strategies.

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
Analyze your company's internal assets to gain insights into their status, quantity, and value. This can help in better resource allocation and financial planning.

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