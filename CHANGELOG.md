## v0.6.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#30](https://github.com/turbot/steampipe-plugin-salesforce/pull/30))
- Recompiled plugin with Go version `1.21`. ([#30](https://github.com/turbot/steampipe-plugin-salesforce/pull/30))

## v0.5.0 [2023-08-24]

_What's new?_

- Added a new config argument `naming_convention` to support `api_native` or `snake_case` table and column names. Please check the [Configuration](https://hub.steampipe.io/plugins/turbot/salesforce#configuration) section for more information on how to use the new config argument. ([#24](https://github.com/turbot/steampipe-plugin-salesforce/pull/24))

## v0.4.0 [2023-06-20]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.5.0/CHANGELOG.md#v550-2023-06-16) which significantly reduces API calls and boosts query performance, resulting in faster data retrieval. This update significantly lowers the plugin initialization time of dynamic plugins by avoiding recursing into child folders when not necessary. ([#20](https://github.com/turbot/steampipe-plugin-salesforce/pull/20))

## v0.3.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#12](https://github.com/turbot/steampipe-plugin-salesforce/pull/12))
- Recompiled plugin with Go version `1.19`. ([#12](https://github.com/turbot/steampipe-plugin-salesforce/pull/12))

## v0.2.1 [2023-03-02]

_Bug fixes_

- Fixed formatting for various examples in index and table docs.

## v0.2.0 [2022-06-24]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v330--2022-6-22). ([#9](https://github.com/turbot/steampipe-plugin-salesforce/pull/9))

## v0.1.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#7](https://github.com/turbot/steampipe-plugin-salesforce/pull/7))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#6](https://github.com/turbot/steampipe-plugin-salesforce/pull/6))

## v0.0.1 [2022-02-15]

_What's new?_

- New tables added
  - [salesforce_account](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account)
  - [salesforce_account_contact_role](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_account_contact_role)
  - [salesforce_asset](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_asset)
  - [salesforce_contact](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_contact)
  - [salesforce_contract](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_contract)
  - [salesforce_lead](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_lead)
  - [salesforce_object_permission](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_object_permission)
  - [salesforce_opportunity](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_opportunity)
  - [salesforce_opportunity_contact_role](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_opportunity_contact_role)
  - [salesforce_order](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_order)
  - [salesforce_permission_set](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_permission_set)
  - [salesforce_permission_set_assignment](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_permission_set_assignment)
  - [salesforce_pricebook](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_pricebook)
  - [salesforce_product](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_product)
  - [salesforce_user](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_user)
  - [salesforce_{object_name}](https://hub.steampipe.io/plugins/turbot/salesforce/tables/salesforce_{object_name})
