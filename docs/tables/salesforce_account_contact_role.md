---
title: "Steampipe Table: salesforce_account_contact_role - Query Salesforce Account Contact Roles using SQL"
description: "Allows users to query Salesforce Account Contact Roles, providing specific details about the role of a contact within an account."
---

# Table: salesforce_account_contact_role - Query Salesforce Account Contact Roles using SQL

Salesforce Account Contact Role is a feature within Salesforce that allows users to specify the role that a contact plays within an account. It provides a way to define and manage the relationships between contacts and accounts, offering insights into the hierarchy and responsibilities within an organization. Salesforce Account Contact Role is crucial for managing customer relationships and understanding the dynamics within an account.

## Table Usage Guide

The `salesforce_account_contact_role` table offers insights into the roles of contacts within Salesforce accounts. As a Salesforce administrator or sales representative, you can use this table to understand the responsibilities and hierarchy of contacts within an account. This can help you manage customer relationships more effectively, identify key contacts, and understand the dynamics within an account.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Explore the roles assigned to different contacts within various accounts. This can assist in identifying key contacts and their roles in each account, which is crucial for effective account management and communication strategies.

```sql
select
  id,
  account_id,
  contact_id,
  is_primary,
  role
from
  salesforce_account_contact_role;
```

### List primary account contact role
Explore which contact roles are primarily associated with specific accounts. This is useful for understanding the main points of contact for each account.

```sql
select
  id,
  account_id,
  contact_id,
  is_primary
from
  salesforce_account_contact_role
where
  is_primary;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Gain insights into the primary contacts associated with each account, which can be useful in understanding the account's key relationships and interactions.

```sql
select
  "ID",
  "AccountID",
  "ContactID",
  "IsPrimary"
from
  "AccountContactRole";
```

### List primary account contact role (with API Native naming convention)
Explore which contacts are designated as the primary point of contact for their respective accounts. This is particularly useful for focusing communication efforts and understanding the main contact for each account.

```sql
select
  "ID",
  "AccountID",
  "ContactID",
  "IsPrimary"
from
  "AccountContactRole"
where
  "IsPrimary";
```

### Show approver account contact roles
Identify the primary contacts who hold the 'Approver' role within an organization. This can be useful to understand who has the authority to approve actions or changes within the system.

```sql
select
  "ID",
  "AccountID",
  "ContactID",
  "IsPrimary",
  "Role"
from
  "AccountContactRole"
where
  "Role" = 'Approver';
```

### Show account contact roles created in last 30 days
Identify recent changes in your account's contact roles to understand any alterations made within the past month.

```sql
select
  "ID",
  "AccountID",
  "ContactID",
  "IsPrimary",
  "Role"
from
  "AccountContactRole"
where
  "CreatedDate" <= now() - interval '30' day;
```