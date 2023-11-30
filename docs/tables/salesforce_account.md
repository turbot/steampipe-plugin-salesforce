---
title: "Steampipe Table: salesforce_account - Query Salesforce Accounts using SQL"
description: "Allows users to query Salesforce Accounts, providing insights into customer account details and related information."
---

# Table: salesforce_account - Query Salesforce Accounts using SQL

Salesforce Accounts are companies or individuals that are involved in a business relationship. These can include customers, competitors, and partners. Accounts are used to store information such as name, address, and phone numbers, for easy reference.

## Table Usage Guide

The `salesforce_account` table provides insights into Salesforce Accounts. As a Sales or CRM Manager, explore account-specific details through this table, including name, address, and phone numbers. Utilize it to uncover information about accounts, such as those with recent activities, the relationships between accounts, and the verification of account details.

## Examples

### Basic info
Explore the general information about your Salesforce accounts, such as their names, descriptions, annual revenues, ownership types, industries, creation dates, ratings, and websites. This can help you gain insights into your accounts' performance, their industries, and how they have been rated over time.

```sql
select
  id,
  name,
  description,
  annual_revenue,
  ownership,
  industry,
  created_date,
  rating,
  website
from
  salesforce_account
```

### List number of accounts by industry type
Discover the segments that each industry occupies in your Salesforce accounts. This can help you understand the distribution of your accounts across different industries, providing valuable insights for business strategy planning and decision making.

```sql
select
  count(*),
  industry
from
  salesforce_account
group by
  industry;
```

### List number of accounts by ownership
Explore the distribution of accounts based on ownership to better understand the structure of your Salesforce accounts. This can help in assessing the balance of account ownership and strategizing future sales and marketing efforts.

```sql
select
  count(*),
  ownership
from
  salesforce_account
group by
  ownership;
```

### List accounts with hot rating
This query can be used to identify high-priority accounts in your Salesforce database by finding those labeled with a 'Hot' rating. This can help sales teams focus their efforts on the most promising leads, improving efficiency and potentially increasing sales.

```sql
select
  id,
  name,
  description,
  annual_revenue,
  ownership,
  industry,
  created_date,
  rating,
  website
from
  salesforce_account
where
  rating = 'Hot'
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Gain insights into various aspects of your business accounts such as revenue, ownership, industry type, and rating. This query is useful for a comprehensive overview, helping in strategic decision-making and performance evaluation.

```sql
select
  "ID",
  "Name",
  "Description",
  "AnnualRevenue",
  "Ownership",
  "Industry",
  "Rating"
from
  "Account";
```

### List number of accounts by industry type (with API Native naming convention)
Determine the distribution of accounts across different industry types. This can be useful for understanding the diversity of your client base and tailoring your services accordingly.

```sql
select
  count(*),
  "Industry"
from
  "Account"
group by
  "Industry";
```

### Show details about the turbot account
Gain insights into the Turbot account's details such as annual revenue, ownership, industry, and rating to better understand its financial and industrial standing.

```sql
select
  "ID",
  "Name",
  "Description",
  "AnnualRevenue",
  "Ownership",
  "Industry",
  "Rating"
from
  "Account"
where
  "Name" = 'turbot';
```

### Show customer accounts
Explore which customer accounts exist in your database to gain insights into their annual revenue, industry, and rating. This can help in understanding the financial health and industry distribution of your customer base.

```sql
select
  "ID",
  "Name",
  "Description",
  "AnnualRevenue",
  "Ownership",
  "Industry",
  "Rating"
from
  "Account"
where
  "Type" = 'Customer';
```