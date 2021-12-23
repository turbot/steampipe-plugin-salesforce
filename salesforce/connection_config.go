package salesforce

import (
	"context"
	"fmt"

	"github.com/simpleforce/simpleforce"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type salesforceConfig struct {
	URL      *string `cty:"url"`
	User     *string `cty:"user"`
	Password *string `cty:"password"`
	Token    *string `cty:"token"`
	ClientId *string `cty:"client_id"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"user": {
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

func connect(_ context.Context, d *plugin.QueryData) (*simpleforce.Client, error) {
	config := GetConfig(d.Connection)

	client := simpleforce.NewClient(*config.URL, simpleforce.DefaultClientID, simpleforce.DefaultAPIVersion)
	if client == nil {
		return nil, fmt.Errorf("couldn't get salesforce client. Clent generation error.")
	}
	err := client.LoginPassword(*config.User, *config.Password, *config.Token)
	if err != nil {
		return nil, fmt.Errorf("client LoginPassword Error")
	}

	return client, nil
}
