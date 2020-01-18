package confluence

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns the ResourceProvider for Confluence
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"site": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cloud Confluence Site Name (the name before atlassian.net)",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_SITE", nil),
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User's email address",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_USER", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confluence API Token for user",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_TOKEN", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"confluence_content": resourceContent(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return NewClient(&NewClientInput{
		site:  d.Get("site").(string),
		token: d.Get("token").(string),
		user:  d.Get("user").(string),
	}), nil
}
