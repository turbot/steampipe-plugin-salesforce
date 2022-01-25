# Table: salesforce_contact

Represents a contact, which is a person associated with an account.

## Examples

### Basic info

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
