package confluence

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
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
	}
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		meta, err := NewClient(&NewClientInput{
			site:  d.Get("site").(string),
			token: d.Get("token").(string),
			user:  d.Get("user").(string),
		})
		if err != nil {
			return nil, err
		}
		return meta, nil
	}
}
