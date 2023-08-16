# Table: salesforce_contract

Represents a contract (a business agreement) associated with an Account.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_contract#api_native_examples).

## Examples

### Basic info

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

### Basic info (with API Native naming convention)

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
