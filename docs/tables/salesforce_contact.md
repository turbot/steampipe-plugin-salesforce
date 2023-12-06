---
title: "Steampipe Table: salesforce_contact - Query Salesforce Contacts using SQL"
description: "Allows users to query Contacts in Salesforce, providing details such as contact name, email, phone number, and associated account information."
---

# Table: salesforce_contact - Query Salesforce Contacts using SQL

Salesforce Contacts is a feature within Salesforce CRM that allows businesses to store and manage information about their customers. It provides a centralized way to keep track of customer details, including their contact information, associated accounts, and any activities or opportunities related to them. Salesforce Contacts helps businesses stay informed about their customer base and maintain strong relationships with them.

## Table Usage Guide

The `salesforce_contact` table provides insights into Contacts within Salesforce CRM. As a sales representative or a customer relationship manager, explore contact-specific details through this table, including name, email, phone number, and associated account information. Utilize it to uncover information about contacts, such as those who are associated with specific accounts, the activities related to them, and their communication details.

**Important Notes**
- If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account#api_native_examples).

## Examples

### Basic info
Explore the basic information of your Salesforce contacts to understand their source, department, and role within their organization. This can help in tailoring communication and marketing strategies to each contact.

```sql
select
  id,
  name,
  account_id,
  email,
  lead_source,
  title,
  department
from
  salesforce_contact;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)
Discover the segments that allow you to gain insights into your contact list, such as identifying the associated account and department for each contact, without delving into overly technical details.

```sql
select
  "ID",
  "Name",
  "AccountID",
  "Email",
  "Department"
from
  "Contact";
```

### List deleted contacts
Uncover the details of contacts that have been deleted from your database. This can be useful for auditing purposes or for retrieving lost information.

```sql
select
  "ID",
  "Name",
  "AccountID",
  "Email",
  "Department"
from
  "Contact"
where
  "IsDeleted";
```

### Show contacts created in last 30 days
Explore which contacts have been added within the past month. This can help you keep track of recent additions and ensure that no new contacts have been missed.

```sql
select
  "ID",
  "Name",
  "AccountID",
  "Email",
  "Department"
from
  "Contact"
where
  "CreatedDate" <= now() - interval '30' day;
```