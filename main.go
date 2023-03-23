package main

import (
	"github.com/turbot/steampipe-plugin-salesforce/salesforce"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: salesforce.Plugin})
}
