# Salesforce Plugin for Steampipe - A streamlined way to query Salesforce via ODBC via this Steampipe plugin. 

This fork is a (working) experiment and thus should not be relied on for production environments. 
Summary of changes:
- disables conversion of field API names to snake_case  (eg. CreatedDate to created_date)
- lower-cases all field API names for simplicity (and convenience) when writing sql queries
- reduced the 'static tables' intitialized down to a single mostly unused object in Salesforce 'Pricebook2' (ignores all the salesforce_ schemas that come out of the box with this plugin) thus allowing for dynamically mapped object schemas

Forked from https://github.com/turbot/steampipe-plugin-salesforce on March 22, 2023.

Built for linux with libraries embeded using...

'''shell
go build -o -arch=amd64 -os=linux ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/salesforce@latest/steampipe-plugin-salesforce.plugin *.go
'''

---8<--- remainder of original Readme.md is below this line ----

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
  name,
  amount,
  close_date
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
