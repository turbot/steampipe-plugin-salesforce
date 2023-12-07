---
title: "Steampipe Table: salesforce_lead - Query Salesforce Leads using SQL"
description: "Allows users to query Leads in Salesforce, providing detailed information about potential sales opportunities and customer data."
---

# Table: salesforce_lead - Query Salesforce Leads using SQL

Salesforce Leads are a key resource in Salesforce, a customer relationship management solution. They represent potential sales opportunities and contain information about individuals or representatives of organizations who are interested in the product or service that the company offers. Leads can be converted into accounts, contacts, or opportunities when they are qualified.

## Table Usage Guide

The `salesforce_lead` table provides insights into Leads within Salesforce. As a sales manager or marketing professional, explore Lead-specific details through this table, including personal information, source, status, and associated company details. Utilize it to track and analyze potential sales opportunities, understand customer behavior, and plan targeted marketing strategies.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Explore the various leads in your Salesforce system to understand their status and gain insights into their associated industries. This can assist in identifying potential conversion opportunities and assessing lead quality.

```sql+postgres
select
  id,
  name,
  company,
  email,
  industry,
  is_converted,
  rating,
  status,
  website
from
  salesforce_lead;
```

```sql+sqlite
select
  id,
  name,
  company,
  email,
  industry,
  is_converted,
  rating,
  status,
  website
from
  salesforce_lead;
```

### List number of leads by industry type
Analyze the distribution of leads across different industry sectors. This is beneficial for understanding which industries are more engaged with your business, aiding in strategic decision-making.

```sql+postgres
select
  count(*),
  industry
from
  salesforce_lead
group by
  industry;
```

```sql+sqlite
select
  count(*),
  industry
from
  salesforce_lead
group by
  industry;
```

### List number of leads by status
Determine the distribution of leads based on their status. This can provide insights into your sales pipeline and help pinpoint areas for improvement.

```sql+postgres
select
  count(*),
  status
from
  salesforce_lead
group by
  status;
```

```sql+sqlite
select
  count(*),
  status
from
  salesforce_lead
group by
  status;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Discover the segments that are converted and their respective ratings in a particular industry. This can be used to pinpoint specific areas for potential business growth and improvement.

```sql+postgres
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead";
```

```sql+sqlite
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead";
```

### List number of leads by industry type (with API Native naming convention)
Determine the distribution of leads across various industry types. This can help in identifying which industries are more engaged with your services, enabling targeted marketing efforts.

```sql+postgres
select
  count(*),
  "Industry"
from
  "Lead"
group by
  "Industry";
```

```sql+sqlite
select
  count(*),
  "Industry"
from
  "Lead"
group by
  "Industry";
```

### List cold rated leads
Discover the segments that are identified as 'cold' leads in your business. This is useful to target specific marketing strategies to improve conversion rates.

```sql+postgres
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead"
where
  "Rating" = 'Cold';
```

```sql+sqlite
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead"
where
  "Rating" = 'Cold';
```

### List qualified leads
Explore which leads are qualified, allowing you to focus your efforts on potential customers who have been pre-identified as likely to convert. This can help streamline your sales process and improve efficiency.

```sql+postgres
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead"
where
  "Status" = 'Qualified';
```

```sql+sqlite
select
  "ID",
  "Name",
  "Industry",
  "IsConverted",
  "Rating",
  "Status",
  "Website"
from
  "Lead"
where
  "Status" = 'Qualified';
```