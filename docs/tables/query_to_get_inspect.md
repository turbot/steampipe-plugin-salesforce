Query data from the table called {object_name}. A table is automatically created to represent each Object.

```bash
steampipe query --output=json "SELECT table_name, column_name, column_default, is_nullable, data_type, table_schema, ( Coalesce( (SELECT pg_catalog.Col_description(c.oid, cols.ordinal_position :: INT) FROM   pg_catalog.pg_class c WHERE  c.relname = cols.table_name AND c.relnamespace = (SELECT oid FROM   pg_catalog.pg_namespace WHERE  nspname = cols.table_schema)), '') ) AS column_comment, ( Coalesce( (SELECT pg_catalog.Obj_description(c.oid) FROM   pg_catalog.pg_class c WHERE  c.relname = cols.table_name AND c.relnamespace = (SELECT oid FROM   pg_catalog.pg_namespace WHERE  nspname = cols.table_schema)), '') ) AS table_comment FROM   information_schema.COLUMNS cols WHERE  cols.table_schema = 'salesforce' ORDER  BY cols.table_name, cols.column_name"
```

```json
[
  {
    "column_comment": "The Account Number.",
    "column_default": null,
    "column_name": "account_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The source of the account record. For example, Advertisement, Data.com, or Trade Show.",
    "column_default": null,
    "column_name": "account_source",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Active.",
    "column_default": null,
    "column_name": "active__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Estimated annual revenue of the account.",
    "column_default": null,
    "column_name": "annual_revenue",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The billing adress of the account.",
    "column_default": null,
    "column_name": "billing_address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates the record's clean status as compared with Data.com. Values are: Matched, Different,Acknowledged,NotFound,Inactive,Pending, SelectMatch, or Skipped.",
    "column_default": null,
    "column_name": "clean_status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The id of the user who created the account.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The creation date and time of the account.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Customer Priority.",
    "column_default": null,
    "column_name": "customer_priority__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The D&B Company ID.",
    "column_default": null,
    "column_name": "dandb_company_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Text description of the account.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The D-U-N-S Number.",
    "column_default": null,
    "column_name": "duns_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Account Fax.",
    "column_default": null,
    "column_name": "fax",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the account in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Primary business of account.",
    "column_default": null,
    "column_name": "industry",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether the object has been moved to the Recycle Bin (true) or not (false).",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Data.com Key.",
    "column_default": null,
    "column_name": "jigsaw",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Jigsaw Company ID.",
    "column_default": null,
    "column_name": "jigsaw_company_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Activity.",
    "column_default": null,
    "column_name": "last_activity_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The id of the user who last changed the contact fields, including modification date and time.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The date and time of last modification to account.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Master Record ID.",
    "column_default": null,
    "column_name": "master_record_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The NAICS Code.",
    "column_default": null,
    "column_name": "naics_code",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The NAICS Description.",
    "column_default": null,
    "column_name": "naics_desc",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Name of the account.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Number of employees working at the company represented by this account.",
    "column_default": null,
    "column_name": "number_of_employees",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Number of Locations.",
    "column_default": null,
    "column_name": "numberof_locations__c",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Operating Hour ID.",
    "column_default": null,
    "column_name": "operating_hours_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ID of the user who currently owns this account. Default value is the user logged in to the API to perform the create.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Ownership type for the account, for example Private, Public, or Subsidiary.",
    "column_default": null,
    "column_name": "ownership",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Parent Account ID.",
    "column_default": null,
    "column_name": "parent_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The contact's primary phone number.",
    "column_default": null,
    "column_name": "phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Photo URL.",
    "column_default": null,
    "column_name": "photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The account's prospect rating, for example Hot, Warm, or Cold.",
    "column_default": null,
    "column_name": "rating",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The shipping adress of the account.",
    "column_default": null,
    "column_name": "shipping_address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Standard Industrial Classification code of the company's main business categorization, for example, 57340 for Electronics.",
    "column_default": null,
    "column_name": "sic",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SIC Description.",
    "column_default": null,
    "column_name": "sic_desc",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Account Site.",
    "column_default": null,
    "column_name": "site",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SLA.",
    "column_default": null,
    "column_name": "sla__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SLA Expiration Date.",
    "column_default": null,
    "column_name": "sla_expiration_date__c",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SLA Serial Number.",
    "column_default": null,
    "column_name": "sla_serial_number__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Test Custom Field.",
    "column_default": null,
    "column_name": "test_custom_field__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The stock market symbol for this account.",
    "column_default": null,
    "column_name": "ticker_symbol",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "A name, different from its legal name, that an org may use for conducting business. Similar to “Doing business as” or \"DBA\".",
    "column_default": null,
    "column_name": "tradestyle",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Type of account, for example, Customer, Competitor, or Partner.",
    "column_default": null,
    "column_name": "type",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Upsell Opportunity.",
    "column_default": null,
    "column_name": "upsell_opportunity__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The website of this account, for example, www.acme.com.",
    "column_default": null,
    "column_name": "website",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Year Started.",
    "column_default": null,
    "column_name": "year_started",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).",
    "table_name": "salesforce_account",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the Account.",
    "column_default": null,
    "column_name": "account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the Contact associated with this account.",
    "column_default": null,
    "column_name": "contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who created contact role record.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date and time of the creation of the contact role record.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the account contact role in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Specifies whether the Contact plays the primary role on the Account (true) or not (false). Note that each account has only one primary contact role. Label is Primary. Default value is false.",
    "column_default": null,
    "column_name": "is_primary",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who most recently changed the contact role record.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date of most recent change in the contact role record.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Name of the role played by the Contact on this Account, such as Decision Maker, Approver, Buyer, and so on.",
    "column_default": null,
    "column_name": "role",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Account.",
    "table_name": "salesforce_account_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Account ID.",
    "column_default": null,
    "column_name": "account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Asset Level.",
    "column_default": null,
    "column_name": "asset_level",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Asset Provided By ID.",
    "column_default": null,
    "column_name": "asset_provided_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Asset Serviced By ID.",
    "column_default": null,
    "column_name": "asset_serviced_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact ID.",
    "column_default": null,
    "column_name": "contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created By ID.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Description.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Asset ID.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Install Date.",
    "column_default": null,
    "column_name": "install_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Competitor Asset.",
    "column_default": null,
    "column_name": "is_competitor_product",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Internal Asset.",
    "column_default": null,
    "column_name": "is_internal",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified By ID.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Asset Name.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Owner ID.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Parent Asset ID.",
    "column_default": null,
    "column_name": "parent_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Price.",
    "column_default": null,
    "column_name": "price",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Product ID.",
    "column_default": null,
    "column_name": "product_2_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Product Code.",
    "column_default": null,
    "column_name": "product_code",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Purchase Date.",
    "column_default": null,
    "column_name": "purchase_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Quantity.",
    "column_default": null,
    "column_name": "quantity",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Root Asset ID.",
    "column_default": null,
    "column_name": "root_asset_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Serial Number.",
    "column_default": null,
    "column_name": "serial_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Status.",
    "column_default": null,
    "column_name": "status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Product SKU.",
    "column_default": null,
    "column_name": "stock_keeping_unit",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Usage End Date.",
    "column_default": null,
    "column_name": "usage_end_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Asset.",
    "table_name": "salesforce_asset",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Actual Cost in Campaign.",
    "column_default": null,
    "column_name": "actual_cost",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Value Opportunities in Campaign.",
    "column_default": null,
    "column_name": "amount_all_opportunities",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Value Won Opportunities in Campaign.",
    "column_default": null,
    "column_name": "amount_won_opportunities",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Budgeted Cost in Campaign.",
    "column_default": null,
    "column_name": "budgeted_cost",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Record Type ID.",
    "column_default": null,
    "column_name": "campaign_member_record_type_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created By ID.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Description.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The End Date.",
    "column_default": null,
    "column_name": "end_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Expected Response (%).",
    "column_default": null,
    "column_name": "expected_response",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Expected Revenue in Campaign.",
    "column_default": null,
    "column_name": "expected_revenue",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Campaign ID.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Active.",
    "column_default": null,
    "column_name": "is_active",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Activity.",
    "column_default": null,
    "column_name": "last_activity_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified By ID.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Name.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contacts in Campaign.",
    "column_default": null,
    "column_name": "number_of_contacts",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Converted Leads in Campaign.",
    "column_default": null,
    "column_name": "number_of_converted_leads",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Leads in Campaign.",
    "column_default": null,
    "column_name": "number_of_leads",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Opportunities in Campaign.",
    "column_default": null,
    "column_name": "number_of_opportunities",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Responses in Campaign.",
    "column_default": null,
    "column_name": "number_of_responses",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Won Opportunities in Campaign.",
    "column_default": null,
    "column_name": "number_of_won_opportunities",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Num Sent in Campaign.",
    "column_default": null,
    "column_name": "number_sent",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Owner ID.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Parent Campaign ID.",
    "column_default": null,
    "column_name": "parent_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Start Date.",
    "column_default": null,
    "column_name": "start_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Status.",
    "column_default": null,
    "column_name": "status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Type.",
    "column_default": null,
    "column_name": "type",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Campaign.",
    "table_name": "salesforce_campaign",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Account ID.",
    "column_default": null,
    "column_name": "account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Asset ID.",
    "column_default": null,
    "column_name": "asset_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Case Number.",
    "column_default": null,
    "column_name": "case_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Closed Date.",
    "column_default": null,
    "column_name": "closed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Internal Comments.",
    "column_default": null,
    "column_name": "comments",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact Email.",
    "column_default": null,
    "column_name": "contact_email",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact Fax.",
    "column_default": null,
    "column_name": "contact_fax",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact ID.",
    "column_default": null,
    "column_name": "contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact Mobile.",
    "column_default": null,
    "column_name": "contact_mobile",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact Phone.",
    "column_default": null,
    "column_name": "contact_phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created By ID.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Description.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Engineering Req Number.",
    "column_default": null,
    "column_name": "engineering_req_number__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Case ID.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Closed.",
    "column_default": null,
    "column_name": "is_closed",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Escalated.",
    "column_default": null,
    "column_name": "is_escalated",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified By ID.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Case Origin.",
    "column_default": null,
    "column_name": "origin",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Owner ID.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Parent Case ID.",
    "column_default": null,
    "column_name": "parent_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Potential Liability.",
    "column_default": null,
    "column_name": "potential_liability__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Priority.",
    "column_default": null,
    "column_name": "priority",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Product.",
    "column_default": null,
    "column_name": "product__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Case Reason.",
    "column_default": null,
    "column_name": "reason",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SLA Violation.",
    "column_default": null,
    "column_name": "sla_violation__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Source ID.",
    "column_default": null,
    "column_name": "source_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Status.",
    "column_default": null,
    "column_name": "status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Subject.",
    "column_default": null,
    "column_name": "subject",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Company.",
    "column_default": null,
    "column_name": "supplied_company",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Address.",
    "column_default": null,
    "column_name": "supplied_email",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Name.",
    "column_default": null,
    "column_name": "supplied_name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Phone.",
    "column_default": null,
    "column_name": "supplied_phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Case Type.",
    "column_default": null,
    "column_name": "type",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Case.",
    "table_name": "salesforce_case",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Account ID.",
    "column_default": null,
    "column_name": "account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Assistant's Name.",
    "column_default": null,
    "column_name": "assistant_name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Asst. Phone.",
    "column_default": null,
    "column_name": "assistant_phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Birthdate.",
    "column_default": null,
    "column_name": "birthdate",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Clean Status.",
    "column_default": null,
    "column_name": "clean_status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created By ID.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Department.",
    "column_default": null,
    "column_name": "department",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact Description.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email.",
    "column_default": null,
    "column_name": "email",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Bounced Date.",
    "column_default": null,
    "column_name": "email_bounced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Bounced Reason.",
    "column_default": null,
    "column_name": "email_bounced_reason",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Business Fax.",
    "column_default": null,
    "column_name": "fax",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Home Phone.",
    "column_default": null,
    "column_name": "home_phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact ID.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Individual ID.",
    "column_default": null,
    "column_name": "individual_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Is Email Bounced.",
    "column_default": null,
    "column_name": "is_email_bounced",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Data.com Key.",
    "column_default": null,
    "column_name": "jigsaw",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Jigsaw Contact ID.",
    "column_default": null,
    "column_name": "jigsaw_contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Languages.",
    "column_default": null,
    "column_name": "languages__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Activity.",
    "column_default": null,
    "column_name": "last_activity_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Stay-in-Touch Request Date.",
    "column_default": null,
    "column_name": "last_cu_request_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Stay-in-Touch Save Date.",
    "column_default": null,
    "column_name": "last_cu_update_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified By ID.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Lead Source.",
    "column_default": null,
    "column_name": "lead_source",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Level.",
    "column_default": null,
    "column_name": "level__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Mailing Address.",
    "column_default": null,
    "column_name": "mailing_address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Master Record ID.",
    "column_default": null,
    "column_name": "master_record_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Mobile Phone.",
    "column_default": null,
    "column_name": "mobile_phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Full Name.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Other Address.",
    "column_default": null,
    "column_name": "other_address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Other Phone.",
    "column_default": null,
    "column_name": "other_phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Owner ID.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Business Phone.",
    "column_default": null,
    "column_name": "phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Photo URL.",
    "column_default": null,
    "column_name": "photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Reports To ID.",
    "column_default": null,
    "column_name": "reports_to_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Title.",
    "column_default": null,
    "column_name": "title",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Contact.",
    "table_name": "salesforce_contact",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the Account associated with this contract.",
    "column_default": null,
    "column_name": "account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the User who activated this contract.",
    "column_default": null,
    "column_name": "activated_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date and time when this contract was activated.",
    "column_default": null,
    "column_name": "activated_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Billing address of the account.",
    "column_default": null,
    "column_name": "billing_address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date on which the contract was signed by organization.",
    "column_default": null,
    "column_name": "company_signed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the User who signed the contract.",
    "column_default": null,
    "column_name": "company_signed_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Number of the contract.",
    "column_default": null,
    "column_name": "contract_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Number of months that the contract is valid.",
    "column_default": null,
    "column_name": "contract_term",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who created the contract record.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date and time when contract record was created.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date on which the customer signed the contract.",
    "column_default": null,
    "column_name": "customer_signed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the Contact who signed this contract.",
    "column_default": null,
    "column_name": "customer_signed_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Title of the contact who signed the contract.",
    "column_default": null,
    "column_name": "customer_signed_title",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Statement describing the contract.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Calculated end date of the contract. This value is calculated by adding the ContractTerm to the start_date.",
    "column_default": null,
    "column_name": "end_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the contract in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether the object has been moved to the Recycle Bin (true) or not (false).",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Value is one of the following, whichever is the most recent. a) Due date of the most recent event logged against the record. b) Due date of the most recently closed task associated with the record.",
    "column_default": null,
    "column_name": "last_activity_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Last date the contract was approved.",
    "column_default": null,
    "column_name": "last_approved_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The id of user who most recently changed the contract record.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The date and time of the last change to contract record.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The timestamp when the current user last accessed this record, a record related to this record, or a list view.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The timestamp when the current user last viewed this record or list view. If this value is null, the user might have only accessed this record or list view (last_referenced_date) but not viewed it.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Number of days ahead of the contract end date (15, 30, 45, 60, 90, and 120). Used to notify the owner in advance that the contract is ending.",
    "column_default": null,
    "column_name": "owner_expiration_notice",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the user who owns the contract.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the pricebook, if any, associated with this contract.",
    "column_default": null,
    "column_name": "pricebook_2_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Special terms that apply to the contract.",
    "column_default": null,
    "column_name": "special_terms",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Start date for this contract.",
    "column_default": null,
    "column_name": "start_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The picklist of values that indicate order status. A contract can be Draft, InApproval, or Activated.",
    "column_default": null,
    "column_name": "status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Status Category.",
    "column_default": null,
    "column_name": "status_code",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The date and time when contract was last modified by a user or by an automated process.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a contract (a business agreement) associated with an Account.",
    "table_name": "salesforce_contract",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created By ID.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Record ID.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified By ID.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Custom APP Name.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Owner ID.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Custom APP.",
    "table_name": "salesforce_custom_app__c",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Street address for the lead.",
    "column_default": null,
    "column_name": "address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Annual revenue for the lead's company.",
    "column_default": null,
    "column_name": "annual_revenue",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Clean Status.",
    "column_default": null,
    "column_name": "clean_status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The lead's company.",
    "column_default": null,
    "column_name": "company",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Company D-U-N-S Number.",
    "column_default": null,
    "column_name": "company_duns_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Converted Account ID.",
    "column_default": null,
    "column_name": "converted_account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Converted Contact ID.",
    "column_default": null,
    "column_name": "converted_contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date on which this lead was converted.",
    "column_default": null,
    "column_name": "converted_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Converted Opportunity ID.",
    "column_default": null,
    "column_name": "converted_opportunity_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who created the lead.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Creation date and time of the lead.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Current Generator(s).",
    "column_default": null,
    "column_name": "current_generators__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The D&B Company ID.",
    "column_default": null,
    "column_name": "dandb_company_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Description.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The lead's email address.",
    "column_default": null,
    "column_name": "email",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Bounced Date.",
    "column_default": null,
    "column_name": "email_bounced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Bounced Reason.",
    "column_default": null,
    "column_name": "email_bounced_reason",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Fax.",
    "column_default": null,
    "column_name": "fax",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the lead in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Individual ID.",
    "column_default": null,
    "column_name": "individual_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Primary business of lead's company.",
    "column_default": null,
    "column_name": "industry",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether the lead has been converted (true) or not (false).",
    "column_default": null,
    "column_name": "is_converted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Unread By Owner.",
    "column_default": null,
    "column_name": "is_unread_by_owner",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Data.com Key.",
    "column_default": null,
    "column_name": "jigsaw",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Jigsaw Contact ID.",
    "column_default": null,
    "column_name": "jigsaw_contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Activity.",
    "column_default": null,
    "column_name": "last_activity_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who last changed the lead record.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date and time of the last changes to lead record.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Source of lead, for example, Advertisement, Partner, or Web.",
    "column_default": null,
    "column_name": "lead_source",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Master Record ID.",
    "column_default": null,
    "column_name": "master_record_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Mobile Phone.",
    "column_default": null,
    "column_name": "mobile_phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Name of the lead, as displayed on lead detail page.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Number of employees at the lead's company.",
    "column_default": null,
    "column_name": "number_of_employees",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Number of Locations.",
    "column_default": null,
    "column_name": "numberof_locations__c",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the assigned owner of the lead.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Lead's primary phone number.",
    "column_default": null,
    "column_name": "phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Photo URL.",
    "column_default": null,
    "column_name": "photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Primary.",
    "column_default": null,
    "column_name": "primary__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Product Interest.",
    "column_default": null,
    "column_name": "product_interest__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates value or prospect of the lead, for example, Hot, Warm, Cold.",
    "column_default": null,
    "column_name": "rating",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SIC Code.",
    "column_default": null,
    "column_name": "sic_code__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Status of the lead, for example, Open, Contacted, Qualified.",
    "column_default": null,
    "column_name": "status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Title.",
    "column_default": null,
    "column_name": "title",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "URL of the lead's company's website.",
    "column_default": null,
    "column_name": "website",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a prospect or lead.",
    "table_name": "salesforce_lead",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the account associated with this opportunity.",
    "column_default": null,
    "column_name": "account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Estimated total sale amount. For opportunities with products, the amount is the sum of the related products.",
    "column_default": null,
    "column_name": "amount",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of a related Campaign. This field is defined only for those organizations that have the campaign feature Campaigns enabled.",
    "column_default": null,
    "column_name": "campaign_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date when the opportunity is expected to close.",
    "column_default": null,
    "column_name": "close_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who created the opportunity.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The creation date and time of the opportunity.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Current Generator(s).",
    "column_default": null,
    "column_name": "current_generators__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Delivery/Installation Status.",
    "column_default": null,
    "column_name": "delivery_installation_status__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Description of the opportunity.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Calculated revenue based on the Amount and Probability fields.",
    "column_default": null,
    "column_name": "expected_revenue",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Fiscal Period.",
    "column_default": null,
    "column_name": "fiscal",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Forecast category name displayed in reports, opportunity detail and edit pages, opportunity searches, and opportunity list views.",
    "column_default": null,
    "column_name": "forecast_category",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Name of the forecast category.",
    "column_default": null,
    "column_name": "forecast_category_name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether an opportunity has an open event or task (true) or not (false).",
    "column_default": null,
    "column_name": "has_open_activity",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether the opportunity has associated line items. A value of true means that Opportunity line items have been created for the opportunity.",
    "column_default": null,
    "column_name": "has_opportunity_line_item",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether an opportunity has an overdue task (true) or not (false).",
    "column_default": null,
    "column_name": "has_overdue_task",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the opportunity in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates that the opportunity is closed.",
    "column_default": null,
    "column_name": "is_closed",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates that the opportunity is deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates that the opportunity is private.",
    "column_default": null,
    "column_name": "is_private",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates that the quote or proposal has been signed or electronically accepted.",
    "column_default": null,
    "column_name": "is_won",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Value is one of the following, whichever is the most recent of a) Due date of the most recent event logged against the record or b) Due date of the most recently closed task associated with the record.",
    "column_default": null,
    "column_name": "last_activity_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The id of the user who last modified the oppurtinity record.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The data and time of the last modification of the oppurtinity record.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Source of this opportunity, such as Advertisement or Trade Show.",
    "column_default": null,
    "column_name": "lead_source",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Main Competitor(s).",
    "column_default": null,
    "column_name": "main_competitors__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "A name for this opportunity.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Description of next task in closing opportunity.",
    "column_default": null,
    "column_name": "next_step",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Order Number.",
    "column_default": null,
    "column_name": "order_number__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the User who has been assigned to work this opportunity.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of a related Pricebook2 object. The Pricebook2Id field indicates which Pricebook2 applies to this opportunity. The Pricebook2Id field is defined only for those organizations that have products enabled as a feature.",
    "column_default": null,
    "column_name": "pricebook_2_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Percentage of estimated confidence in closing the opportunity.",
    "column_default": null,
    "column_name": "probability",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Current stage of opportunity.",
    "column_default": null,
    "column_name": "stage_name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The date and time when opportunity was last modified by a user or by an automated process.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Number of items included in this opportunity. Used in quantity-based forecasting.",
    "column_default": null,
    "column_name": "total_opportunity_quantity",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Tracking Number.",
    "column_default": null,
    "column_name": "tracking_number__c",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Type of opportunity, such as Existing Business or New Business.",
    "column_default": null,
    "column_name": "type",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an opportunity, which is a sale or pending deal.",
    "table_name": "salesforce_opportunity",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of an associated Contact.",
    "column_default": null,
    "column_name": "contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who created contact role record.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date and time of the creation of the contact role record.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the opportunity contact role in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether the associated Contact plays the primary role on the Opportunity (true) or not (false).",
    "column_default": null,
    "column_name": "is_primary",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who most recently changed the contact role record.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date of most recent change in the contact role record.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Required. ID of an associated Opportunity.",
    "column_default": null,
    "column_name": "opportunity_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Name of the role played by the associated Contact on the Opportunity, such as Business User or Decision Maker.",
    "column_default": null,
    "column_name": "role",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents the role that a Contact plays on an Opportunity.",
    "table_name": "salesforce_opportunity_contact_role",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the Account associated with this order.",
    "column_default": null,
    "column_name": "account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the User who activated this order.",
    "column_default": null,
    "column_name": "activated_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date and time when the order was activated.",
    "column_default": null,
    "column_name": "activated_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the contact that the order is billed to.",
    "column_default": null,
    "column_name": "bill_to_contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The billing adress for the order.",
    "column_default": null,
    "column_name": "billing_address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the user who authorized the account associated with the order.",
    "column_default": null,
    "column_name": "company_authorized_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The date on which your organization authorized the order.",
    "column_default": null,
    "column_name": "company_authorized_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the contract associated with this order. Can only be updated when the order's StatusCode value is Draft.",
    "column_default": null,
    "column_name": "contract_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who created the order record.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Creation date and time of the order record.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the contact who authorized the order.",
    "column_default": null,
    "column_name": "customer_authorized_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date on which the contact authorized the order.",
    "column_default": null,
    "column_name": "customer_authorized_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Description of the order.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date at which the order becomes effective.",
    "column_default": null,
    "column_name": "effective_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date at which the order ends.",
    "column_default": null,
    "column_name": "end_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the order in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates that the order is deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Determines whether an order is a reduction order.",
    "column_default": null,
    "column_name": "is_reduction_order",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who most recently changed the order record.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date of most recent change in the order record.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The timestamp when the current user last accessed this record, a record related to this record, or a list view.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The timestamp when the current user last viewed this record or list view. If this value is null, the user might have only accessed this record or list view (LastReferencedDate) but not viewed it.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Title for the order that distinguishes it from other orders.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Order number assigned to this order.",
    "column_default": null,
    "column_name": "order_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Reference number assigned to the order.",
    "column_default": null,
    "column_name": "order_reference_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Optional. ID of the original order that a reduction order is reducing, if the reduction order is reducing a single order.",
    "column_default": null,
    "column_name": "original_order_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": " ID of the User or queue that owns this order.",
    "column_default": null,
    "column_name": "owner_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date of the purchase order.",
    "column_default": null,
    "column_name": "po_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Number identifying the purchase order.",
    "column_default": null,
    "column_name": "po_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the price book associated with this order.",
    "column_default": null,
    "column_name": "pricebook_2_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the contact that the order is shipped to.",
    "column_default": null,
    "column_name": "ship_to_contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The shipping adress for the order.",
    "column_default": null,
    "column_name": "shipping_address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Status field specifies the current state of an order. Status strings represent its current state (Draft or Activated).",
    "column_default": null,
    "column_name": "status",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Status code of the stage that the order has reached in the order business process.",
    "column_default": null,
    "column_name": "status_code",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The date and time when order record was last modified by a user or by an automated process.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Total amount of the order.",
    "column_default": null,
    "column_name": "total_amount",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Type of order.",
    "column_default": null,
    "column_name": "type",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents an order associated with a contract or an account.",
    "table_name": "salesforce_order",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created By ID.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Description.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Price Book ID.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Active.",
    "column_default": null,
    "column_name": "is_active",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Archived.",
    "column_default": null,
    "column_name": "is_archived",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Is Standard Price Book.",
    "column_default": null,
    "column_name": "is_standard",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified By ID.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Price Book Name.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book.",
    "table_name": "salesforce_pricebook",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created By ID.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Price Book Entry ID.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Active.",
    "column_default": null,
    "column_name": "is_active",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Deleted.",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified By ID.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Product Name.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Price Book ID.",
    "column_default": null,
    "column_name": "pricebook_2_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Product ID.",
    "column_default": null,
    "column_name": "product_2_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Product Code.",
    "column_default": null,
    "column_name": "product_code",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The List Price.",
    "column_default": null,
    "column_name": "unit_price",
    "data_type": "double precision",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Use Standard Price.",
    "column_default": null,
    "column_name": "use_standard_price",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Price Book Entry.",
    "table_name": "salesforce_pricebook_entry",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The id of the user who created the product record, with the date and time of creation.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date and time of the creation of the product role record.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Description of the product.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "URL leading to a specific version of a record in the linked external data source.",
    "column_default": null,
    "column_name": "display_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The id of the related external data source.",
    "column_default": null,
    "column_name": "external_data_source_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The unique identifier of a record in the linked external data source.",
    "column_default": null,
    "column_name": "external_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Name of the product family associated with this record.",
    "column_default": null,
    "column_name": "family",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the product in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates that the product is ready for use in a price book, opportunity, or quote, and whether you can see the product in views.",
    "column_default": null,
    "column_name": "is_active",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Describes whether the product is archived. The default value is false.",
    "column_default": null,
    "column_name": "is_archived",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether the object has been moved to the Recycle Bin (true) or not (false).",
    "column_default": null,
    "column_name": "is_deleted",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who most recently changed the product record.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Date of most recent change in the product record.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The timestamp when the current user last viewed product record.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The timestamp when the current user last viewed this record. If this value is null, this record might only have been referenced (last_referenced_date) and not viewed by the current user.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The product's name.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The internal code or product number that you use to identify the product.",
    "column_default": null,
    "column_name": "product_code",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unit of the product—for example, kilograms, liters, or cases.",
    "column_default": null,
    "column_name": "quantity_unit_of_measure",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The product's SKU, which can be used with or in place of the Product Code field.",
    "column_default": null,
    "column_name": "stock_keeping_unit",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The date and time when order record was last modified by a user or by an automated process.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a product that organization sells.",
    "table_name": "salesforce_product",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created By ID.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Description.",
    "column_default": null,
    "column_name": "description",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Profile ID.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified By ID.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Name.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Access Experience Management.",
    "column_default": null,
    "column_name": "permissions_access_cmc",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Activate Contracts.",
    "column_default": null,
    "column_name": "permissions_activate_contract",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Activate Orders.",
    "column_default": null,
    "column_name": "permissions_activate_order",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Add Tableau CRM Remote Connections.",
    "column_default": null,
    "column_name": "permissions_add_analytics_remote_connections",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Add People to Direct Messages.",
    "column_default": null,
    "column_name": "permissions_add_direct_message_members",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email-Based Identity Verification Option.",
    "column_default": null,
    "column_name": "permissions_allow_email_ic",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Lightning Login User.",
    "column_default": null,
    "column_name": "permissions_allow_lightning_login",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Knowledge One.",
    "column_default": null,
    "column_name": "permissions_allow_universal_search",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View and Edit Converted Leads.",
    "column_default": null,
    "column_name": "permissions_allow_view_edit_converted_leads",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Allow View Knowledge.",
    "column_default": null,
    "column_name": "permissions_allow_view_knowledge",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Apex REST Services.",
    "column_default": null,
    "column_name": "permissions_apex_rest_services",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The API Enabled.",
    "column_default": null,
    "column_name": "permissions_api_enabled",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Archive Articles.",
    "column_default": null,
    "column_name": "permissions_archive_articles",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Assign Permission Sets.",
    "column_default": null,
    "column_name": "permissions_assign_permission_sets",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Assign Topics.",
    "column_default": null,
    "column_name": "permissions_assign_topics",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Assign Chat Skills to Users.",
    "column_default": null,
    "column_name": "permissions_assign_user_to_skill",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Author Apex.",
    "column_default": null,
    "column_name": "permissions_author_apex",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Use Einstein Activity Capture.",
    "column_default": null,
    "column_name": "permissions_automatic_activity_capture",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Bulk API Hard Delete.",
    "column_default": null,
    "column_name": "permissions_bulk_api_hard_delete",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Run Macros on Multiple Records.",
    "column_default": null,
    "column_name": "permissions_bulk_macros_allowed",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Bypass Email Approval.",
    "column_default": null,
    "column_name": "permissions_bypass_email_approval",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Campaign Influence.",
    "column_default": null,
    "column_name": "permissions_campaign_influence_2",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Can Approve Feed Post and Comment.",
    "column_default": null,
    "column_name": "permissions_can_approve_feed_post",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Dataset Recipes.",
    "column_default": null,
    "column_name": "permissions_can_edit_data_prep_recipe",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Insert System Field Values for Chatter Feeds.",
    "column_default": null,
    "column_name": "permissions_can_insert_feed_system_fields",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Drag-and-Drop Dashboard Builder.",
    "column_default": null,
    "column_name": "permissions_can_use_new_dashboard_builder",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Verify Answers to Chatter Questions.",
    "column_default": null,
    "column_name": "permissions_can_verify_comment",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Change Dashboard Colors.",
    "column_default": null,
    "column_name": "permissions_change_dashboard_colors",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Allow Inclusion of Code Snippets from UI.",
    "column_default": null,
    "column_name": "permissions_chatter_compose_ui_codesnippet",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit My Own Posts.",
    "column_default": null,
    "column_name": "permissions_chatter_edit_own_post",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Posts on Records I Own.",
    "column_default": null,
    "column_name": "permissions_chatter_edit_own_record_post",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create Public Links.",
    "column_default": null,
    "column_name": "permissions_chatter_file_link",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Chatter Internal User.",
    "column_default": null,
    "column_name": "permissions_chatter_internal_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Invite Customers To Chatter.",
    "column_default": null,
    "column_name": "permissions_chatter_invite_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create and Own New Chatter Groups.",
    "column_default": null,
    "column_name": "permissions_chatter_own_groups",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Close Conversation Threads.",
    "column_default": null,
    "column_name": "permissions_close_conversations",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Configure Custom Recommendations.",
    "column_default": null,
    "column_name": "permissions_config_custom_recs",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Configure Messaging.",
    "column_default": null,
    "column_name": "permissions_configure_live_message",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Connect Organization to Environment Hub.",
    "column_default": null,
    "column_name": "permissions_connect_org_to_environment_hub",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Salesforce CRM Content.",
    "column_default": null,
    "column_name": "permissions_content_administrator",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Files Connect Cloud.",
    "column_default": null,
    "column_name": "permissions_content_hub_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Access Libraries.",
    "column_default": null,
    "column_name": "permissions_content_workspaces",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Convert Leads.",
    "column_default": null,
    "column_name": "permissions_convert_leads",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create and Customize Dashboards.",
    "column_default": null,
    "column_name": "permissions_create_customize_dashboards",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create and Customize List Views.",
    "column_default": null,
    "column_name": "permissions_create_customize_filters",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create and Customize Reports.",
    "column_default": null,
    "column_name": "permissions_create_customize_reports",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create Dashboard Folders.",
    "column_default": null,
    "column_name": "permissions_create_dashboard_folders",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create AppExchange Packages.",
    "column_default": null,
    "column_name": "permissions_create_multiforce",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create Report Folders.",
    "column_default": null,
    "column_name": "permissions_create_report_folders",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Report Builder (Lightning Experience).",
    "column_default": null,
    "column_name": "permissions_create_report_in_lightning",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create Topics.",
    "column_default": null,
    "column_name": "permissions_create_topics",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create Libraries.",
    "column_default": null,
    "column_name": "permissions_create_workspaces",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Access Custom Mobile Apps.",
    "column_default": null,
    "column_name": "permissions_custom_mobile_apps_access",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Show Custom Sidebar On All Pages.",
    "column_default": null,
    "column_name": "permissions_custom_sidebar_on_all_pages",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Customize Application.",
    "column_default": null,
    "column_name": "permissions_customize_application",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Weekly Data Export.",
    "column_default": null,
    "column_name": "permissions_data_export",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Multi-Factor Authentication in User Interface.",
    "column_default": null,
    "column_name": "permissions_delegated_two_factor",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Delete Activated Contracts.",
    "column_default": null,
    "column_name": "permissions_delete_activated_contract",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Delete Topics.",
    "column_default": null,
    "column_name": "permissions_delete_topics",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create Content Deliveries.",
    "column_default": null,
    "column_name": "permissions_distribute_from_pers_wksp",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Activated Orders.",
    "column_default": null,
    "column_name": "permissions_edit_activated_orders",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Letterheads.",
    "column_default": null,
    "column_name": "permissions_edit_brand_templates",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Case Comments.",
    "column_default": null,
    "column_name": "permissions_edit_case_comments",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Events.",
    "column_default": null,
    "column_name": "permissions_edit_event",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit HTML Templates.",
    "column_default": null,
    "column_name": "permissions_edit_html_templates",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Articles.",
    "column_default": null,
    "column_name": "permissions_edit_knowledge",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit My Dashboards.",
    "column_default": null,
    "column_name": "permissions_edit_my_dashboards",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit My Reports.",
    "column_default": null,
    "column_name": "permissions_edit_my_reports",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Opportunity Product Sales Price.",
    "column_default": null,
    "column_name": "permissions_edit_opp_line_item_unit_price",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Public Documents.",
    "column_default": null,
    "column_name": "permissions_edit_public_documents",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Public List Views.",
    "column_default": null,
    "column_name": "permissions_edit_public_filters",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Public Classic Email Templates.",
    "column_default": null,
    "column_name": "permissions_edit_public_templates",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Read Only Fields.",
    "column_default": null,
    "column_name": "permissions_edit_readonly_fields",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Tasks.",
    "column_default": null,
    "column_name": "permissions_edit_task",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Topics.",
    "column_default": null,
    "column_name": "permissions_edit_topics",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Article Translation - Edit.",
    "column_default": null,
    "column_name": "permissions_edit_translation",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Administration.",
    "column_default": null,
    "column_name": "permissions_email_administration",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Mass Email.",
    "column_default": null,
    "column_name": "permissions_email_mass",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Send Email.",
    "column_default": null,
    "column_name": "permissions_email_single",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Email Templates.",
    "column_default": null,
    "column_name": "permissions_email_template_management",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Show App Launcher in Experience Cloud Sites.",
    "column_default": null,
    "column_name": "permissions_enable_community_app_launcher",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Send Outbound Messages.",
    "column_default": null,
    "column_name": "permissions_enable_notifications",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Export Reports.",
    "column_default": null,
    "column_name": "permissions_export_report",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Pin Posts in Feeds.",
    "column_default": null,
    "column_name": "permissions_feed_pinning",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Require Flow User Feature License.",
    "column_default": null,
    "column_name": "permissions_flow_ufl_required",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Multi-Factor Authentication for User Interface Logins.",
    "column_default": null,
    "column_name": "permissions_force_two_factor",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Experiences.",
    "column_default": null,
    "column_name": "permissions_govern_networks",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Hide the Seen By List.",
    "column_default": null,
    "column_name": "permissions_hide_read_by_list",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Use Identity Connect.",
    "column_default": null,
    "column_name": "permissions_identity_connect",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Use Identity Features.",
    "column_default": null,
    "column_name": "permissions_identity_enabled",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Import Custom Objects.",
    "column_default": null,
    "column_name": "permissions_import_custom_objects",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Import Leads.",
    "column_default": null,
    "column_name": "permissions_import_leads",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Import Personal Contacts.",
    "column_default": null,
    "column_name": "permissions_import_personal",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Tableau CRM.",
    "column_default": null,
    "column_name": "permissions_insights_app_admin",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create and Edit Tableau CRM Dashboards.",
    "column_default": null,
    "column_name": "permissions_insights_app_dashboard_editor",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Edit Tableau CRM Dataflows.",
    "column_default": null,
    "column_name": "permissions_insights_app_elt_editor",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Upload External Data to Tableau CRM.",
    "column_default": null,
    "column_name": "permissions_insights_app_upload_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Use Tableau CRM.",
    "column_default": null,
    "column_name": "permissions_insights_app_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create Tableau CRM Apps.",
    "column_default": null,
    "column_name": "permissions_insights_create_application",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Download AppExchange Packages.",
    "column_default": null,
    "column_name": "permissions_install_multiforce",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Lightning Console User.",
    "column_default": null,
    "column_name": "permissions_lightning_console_allowed_for_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Lightning Experience User.",
    "column_default": null,
    "column_name": "permissions_lightning_experience_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Allow sending of List Emails.",
    "column_default": null,
    "column_name": "permissions_list_email_send",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Messaging Agent.",
    "column_default": null,
    "column_name": "permissions_live_message_agent",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Remain in Salesforce Classic.",
    "column_default": null,
    "column_name": "permissions_ltng_promo_reserved_01_user_perm",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Reporting Snapshots.",
    "column_default": null,
    "column_name": "permissions_manage_analytic_snapshots",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Auth. Providers.",
    "column_default": null,
    "column_name": "permissions_manage_auth_providers",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Business Hours Holidays.",
    "column_default": null,
    "column_name": "permissions_manage_business_hour_holidays",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Call Centers.",
    "column_default": null,
    "column_name": "permissions_manage_call_centers",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Cases.",
    "column_default": null,
    "column_name": "permissions_manage_cases",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Categories.",
    "column_default": null,
    "column_name": "permissions_manage_categories",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Certificates.",
    "column_default": null,
    "column_name": "permissions_manage_certificates",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Chatter Messages and Direct Messages.",
    "column_default": null,
    "column_name": "permissions_manage_chatter_messages",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Content Permissions.",
    "column_default": null,
    "column_name": "permissions_manage_content_permissions",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Content Properties.",
    "column_default": null,
    "column_name": "permissions_manage_content_properties",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage record types and layouts for Files.",
    "column_default": null,
    "column_name": "permissions_manage_content_types",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Custom Permissions.",
    "column_default": null,
    "column_name": "permissions_manage_custom_permissions",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Custom Report Types.",
    "column_default": null,
    "column_name": "permissions_manage_custom_report_types",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Dashboards in Public Folders.",
    "column_default": null,
    "column_name": "permissions_manage_dashbds_in_pub_folders",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Data Categories.",
    "column_default": null,
    "column_name": "permissions_manage_data_categories",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Data Integrations.",
    "column_default": null,
    "column_name": "permissions_manage_data_integrations",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Dynamic Dashboards.",
    "column_default": null,
    "column_name": "permissions_manage_dynamic_dashboards",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Email Client Configurations.",
    "column_default": null,
    "column_name": "permissions_manage_email_client_config",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Encryption Keys.",
    "column_default": null,
    "column_name": "permissions_manage_encryption_keys",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Entitlements.",
    "column_default": null,
    "column_name": "permissions_manage_entitlements",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Lightning Sync.",
    "column_default": null,
    "column_name": "permissions_manage_exchange_config",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Health Check.",
    "column_default": null,
    "column_name": "permissions_manage_health_check",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Flow.",
    "column_default": null,
    "column_name": "permissions_manage_interaction",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Internal Users.",
    "column_default": null,
    "column_name": "permissions_manage_internal_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage IP Addresses.",
    "column_default": null,
    "column_name": "permissions_manage_ip_addresses",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Salesforce Knowledge.",
    "column_default": null,
    "column_name": "permissions_manage_knowledge",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Knowledge Article Import/Export.",
    "column_default": null,
    "column_name": "permissions_manage_knowledge_import_export",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Leads.",
    "column_default": null,
    "column_name": "permissions_manage_leads",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Login Access Policies.",
    "column_default": null,
    "column_name": "permissions_manage_login_access_policies",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Mobile Configurations.",
    "column_default": null,
    "column_name": "permissions_manage_mobile",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create and Set Up Experiences.",
    "column_default": null,
    "column_name": "permissions_manage_networks",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Password Policies.",
    "column_default": null,
    "column_name": "permissions_manage_password_policies",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Profiles and Permission Sets.",
    "column_default": null,
    "column_name": "permissions_manage_profiles_permissionsets",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Next Best Action Recommendations.",
    "column_default": null,
    "column_name": "permissions_manage_propositions",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage All Private Reports and Dashboards.",
    "column_default": null,
    "column_name": "permissions_manage_pvt_rpts_and_dashbds",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Next Best Action Strategies.",
    "column_default": null,
    "column_name": "permissions_manage_recommendation_strategies",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Connected Apps.",
    "column_default": null,
    "column_name": "permissions_manage_remote_access",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Reports in Public Folders.",
    "column_default": null,
    "column_name": "permissions_manage_reports_in_pub_folders",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Roles.",
    "column_default": null,
    "column_name": "permissions_manage_roles",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Promoted Search Terms.",
    "column_default": null,
    "column_name": "permissions_manage_search_promotion_rules",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Session Permission Set Activations.",
    "column_default": null,
    "column_name": "permissions_manage_session_permission_sets",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Sharing.",
    "column_default": null,
    "column_name": "permissions_manage_sharing",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Published Solutions.",
    "column_default": null,
    "column_name": "permissions_manage_solutions",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Surveys.",
    "column_default": null,
    "column_name": "permissions_manage_surveys",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Synonyms.",
    "column_default": null,
    "column_name": "permissions_manage_synonyms",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Tableau CRM Templated Apps.",
    "column_default": null,
    "column_name": "permissions_manage_templated_app",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Multi-Factor Authentication in API.",
    "column_default": null,
    "column_name": "permissions_manage_two_factor",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Unlisted Groups.",
    "column_default": null,
    "column_name": "permissions_manage_unlisted_groups",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Users.",
    "column_default": null,
    "column_name": "permissions_manage_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Mass Edits from Lists.",
    "column_default": null,
    "column_name": "permissions_mass_inline_edit",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Merge Topics.",
    "column_default": null,
    "column_name": "permissions_merge_topics",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Moderate Chatter.",
    "column_default": null,
    "column_name": "permissions_moderate_chatter",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Moderate Experience Cloud Site Users.",
    "column_default": null,
    "column_name": "permissions_moderate_network_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Modify All Data.",
    "column_default": null,
    "column_name": "permissions_modify_all_data",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Modify Metadata Through Metadata API Functions.",
    "column_default": null,
    "column_name": "permissions_modify_metadata",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Report Builder.",
    "column_default": null,
    "column_name": "permissions_new_report_builder",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Create and Update Second-Generation Packages.",
    "column_default": null,
    "column_name": "permissions_packaging_2",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Password Never Expires.",
    "column_default": null,
    "column_name": "permissions_password_never_expires",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Hide Option to Switch to Salesforce Classic.",
    "column_default": null,
    "column_name": "permissions_prevent_classic_experience",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Publish Articles.",
    "column_default": null,
    "column_name": "permissions_publish_articles",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Upload AppExchange Packages.",
    "column_default": null,
    "column_name": "permissions_publish_multiforce",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Article Translation - Publish.",
    "column_default": null,
    "column_name": "permissions_publish_translation",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Remove People from Direct Messages.",
    "column_default": null,
    "column_name": "permissions_remove_direct_message_members",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Reset User Passwords and Unlock Users.",
    "column_default": null,
    "column_name": "permissions_reset_passwords",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Run Flows.",
    "column_default": null,
    "column_name": "permissions_run_flow",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Run Reports.",
    "column_default": null,
    "column_name": "permissions_run_reports",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Sales Console.",
    "column_default": null,
    "column_name": "permissions_sales_console",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Schedule Reports.",
    "column_default": null,
    "column_name": "permissions_schedule_reports",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Select Files from Salesforce.",
    "column_default": null,
    "column_name": "permissions_select_files_from_salesforce",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Send announcement emails.",
    "column_default": null,
    "column_name": "permissions_send_announcement_emails",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Send Stay-in-Touch Requests.",
    "column_default": null,
    "column_name": "permissions_send_sit_requests",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Share internal Knowledge articles externally.",
    "column_default": null,
    "column_name": "permissions_share_internal_articles",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Show Company Name as Site Role.",
    "column_default": null,
    "column_name": "permissions_show_company_name_as_user_badge",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Import Solutions.",
    "column_default": null,
    "column_name": "permissions_solution_import",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Use Einstein Activity Capture Standard.",
    "column_default": null,
    "column_name": "permissions_std_automatic_activity_capture",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Article Translation - Submit for Translation.",
    "column_default": null,
    "column_name": "permissions_submit_for_translation",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manage Macros Users Can't Undo.",
    "column_default": null,
    "column_name": "permissions_submit_macros_allowed",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Subscribe to Dashboards: Add Recipients.",
    "column_default": null,
    "column_name": "permissions_subscribe_dashboard_to_other_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Subscribe to Reports: Add Recipients.",
    "column_default": null,
    "column_name": "permissions_subscribe_report_to_other_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Subscribe to Reports: Set Running User.",
    "column_default": null,
    "column_name": "permissions_subscribe_reports_run_as_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Subscribe to Dashboards.",
    "column_default": null,
    "column_name": "permissions_subscribe_to_lightning_dashboards",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Subscribe to Reports.",
    "column_default": null,
    "column_name": "permissions_subscribe_to_lightning_reports",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Transfer Cases.",
    "column_default": null,
    "column_name": "permissions_transfer_any_case",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Transfer Record.",
    "column_default": null,
    "column_name": "permissions_transfer_any_entity",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Transfer Leads.",
    "column_default": null,
    "column_name": "permissions_transfer_any_lead",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Multi-Factor Authentication for API Logins.",
    "column_default": null,
    "column_name": "permissions_two_factor_api",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Use Team Reassignment Wizards.",
    "column_default": null,
    "column_name": "permissions_use_team_reassign_wizards",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Use Tableau CRM Templated Apps.",
    "column_default": null,
    "column_name": "permissions_use_templated_app",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Allow Access to Customized Actions.",
    "column_default": null,
    "column_name": "permissions_use_web_link",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View All Activities.",
    "column_default": null,
    "column_name": "permissions_view_all_activities",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View All Data.",
    "column_default": null,
    "column_name": "permissions_view_all_data",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View All Users.",
    "column_default": null,
    "column_name": "permissions_view_all_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Content in Portals.",
    "column_default": null,
    "column_name": "permissions_view_content",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Access to view Data Assessment.",
    "column_default": null,
    "column_name": "permissions_view_data_assessment",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Data Categories in Setup.",
    "column_default": null,
    "column_name": "permissions_view_data_categories",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Real-Time Event Monitoring Data.",
    "column_default": null,
    "column_name": "permissions_view_data_leakage_events",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Encrypted Data.",
    "column_default": null,
    "column_name": "permissions_view_encrypted_data",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Event Log Files.",
    "column_default": null,
    "column_name": "permissions_view_event_log_files",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Health Check.",
    "column_default": null,
    "column_name": "permissions_view_health_check",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Help Link.",
    "column_default": null,
    "column_name": "permissions_view_help_link",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View My Team's Dashboards.",
    "column_default": null,
    "column_name": "permissions_view_my_teams_dashboards",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Access to View-Only Licensed Templates and Apps.",
    "column_default": null,
    "column_name": "permissions_view_only_embedded_app_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Login Forensics Events.",
    "column_default": null,
    "column_name": "permissions_view_platform_events",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Dashboards in Public Folders.",
    "column_default": null,
    "column_name": "permissions_view_public_dashboards",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Reports in Public Folders.",
    "column_default": null,
    "column_name": "permissions_view_public_reports",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Roles and Role Hierarchy.",
    "column_default": null,
    "column_name": "permissions_view_roles",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The View Setup and Configuration.",
    "column_default": null,
    "column_name": "permissions_view_setup",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Download Tableau CRM Data.",
    "column_default": null,
    "column_name": "permissions_wave_tabular_download",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Enable WDC Calibration.",
    "column_default": null,
    "column_name": "permissions_work_calibration_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Enable WDC.",
    "column_default": null,
    "column_name": "permissions_work_dot_com_user_perm",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The User License ID.",
    "column_default": null,
    "column_name": "user_license_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The User Type.",
    "column_default": null,
    "column_name": "user_type",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Salesforce Profile.",
    "table_name": "salesforce_profile",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The About Me.",
    "column_default": null,
    "column_name": "about_me",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the Account associated with a Customer Portal user. This field is null for Salesforce users.",
    "column_default": null,
    "column_name": "account_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Address.",
    "column_default": null,
    "column_name": "address",
    "data_type": "jsonb",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The user's alias. For example, jsmith.",
    "column_default": null,
    "column_name": "alias",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The User Photo badge text overlay.",
    "column_default": null,
    "column_name": "badge_text",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Url for banner photo.",
    "column_default": null,
    "column_name": "banner_photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Call Center ID.",
    "column_default": null,
    "column_name": "call_center_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Nickname.",
    "column_default": null,
    "column_name": "community_nickname",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Company Name.",
    "column_default": null,
    "column_name": "company_name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Contact ID.",
    "column_default": null,
    "column_name": "contact_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who created the user including creation date and time.",
    "column_default": null,
    "column_name": "created_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Created Date.",
    "column_default": null,
    "column_name": "created_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Default Notification Frequency when Joining Groups.",
    "column_default": null,
    "column_name": "default_group_notification_frequency",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Delegated Approver ID.",
    "column_default": null,
    "column_name": "delegated_approver_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The company department associated with the user.",
    "column_default": null,
    "column_name": "department",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Chatter Email Highlights Frequency.",
    "column_default": null,
    "column_name": "digest_frequency",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Division.",
    "column_default": null,
    "column_name": "division",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The user's email address.",
    "column_default": null,
    "column_name": "email",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Encoding.",
    "column_default": null,
    "column_name": "email_encoding_key",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The AutoBcc.",
    "column_default": null,
    "column_name": "email_preferences_auto_bcc",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The AutoBccStayInTouch.",
    "column_default": null,
    "column_name": "email_preferences_auto_bcc_stay_in_touch",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The StayInTouchReminder.",
    "column_default": null,
    "column_name": "email_preferences_stay_in_touch_reminder",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The user's employee number.",
    "column_default": null,
    "column_name": "employee_number",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Extension.",
    "column_default": null,
    "column_name": "extension",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Fax.",
    "column_default": null,
    "column_name": "fax",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SAML Federation ID.",
    "column_default": null,
    "column_name": "federation_identifier",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether the user is enabled as a forecast manager (true) or not (false).",
    "column_default": null,
    "column_name": "forecast_enabled",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Url for full-sized Photo.",
    "column_default": null,
    "column_name": "full_photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Unique identifier of the user in Salesforce.",
    "column_default": null,
    "column_name": "id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Individual ID.",
    "column_default": null,
    "column_name": "individual_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Indicates whether the user has access to log in (true) or not (false).",
    "column_default": null,
    "column_name": "is_active",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Show external indicator.",
    "column_default": null,
    "column_name": "is_ext_indicator_visible",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Has Profile Photo.",
    "column_default": null,
    "column_name": "is_profile_photo_active",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Data.com Monthly Addition Limit.",
    "column_default": null,
    "column_name": "jigsaw_import_limit_override",
    "data_type": "bigint",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Language.",
    "column_default": null,
    "column_name": "language_locale_key",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The date and time when the user last successfully logged in. This value is updated if 60 seconds elapses since the user's last login.",
    "column_default": null,
    "column_name": "last_login_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Id of the user who last changed the user fields, including modification date and time.",
    "column_default": null,
    "column_name": "last_modified_by_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Modified Date.",
    "column_default": null,
    "column_name": "last_modified_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Password Change or Reset.",
    "column_default": null,
    "column_name": "last_password_change_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Referenced Date.",
    "column_default": null,
    "column_name": "last_referenced_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Last Viewed Date.",
    "column_default": null,
    "column_name": "last_viewed_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Locale.",
    "column_default": null,
    "column_name": "locale_sid_key",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Manager ID.",
    "column_default": null,
    "column_name": "manager_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Url for Android banner photo.",
    "column_default": null,
    "column_name": "medium_banner_photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Url for medium profile photo.",
    "column_default": null,
    "column_name": "medium_photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Mobile.",
    "column_default": null,
    "column_name": "mobile_phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Display name of the user.",
    "column_default": null,
    "column_name": "name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Sales Anywhere Trial Expiration Date.",
    "column_default": null,
    "column_name": "offline_pda_trial_expiration_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Offline Edition Trial Expiration Date.",
    "column_default": null,
    "column_name": "offline_trial_expiration_date",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Out of office message.",
    "column_default": null,
    "column_name": "out_of_office_message",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Phone.",
    "column_default": null,
    "column_name": "phone",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "ID of the user's Profile.",
    "column_default": null,
    "column_name": "profile_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Admin Info Emails.",
    "column_default": null,
    "column_name": "receives_admin_info_emails",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Info Emails.",
    "column_default": null,
    "column_name": "receives_info_emails",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Sender Address.",
    "column_default": null,
    "column_name": "sender_email",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Sender Name.",
    "column_default": null,
    "column_name": "sender_name",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Email Signature.",
    "column_default": null,
    "column_name": "signature",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Url for IOS banner photo.",
    "column_default": null,
    "column_name": "small_banner_photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Photo.",
    "column_default": null,
    "column_name": "small_photo_url",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Stay-in-Touch Email Note.",
    "column_default": null,
    "column_name": "stay_in_touch_note",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Stay-in-Touch Email Signature.",
    "column_default": null,
    "column_name": "stay_in_touch_signature",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Stay-in-Touch Email Subject.",
    "column_default": null,
    "column_name": "stay_in_touch_subject",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The System Modstamp.",
    "column_default": null,
    "column_name": "system_modstamp",
    "data_type": "timestamp with time zone",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Time Zone.",
    "column_default": null,
    "column_name": "time_zone_sid_key",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Title.",
    "column_default": null,
    "column_name": "title",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Auto-login To Call Center.",
    "column_default": null,
    "column_name": "user_permissions_call_center_auto_login",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Flow User.",
    "column_default": null,
    "column_name": "user_permissions_interaction_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Data.com User.",
    "column_default": null,
    "column_name": "user_permissions_jigsaw_prospecting_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Knowledge User.",
    "column_default": null,
    "column_name": "user_permissions_knowledge_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Chat User.",
    "column_default": null,
    "column_name": "user_permissions_live_agent_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Marketing User.",
    "column_default": null,
    "column_name": "user_permissions_marketing_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Apex Mobile User.",
    "column_default": null,
    "column_name": "user_permissions_mobile_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Offline User.",
    "column_default": null,
    "column_name": "user_permissions_offline_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Salesforce CRM Content User.",
    "column_default": null,
    "column_name": "user_permissions_sf_content_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Site.com Contributor User.",
    "column_default": null,
    "column_name": "user_permissions_siteforce_contributor_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Site.com Publisher User.",
    "column_default": null,
    "column_name": "user_permissions_siteforce_publisher_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Service Cloud User.",
    "column_default": null,
    "column_name": "user_permissions_support_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The WDC User.",
    "column_default": null,
    "column_name": "user_permissions_work_dot_com_user_feature",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ActivityRemindersPopup.",
    "column_default": null,
    "column_name": "user_preferences_activity_reminders_popup",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ApexPagesDeveloperMode.",
    "column_default": null,
    "column_name": "user_preferences_apex_pages_developer_mode",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The CacheDiagnostics.",
    "column_default": null,
    "column_name": "user_preferences_cache_diagnostics",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ContentEmailAsAndWhen.",
    "column_default": null,
    "column_name": "user_preferences_content_email_as_and_when",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ContentNoEmail.",
    "column_default": null,
    "column_name": "user_preferences_content_no_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The CreateLEXAppsWTShown.",
    "column_default": null,
    "column_name": "user_preferences_create_lex_apps_wt_shown",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisCommentAfterLikeEmail.",
    "column_default": null,
    "column_name": "user_preferences_dis_comment_after_like_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisMentionsCommentEmail.",
    "column_default": null,
    "column_name": "user_preferences_dis_mentions_comment_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisProfPostCommentEmail.",
    "column_default": null,
    "column_name": "user_preferences_dis_prof_post_comment_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableAllFeedsEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_all_feeds_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableBookmarkEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_bookmark_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableChangeCommentEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_change_comment_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableEndorsementEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_endorsement_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableFeedbackEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_feedback_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableFileShareNotificationsForApi.",
    "column_default": null,
    "column_name": "user_preferences_disable_file_share_notifications_for_api",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableFollowersEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_followers_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableLaterCommentEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_later_comment_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableLikeEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_like_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableMentionsPostEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_mentions_post_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableMessageEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_message_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableProfilePostEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_profile_post_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableSharePostEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_share_post_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The DisableWorkEmail.",
    "column_default": null,
    "column_name": "user_preferences_disable_work_email",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The EnableAutoSubForFeeds.",
    "column_default": null,
    "column_name": "user_preferences_enable_auto_sub_for_feeds",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The EventRemindersCheckboxDefault.",
    "column_default": null,
    "column_name": "user_preferences_event_reminders_checkbox_default",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ExcludeMailAppAttachments.",
    "column_default": null,
    "column_name": "user_preferences_exclude_mail_app_attachments",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The FavoritesShowTopFavorites.",
    "column_default": null,
    "column_name": "user_preferences_favorites_show_top_favorites",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The FavoritesWTShown.",
    "column_default": null,
    "column_name": "user_preferences_favorites_wt_shown",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The GlobalNavBarWTShown.",
    "column_default": null,
    "column_name": "user_preferences_global_nav_bar_wt_shown",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The GlobalNavGridMenuWTShown.",
    "column_default": null,
    "column_name": "user_preferences_global_nav_grid_menu_wt_shown",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HasCelebrationBadge.",
    "column_default": null,
    "column_name": "user_preferences_has_celebration_badge",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideBiggerPhotoCallout.",
    "column_default": null,
    "column_name": "user_preferences_hide_bigger_photo_callout",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideChatterOnboardingSplash.",
    "column_default": null,
    "column_name": "user_preferences_hide_chatter_onboarding_splash",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideCSNDesktopTask.",
    "column_default": null,
    "column_name": "user_preferences_hide_csn_desktop_task",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideCSNGetChatterMobileTask.",
    "column_default": null,
    "column_name": "user_preferences_hide_csn_get_chatter_mobile_task",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideEndUserOnboardingAssistantModal.",
    "column_default": null,
    "column_name": "user_preferences_hide_end_user_onboarding_assistant_modal",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideLightningMigrationModal.",
    "column_default": null,
    "column_name": "user_preferences_hide_lightning_migration_modal",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideS1BrowserUI.",
    "column_default": null,
    "column_name": "user_preferences_hide_s_1_browser_ui",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideSecondChatterOnboardingSplash.",
    "column_default": null,
    "column_name": "user_preferences_hide_second_chatter_onboarding_splash",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The HideSfxWelcomeMat.",
    "column_default": null,
    "column_name": "user_preferences_hide_sfx_welcome_mat",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The JigsawListUser.",
    "column_default": null,
    "column_name": "user_preferences_jigsaw_list_user",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The LightningExperiencePreferred.",
    "column_default": null,
    "column_name": "user_preferences_lightning_experience_preferred",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The NewLightningReportRunPageEnabled.",
    "column_default": null,
    "column_name": "user_preferences_new_lightning_report_run_page_enabled",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The PathAssistantCollapsed.",
    "column_default": null,
    "column_name": "user_preferences_path_assistant_collapsed",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The PreviewCustomTheme.",
    "column_default": null,
    "column_name": "user_preferences_preview_custom_theme",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The PreviewLightning.",
    "column_default": null,
    "column_name": "user_preferences_preview_lightning",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The RecordHomeReservedWTShown.",
    "column_default": null,
    "column_name": "user_preferences_record_home_reserved_wt_shown",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The RecordHomeSectionCollapseWTShown.",
    "column_default": null,
    "column_name": "user_preferences_record_home_section_collapse_wt_shown",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ReminderSoundOff.",
    "column_default": null,
    "column_name": "user_preferences_reminder_sound_off",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowCityToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_city_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowCityToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_city_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowCountryToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_country_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowCountryToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_country_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowEmailToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_email_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowEmailToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_email_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowFaxToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_fax_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowFaxToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_fax_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowManagerToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_manager_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowManagerToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_manager_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowMobilePhoneToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_mobile_phone_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowMobilePhoneToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_mobile_phone_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowPostalCodeToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_postal_code_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowPostalCodeToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_postal_code_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowProfilePicToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_profile_pic_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowStateToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_state_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowStateToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_state_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowStreetAddressToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_street_address_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowStreetAddressToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_street_address_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowTitleToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_title_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowTitleToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_title_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowWorkPhoneToExternalUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_work_phone_to_external_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The ShowWorkPhoneToGuestUsers.",
    "column_default": null,
    "column_name": "user_preferences_show_work_phone_to_guest_users",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SortFeedByComment.",
    "column_default": null,
    "column_name": "user_preferences_sort_feed_by_comment",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SuppressEventSFXReminders.",
    "column_default": null,
    "column_name": "user_preferences_suppress_event_sfx_reminders",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The SuppressTaskSFXReminders.",
    "column_default": null,
    "column_name": "user_preferences_suppress_task_sfx_reminders",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The TaskRemindersCheckboxDefault.",
    "column_default": null,
    "column_name": "user_preferences_task_reminders_checkbox_default",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The UserDebugModePref.",
    "column_default": null,
    "column_name": "user_preferences_user_debug_mode_pref",
    "data_type": "boolean",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The Role ID.",
    "column_default": null,
    "column_name": "user_role_id",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "The category of user license. Can be one of Standard, PowerPartner, CSPLitePortal, CustomerSuccess, PowerCustomerSuccess, CsnOnly, and Guest.",
    "column_default": null,
    "column_name": "user_type",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  },
  {
    "column_comment": "Login name of the user.",
    "column_default": null,
    "column_name": "username",
    "data_type": "text",
    "is_nullable": "YES",
    "table_comment": "Represents a user in organization.",
    "table_name": "salesforce_user",
    "table_schema": "salesforce"
  }
]
```
