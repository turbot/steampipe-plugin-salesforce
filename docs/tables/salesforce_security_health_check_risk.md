---
title: "Steampipe Table: salesforce_security_health_check_risk - Query Salesforce Security Health Check using SQL"
description: "Allows users to query Salesforce Security Health Check risks, comparing each org security setting against the Salesforce baseline with a risk rating."
---

# Table: salesforce_security_health_check_risk - Query Salesforce Security Health Check using SQL

Salesforce Security Health Check evaluates the org's security settings — MFA enforcement, login IP ranges, session settings, password policies, certificates, and sharing — against a Salesforce-defined baseline. Each row is one setting, with its current org value, the baseline value, and a risk rating. This table is queried through the Salesforce Tooling API.

## Table Usage Guide

The `salesforce_security_health_check_risk` table gives a security administrator a read-only view of the org's security posture. Use it to find settings that deviate from the Salesforce baseline without needing the write-level Metadata API. It is the read-only path to session, password, MFA, and login-IP configuration.

**Important Notes**
- This table is backed by the Salesforce Tooling API and requires a user with the "View Health Check" permission.
- `risk_type` is one of `HIGH_RISK`, `MEDIUM_RISK`, or `MEETS_STANDARD`. `setting_risk_category` is one of `HIGH_RISK`, `MEDIUM_RISK`, `LOW_RISK`, or `INFORMATIONAL`.

## Examples

### Basic info
Review every security setting and how it compares to the Salesforce baseline.

```sql+postgres
select
  setting_group,
  setting,
  org_value,
  standard_value,
  risk_type
from
  salesforce_security_health_check_risk;
```

```sql+sqlite
select
  setting_group,
  setting,
  org_value,
  standard_value,
  risk_type
from
  salesforce_security_health_check_risk;
```

### High-risk settings that fall short of the baseline
Identify the settings that most weaken the org's security posture.

```sql+postgres
select
  setting_group,
  setting,
  org_value,
  standard_value
from
  salesforce_security_health_check_risk
where
  risk_type = 'HIGH_RISK';
```

```sql+sqlite
select
  setting_group,
  setting,
  org_value,
  standard_value
from
  salesforce_security_health_check_risk
where
  risk_type = 'HIGH_RISK';
```

### Password and MFA posture
Focus on the authentication-related settings.

```sql+postgres
select
  setting,
  org_value,
  standard_value,
  risk_type
from
  salesforce_security_health_check_risk
where
  setting_group in ('Identity', 'PasswordPolicies');
```

```sql+sqlite
select
  setting,
  org_value,
  standard_value,
  risk_type
from
  salesforce_security_health_check_risk
where
  setting_group in ('Identity', 'PasswordPolicies');
```
