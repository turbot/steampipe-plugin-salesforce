---
title: "Steampipe Table: salesforce_opportunity_contact_role - Query Salesforce Opportunity Contact Roles using SQL"
description: "Allows users to query Opportunity Contact Roles in Salesforce, providing insights into the role that a contact plays in a specific opportunity."
---

# Table: salesforce_opportunity_contact_role - Query Salesforce Opportunity Contact Roles using SQL

Opportunity Contact Role in Salesforce is a junction object that links Contacts to Opportunities, showing the role that the contact plays in the specific opportunity. This role could be a decision maker, influencer, etc. It provides a way to specify the individuals who are involved in the business deal.

## Table Usage Guide

The `salesforce_opportunity_contact_role` table provides insights into the roles that contacts play in specific opportunities within Salesforce. As a sales manager or business analyst, you can explore details about these roles through this table, including the contact's involvement in the opportunity, their influence level, and other associated metadata. Utilize it to uncover information about the relationships between contacts and opportunities, such as who are the main decision makers or influencers in a particular deal.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Explore which roles are associated with different opportunities in Salesforce to better understand the distribution of responsibilities within your organization. This can help to identify any gaps or overlaps in roles assigned to specific opportunities.

```sql+postgres
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

```sql+sqlite
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
Determine the primary roles within your opportunities to better understand your salesforce interactions. This can help identify key contacts and their respective roles, helping to streamline communication and improve sales strategies.

```sql+postgres
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

```sql+sqlite
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

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Explore the primary contacts associated with specific business opportunities. This can help you understand the key individuals involved in each opportunity and their respective roles.

```sql+postgres
select
  "ID",
  "ContactID",
  "IsPrimary",
  "OpportunityID",
  "StageName"
from
  "OpportunityContactRole";
```

```sql+sqlite
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
Identify the primary roles within opportunities to understand their significance and influence within your organization's operations.

```sql+postgres
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

```sql+sqlite
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
Discover the roles of contacts associated with opportunities that were created in the past month. This could be useful for assessing recent changes in contact roles and understanding their involvement in new opportunities.

```sql+postgres
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

```sql+sqlite
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
  "CreatedDate" <= datetime('now', '-30 day');
```

### Show decision maker opportunity contact roles
Explore which roles in an opportunity are designated as 'Decision Makers'. This can be useful in identifying key individuals within a business opportunity.

```sql+postgres
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

```sql+sqlite
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