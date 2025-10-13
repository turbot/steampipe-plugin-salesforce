## v1.2.0 [2025-10-13]

_Dependencies_

- Recompiled plugin with Go version `1.24`. ([#81](https://github.com/turbot/steampipe-plugin-salesforce/pull/81))
- Recompiled plugin with [steampipe-plugin-sdk v5.13.1](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5131-2025-09-25) that addresses critical and high vulnerabilities in dependent packages. ([#85](https://github.com/turbot/steampipe-plugin-salesforce/pull/85))

## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#72](https://github.com/turbot/steampipe-plugin-salesforce/pull/72))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#72](https://github.com/turbot/steampipe-plugin-salesforce/pull/72))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#61](https://github.com/turbot/steampipe-plugin-salesforce/pull/61))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#61](https://github.com/turbot/steampipe-plugin-salesforce/pull/61))

## v0.7.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#43](https://github.com/turbot/steampipe-plugin-salesforce/pull/43))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#43](https://github.com/turbot/steampipe-plugin-salesforce/pull/43))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-salesforce/blob/main/docs/LICENSE). ([#43](https://github.com/turbot/steampipe-plugin-salesforce/pull/43))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#42](https://github.com/turbot/steampipe-plugin-salesforce/pull/42))

## v0.6.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#33](https://github.com/turbot/steampipe-plugin-salesforce/pull/33))

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
