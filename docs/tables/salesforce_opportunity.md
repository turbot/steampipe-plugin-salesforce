---
title: "Steampipe Table: salesforce_opportunity - Query Salesforce Opportunities using SQL"
description: "Allows users to query Opportunities in Salesforce, specifically to retrieve details about business deals and their current sales stage, providing insights into sales performance and potential revenue."
---

# Table: salesforce_opportunity - Query Salesforce Opportunities using SQL

Salesforce Opportunity is a component of Salesforce Sales Cloud that allows businesses to manage their sales cycles and deals. It provides a way to track and measure potential revenue, and monitor sales stages, from prospecting to closing. Salesforce Opportunity is instrumental in understanding a company's sales pipeline and forecasting sales.

## Table Usage Guide

The `salesforce_opportunity` table provides insights into Opportunities within Salesforce Sales Cloud. As a sales manager or data analyst, explore opportunity-specific details through this table, including deal value, expected close date, and associated account details. Utilize it to uncover information about sales performance, such as deals in the pipeline, the value of potential revenue, and the progress of deals through sales stages.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Gain insights into your sales opportunities by sorting them based on their stage in the sales process. This helps prioritize opportunities that are closer to closing, improving sales efficiency and forecasting.

```sql+postgres
select
  id,
  name,
  amount,
  type,
  stage_name,
  forecast_category,
  is_won
from
  salesforce_opportunity
order by
  stage_name;
```

```sql+sqlite
select
  id,
  name,
  amount,
  type,
  stage_name,
  forecast_category,
  is_won
from
  salesforce_opportunity
order by
  stage_name;
```

### List only won opportunities
Explore which sales opportunities have been successful to understand the areas of high revenue generation. This can help in strategizing future sales approaches.

```sql+postgres
select
  id,
  name,
  amount,
  type,
  is_won
from
  salesforce_opportunity
where
  is_won
order by
  amount;
```

```sql+sqlite
select
  id,
  name,
  amount,
  type,
  is_won
from
  salesforce_opportunity
where
  is_won
order by
  amount;
```

### List number of opportunities with respective stage
Determine the distribution of sales opportunities across different stages. This can help you understand where most of your opportunities are concentrated, allowing you to strategize and allocate resources effectively.

```sql+postgres
select
  count(*),
  stage_name
from
  salesforce_opportunity
group by
  stage_name;
```

```sql+sqlite
select
  count(*),
  stage_name
from
  salesforce_opportunity
group by
  stage_name;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore the opportunities in your pipeline, focusing on their names, amounts, types, and forecast categories, to gain insights into their stages. This is useful for understanding the distribution and potential value of your opportunities, which can inform your sales strategy and prioritization.

```sql+postgres
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
order by
  "StageName";
```

```sql+sqlite
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
order by
  "StageName";
```

### List only won opportunities (with API Native naming convention)
Discover the segments that have been successful by identifying instances where opportunities have been won. This can help assess the elements within these successful opportunities to strategize for future prospects.

```sql+postgres
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory",
  "IsWon"
from
  "Opportunity"
where
  "IsWon"
order by
  "Amount";
```

```sql+sqlite
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory",
  "IsWon"
from
  "Opportunity"
where
  "IsWon"
order by
  "Amount";
```

### List closed opportunities
Explore which business opportunities have already reached their close date. This can help in assessing past performance and strategizing for future opportunities.

```sql+postgres
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
where
  "CloseDate" <= now();
```

```sql+sqlite
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
where
  "CloseDate" <= datetime('now');
```

### List opportunities with open activity
Explore which opportunities have an open activity to determine areas for potential business growth. This can help in identifying instances where immediate action is required, aiding in effective decision making.

```sql+postgres
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
where
  "HasOpenActivity";
```

```sql+sqlite
select
  "ID",
  "Name",
  "Amount",
  "Type",
  "StageName",
  "ForecastCategory"
from
  "Opportunity"
where
  "HasOpenActivity";
```