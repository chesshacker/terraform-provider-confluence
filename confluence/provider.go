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
				Description: "Confluence hostname (<name>.atlassian.net if using Cloud Confluence, otherwise hostname)",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_SITE", nil),
			},
			"site_scheme": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Optional https or http scheme to use for API calls",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_SITE_SCHEME", "https"),
			},
			"public_site": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Optional public Confluence Server hostname if different than API hostname",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_PUBLIC_SITE", ""),
			},
			"public_site_scheme": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Optional https or http scheme to use for public site URLs",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_PUBLIC_SITE_SCHEME", "https"),
			},
			"context": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confluence path context (Will default to /wiki if using an atlassian.net hostname)",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_CONTEXT", ""),
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User's email address for Cloud Confluence or username for Confluence Server",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_USER", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confluence API Token for Cloud Confluence or password for Confluence Server",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_TOKEN", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"confluence_content":    resourceContent(),
			"confluence_attachment": resourceAttachment(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return NewClient(&NewClientInput{
		site:             d.Get("site").(string),
		siteScheme:       d.Get("site_scheme").(string),
		publicSite:       d.Get("public_site").(string),
		publicSiteScheme: d.Get("public_site_scheme").(string),
		context:          d.Get("context").(string),
		token:            d.Get("token").(string),
		user:             d.Get("user").(string),
	}), nil
}
