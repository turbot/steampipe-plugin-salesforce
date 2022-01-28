# Table: salesforce_permission_set

Represents a set of permissions that's used to grant more access to one or more users without changing their profile or reassigning profiles.
PermissionSet has a read-only child relationship with PermissionSetGroup. PermissionSet contains the aggregated permissions for the group.
You can use permission sets to grant access, but not to deny access.

## Examples

### Basic info

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
