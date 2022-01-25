# Table: salesforce_opportunity_contact_role

Represents the role that a Contact plays on an Opportunity.

## Examples

### Basic info

```sql
select
  id,
  contact_id,
  is_primary,
  opportunity_id,
  stage_name,
  role
from
  salesforce_opportunity_contact_role;
```

### List primary opportunity contact role

```sql
select
  id,
  opportunity_id,
  contact_id,
  is_primary
from
  salesforce_opportunity_contact_role
where
  is_primary;
```
