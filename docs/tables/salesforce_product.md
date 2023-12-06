---
title: "Steampipe Table: salesforce_product - Query Salesforce Products using SQL"
description: "Allows users to query Products in Salesforce, providing detailed information about each product, including its standard and custom fields."
---

# Table: salesforce_product - Query Salesforce Products using SQL

Salesforce Products are used to represent the items you sell. They can be goods, services, or digital content that are part of your business model. Products are typically associated with Opportunities or Quotes in Salesforce.

## Table Usage Guide

The `salesforce_product` table provides insights into Products within Salesforce. As a Sales or Business Analyst, explore product-specific details through this table, including standard and custom fields. Utilize it to uncover information about products, such as their associated Opportunities or Quotes, and the verification of custom fields.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).
- If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select id, name from salesforce_product` would become `select "ID", "Name" from "Product2"`.

## Examples

### Basic info
Explore which products in your Salesforce inventory are active and who created them. This information can be useful for auditing purposes and to understand the distribution of product responsibilities within your team.

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
Discover the segments that consist of inactive products in your Salesforce database. This can help you identify which products are not currently in use, allowing you to make informed decisions about inventory management and product offerings.

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

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore which products are active and in stock by assessing key attributes such as product ID, name, and family. This can help in managing inventory and understanding product performance.

```sql
select
  "ID",
  "Name",
  "Family",
  "IsActive",
  "StockKeepingUnit"
from
  "Product2";
```

### List inactive products (with API Native naming convention)
Discover the segments that contain inactive products in the inventory to better manage stock and potentially discontinue or replenish certain items.

```sql
select
  "ID",
  "Name",
  "Family",
  "IsActive",
  "StockKeepingUnit"
from
  "Product2"
where
  not "IsActive";
```

### Show archived products
Explore which products have been archived in your inventory, allowing you to keep track of items no longer actively sold or in circulation.

```sql
select
  "ID",
  "Name",
  "Family",
  "IsActive",
  "StockKeepingUnit"
from
  "Product2"
where
  "IsArchived";
```

### Show deleted products
Uncover the details of products that have been removed from your inventory. This is useful for tracking product lifecycle and managing stock levels effectively.

```sql
select
  "ID",
  "Name",
  "Family",
  "IsActive",
  "StockKeepingUnit"
from
  "Product2"
where
  "IsDeleted";
```