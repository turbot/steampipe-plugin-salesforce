# Table: salesforce_contract

Represents a contract (a business agreement) associated with an Account.

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
