# Table: salesforce_account

Represents an individual account, which is an organization or person involved with business (such as customers, competitors, and partners).

## Examples

### Basic info

```sql
select
  name,
  industry,
  type,
  ownership,
  rating
from
  salesforce_account;
```

### List Active users

```sql
select
  username,
  alias,
  user_type,
  is_active,
  last_login_date
from
  salesforce_user
where
  is_active;
```

### List Standard users

```sql
select
  username,
  alias,
  user_type,
  is_active,
  last_login_date
from
  salesforce_user
where
  user_type = 'Standard';
```

```json
{
  "AccountNumber": "CD355118",
  "AccountSource": null,
  "Active__c": "Yes",
  "AnnualRevenue": 5600000000,
  "CleanStatus": "Pending",
  "CreatedById": "0055j000003GU6IAAW",
  "CreatedDate": "2021-12-23T07:50:29.000+0000",
  "CustomerPriority__c": "High",
  "DandbCompanyId": null,
  "Description": "World's third largest oil and gas company.",
  "DunsNumber": null,
  "Fax": "(212) 842-5501",
  "Id": "0015j00000Erg8HAAR",
  "Industry": "Energy",
  "IsDeleted": false,
  "Jigsaw": null,
  "JigsawCompanyId": null,
  "LastActivityDate": null,
  "LastModifiedById": "0055j000003GU6IAAW",
  "LastModifiedDate": "2021-12-23T07:50:29.000+0000",
  "LastReferencedDate": "2022-01-10T13:26:18.000+0000",
  "LastViewedDate": "2022-01-10T13:26:18.000+0000",
  "MasterRecordId": null,
  "NaicsCode": null,
  "NaicsDesc": null,
  "Name": "United Oil \u0026 Gas Corp.",
  "NumberOfEmployees": 145000,
  "NumberofLocations__c": 955,
  "OperatingHoursId": null,
  "OwnerId": "0055j000003GU6IAAW",
  "Ownership": "Public",
  "ParentId": null,
  "Phone": "(212) 842-5500",
  "PhotoUrl": "/services/images/photo/0015j00000Erg8HAAR",
  "Rating": "Hot",
  "SLAExpirationDate__c": "2022-07-20",
  "SLASerialNumber__c": "6654",
  "SLA__c": "Platinum",
  "Sic": "4437",
  "SicDesc": null,
  "Site": null,
  "SystemModstamp": "2021-12-23T07:50:29.000+0000",
  "TickerSymbol": "UOS",
  "Tradestyle": null,
  "Type": "Customer - Direct",
  "UpsellOpportunity__c": "Yes",
  "Website": "http://www.uos.com",
  "YearStarted": null,
  "__client__": {},
  "attributes": {
    "type": "Account",
    "url": "/services/data/v43.0/sobjects/Account/0015j00000Erg8HAAR"
  }
}
```
