package salesforce

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type salesforceConfig struct {
	URL        *string   `cty:"url"`
	Username   *string   `cty:"username"`
	Password   *string   `cty:"password"`
	Token      *string   `cty:"token"`
	ClientId   *string   `cty:"client_id"`
	APIVersion *string   `cty:"api_version"`
	Objects    *[]string `cty:"objects"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
	"client_id": {
		Type: schema.TypeString,
	},
	"api_version": {
		Type: schema.TypeString,
	},
	"objects": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{
			Type: schema.TypeString,
		},
	},
}

func ConfigInstance() interface{} {
	return &salesforceConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) salesforceConfig {
	if connection == nil || connection.Config == nil {
		return salesforceConfig{}
	}
	config, _ := connection.Config.(salesforceConfig)
	return config
}
