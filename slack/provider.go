package slack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SLACK_TOKEN", nil),
				Description: "The API key to use for interacting with your Slack team.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"slack_channel": resourceChannel(),
			"slack_user":    resourceUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Token: d.Get("api_token").(string),
	}
	return config, nil
}
