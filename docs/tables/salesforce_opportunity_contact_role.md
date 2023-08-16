# Table: salesforce_opportunity_contact_role

Represents the role that a Contact plays on an Opportunity.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_opportunity_contact_role#show_opportunity_contact_roles_created_in_last_30_days).

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

### List primary opportunity contact roles

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

## API Native Examples

### Basic info (with API Native naming convention)

```sql
select
  "ID",
  "ContactID",
  "IsPrimary",
  "OpportunityID",
  "StageName"
from
  "OpportunityContactRole";
```

### List primary opportunity contact roles (with API Native naming convention)

```sql
select
  "ID",
  "ContactID",
  "IsPrimary",
  "OpportunityID",
  "StageName"
from
  "OpportunityContactRole"
where
  "IsPrimary";
```

### Show opportunity contact roles created in last 30 days

```sql
select
  "ID",
  "ContactID",
  "IsPrimary",
  "OpportunityID",
  "StageName",
  "Role"
from
  "OpportunityContactRole"
where
  "CreatedDate" <= now() - interval '30' day;
```

### Show decision maker opportunity contact roles

```sql
select
  "ID",
  "ContactID",
  "IsPrimary",
  "OpportunityID",
  "StageName",
  "Role"
from
  "OpportunityContactRole"
where
  "Role" = 'Decision Maker';
```
