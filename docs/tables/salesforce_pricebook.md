---
title: "Steampipe Table: salesforce_pricebook - Query Salesforce Pricebooks using SQL"
description: "Allows users to query Pricebooks in Salesforce, specifically the list of all available pricebooks and their associated details."
---

# Table: salesforce_pricebook - Query Salesforce Pricebooks using SQL

A Salesforce Pricebook is a list of products and their prices available in Salesforce. It is a key component of Salesforce's product and price management, allowing businesses to manage different prices for the same product. Pricebooks can be specific to an organization, a business unit, or even a specific customer, providing flexibility in pricing strategies.

## Table Usage Guide

The `salesforce_pricebook` table provides insights into Pricebooks within Salesforce. As a sales manager or a business analyst, explore pricebook-specific details through this table, including related products, standard prices, and associated metadata. Utilize it to uncover information about pricing strategies, such as those specific to different business units or customers, and the verification of pricing policies.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).
- If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select name, is_active from salesforce_pricebook` would become `select "Name", "IsActive" from "Pricebook2"`.

## Examples

### Basic info
Analyze the settings to understand the status and creation details of price books in Salesforce. This is useful in assessing the active price books and their creators, which can help in sales strategy planning and management.

```sql+postgres
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook;
```

```sql+sqlite
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook;
```

### List standard price books
Explore the active standard price books in your Salesforce account. This can help you understand which price books have been created and are currently in use, offering insights into your pricing strategies and structures.

```sql+postgres
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook
where
  is_standard;
```

```sql+sqlite
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook
where
  is_standard;
```

### List active price books
Determine the active price books in your Salesforce data. This can help you understand which pricing structures are currently in use, aiding in revenue management and strategic planning.

```sql+postgres
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook
where
  is_active;
```

```sql+sqlite
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook
where
  is_active;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore the active and standard status of different price books in your database, along with their creation details and descriptions. This can help understand the variety and usage of different price books in your system.

```sql+postgres
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2";
```

```sql+sqlite
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2";
```

### List standard price books (with API Native naming convention)
Explore which standard price books are currently active and when they were created to understand their usage and relevance. Similarly, review the archived price books to determine which ones are no longer in use and when they were archived, providing insights into the company's pricing history and potential strategies.

```sql+postgres
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsStandard";
```

```sql+sqlite
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsStandard" = 1;
```

## Show archived price books

```sql+postgres
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsArchived";
```

```sql+sqlite
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsArchived";
```

### Show deleted price books
Discover the segments that have been marked as deleted in your pricing system. This information can be used to track changes, assess the impact of deletions, and potentially recover lost data.

```sql+postgres
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsDeleted";
```

```sql+sqlite
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsDeleted";
```