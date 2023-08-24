# Table: salesforce_pricebook

Represents a price book that contains the list of products that your org sells.

If the naming_convention parameter is set to api_native in the config file, then the table and column names will match what’s in Salesforce. For instance, the query `select name, is_active from salesforce_pricebook` would become `select "Name", "IsActive" from "Pricebook2"`.

If the `naming_convention` configuration argument is set to `api_native`, please see [API Native Examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_pricebook#api_native_examples).

## Examples

### Basic info

```sql
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook;
```

### List standard price books

```sql
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook
where
  is_standard;
```

### List active price books

```sql
select
  name,
  is_active,
  is_standard,
  created_by_id,
  created_date,
  description
from
  salesforce_pricebook
where
  is_active;
```

## API Native Examples

If the `naming_convention` config argument is set to `api_native`, the table and column names will match Salesforce naming conventions.

### Basic info (with API Native naming convention)

```sql
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2";
```

### List standard price books (with API Native naming convention)

```sql
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsStandard";
```

## Show archived price books

```sql
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsArchived";
```

### Show deleted price books

```sql
select
  "Name",
  "IsActive",
  "IsStandard",
  "CreatedById",
  "CreatedDate",
  "Description"
from
  "Pricebook2"
where
  "IsDeleted";
```
