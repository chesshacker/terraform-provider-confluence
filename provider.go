package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"instance": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cloud Confluence Instance Name (the name before atlassian.net)",
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User's email address",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confluence API Token for user",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"confluence_content": resourceContent(),
		},
	}
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		meta, err := NewClient(&NewClientInput{
			instance: d.Get("instance").(string),
			token:    d.Get("token").(string),
			user:     d.Get("user").(string),
		})
		if err != nil {
			return nil, err
		}
		return meta, nil
	}
}
