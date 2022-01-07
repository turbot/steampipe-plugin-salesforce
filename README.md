![image](https://hub.steampipe.io/images/plugins/turbot/salesforce-social-graphic.png)

# Salesforce Plugin for Steampipe

Use SQL to query infrastructure accounts, users, oppurtinities and more from your Salesforce instance.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/salesforce)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/salesforce/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-salesforce/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install salesforce
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/salesforce#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/salesforce#configuration).

Run a query:

```sql
select
  id,
  name,
  amount,
  close_date,
  fiscal_quarter,
  fiscal_year
from
  salesforce_opportunity
where
  is_won;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-salesforce.git
cd steampipe-plugin-salesforce
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/salesforce.spc
```

Try it!

```shell
steampipe query
> .inspect salesforce
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-salesforce/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Salesforce Plugin](https://github.com/turbot/steampipe-plugin-salesforce/labels/help%20wanted)
