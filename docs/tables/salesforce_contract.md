---
title: "Steampipe Table: salesforce_contract - Query Salesforce Contracts using SQL"
description: "Allows users to query Salesforce Contracts, providing details about contracts, including their status, start and end dates, associated accounts, and other related information."
---

# Table: salesforce_contract - Query Salesforce Contracts using SQL

Salesforce Contracts are a record of the agreements between your company and the customers. They contain details such as the contract term (start and end date), contract status, associated account, and other related information. Contracts in Salesforce help to maintain the details of the service or support agreed upon during a sale.

## Table Usage Guide

The `salesforce_contract` table provides insights into contracts within Salesforce. As a Salesforce administrator, you can explore contract-specific details through this table, including the contract's status, its start and end dates, and the associated account. Use it to uncover information about contracts, such as those that are nearing their end dates, or to verify the details of a specific contract.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Gain insights into your company's contractual agreements, including the duration and status, to better manage your business relationships and plan for future engagements. This will assist in understanding the overall contract lifecycle within your organization.

```sql
select
  id,
  account_id,
  contract_number,
  contract_term,
  end_date,
  start_date,
  company_signed_date,
  status
from
  salesforce_contract;
```

### List contracts signed in the last six months
Discover the contracts that have been signed in the past six months to understand the recent business dealings and their status. This can be useful in assessing the volume and nature of recent contractual agreements.

```sql
select
  id,
  account_id,
  contract_number,
  contract_term,
  end_date,
  start_date,
  company_signed_date,
  status
from
  salesforce_contract
where
  company_signed_date >= (current_date - interval '6' month);
```

### List number of contracts by status
This query allows you to analyze the distribution of contracts within your Salesforce data according to their status. It aids in assessing the overall contract management efficiency and identifying areas that might need attention or improvement.

```sql
select
  count(*),
  status
from
  salesforce_contract
group by
  status;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Discover the segments that pertain to specific contract details such as term length and status. This can help in understanding the distribution and status of different contracts within an account.

```sql
select
  "ID",
  "AccountID",
  "ContractNumber",
  "ContractTerm",
  "EndDate",
  "StartDate",
  "Status"
from
  "Contract";
```

### List number of contracts by status (with API Native naming convention)
Determine the distribution of contracts based on their status to understand the operational dynamics of your business. This can help you identify areas that may require attention or improvement.

```sql
select
  count(*),
  "Status"
from
  "Contract"
group by
  "Status";
```

### Show draft contracts
Explore which contracts are still in draft status to understand potential upcoming agreements and their terms. This could help in planning resources or assessing future commitments.

```sql
select
  "ID",
  "AccountID",
  "ContractNumber",
  "ContractTerm",
  "EndDate",
  "StartDate",
  "Status"
from
  "Contract"
where
  "Status" = 'Draft';
```

### Show deleted contracts
Discover the contracts that have been deleted to understand any potential impact on business operations or account management. This can be useful in identifying gaps in service or potential revenue loss.

```sql
select
  "ID",
  "AccountID",
  "ContractNumber",
  "ContractTerm",
  "EndDate",
  "StartDate",
  "Status"
from
  "Contract"
where
  "IsDeleted";
```