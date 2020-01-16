package confluence

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

// Body is part of Content
type Body struct {
	Storage *Storage `json:"storage,omitempty"`
}

// Content is a primary resource in Confluence
type Content struct {
	Id      string        `json:"id,omitempty"`
	Type    string        `json:"type,omitempty"`
	Title   string        `json:"title,omitempty"`
	Space   *Space        `json:"space,omitempty"`
	Version *Version      `json:"version,omitempty"`
	Body    *Body         `json:"body,omitempty"`
	Links   *ContentLinks `json:"_links,omitempty"`
}

// ContentLinks is part of Content
type ContentLinks struct {
	Context string `json:"context,omitempty"`
	WebUI   string `json:"webui,omitempty"`
}

// Space is part of Content
type Space struct {
	Key string `json:"key,omitempty"`
}

// Storage is part of Body
type Storage struct {
	Value          string `json:"value,omitempty"`
	Representation string `json:"representation,omitempty"`
}

// Version is part of Content
type Version struct {
	Number int `json:"number,omitempty"`
}

func resourceContent() *schema.Resource {
	return &schema.Resource{
		Create: resourceContentCreate,
		Read:   resourceContentRead,
		Update: resourceContentUpdate,
		Delete: resourceContentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "page",
			},
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_SPACE", nil),
			},
			"body": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: resourceContentDiffBody,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceContentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	contentRequest := contentFromResourceData(d)
	var contentResponse Content
	err := client.Post("/wiki/rest/api/content", contentRequest, &contentResponse)
	if err != nil {
		return err
	}
	d.SetId(contentResponse.Id)
	return resourceContentRead(d, m)
}

func resourceContentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	var contentResponse Content
	u := "/wiki/rest/api/content/" + d.Id() + "?expand=space,body.storage,version"
	err := client.Get(u, &contentResponse) // TODO: contentResponse, err := client.GetContent(d.Id()) ?
	if err != nil {
		d.SetId("")
		return err
	}
	return updateResourceDataFromContent(d, &contentResponse, client)
}

func resourceContentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	contentRequest := contentFromResourceData(d)
	contentRequest.Version.Number++
	var contentResponse Content
	err := client.Put("/wiki/rest/api/content/"+d.Id(), contentRequest, &contentResponse)
	if err != nil {
		d.SetId("")
		return err
	}
	return resourceContentRead(d, m)
}

func resourceContentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	err := client.Delete("/wiki/rest/api/content/" + d.Id())
	if err != nil {
		return err
	}
	// d.SetId("") is automatically called assuming delete returns no errors
	return nil
}

func contentFromResourceData(d *schema.ResourceData) *Content {
	result := &Content{
		Type: d.Get("type").(string),
		Space: &Space{
			Key: d.Get("space").(string),
		},
		Body: &Body{
			Storage: &Storage{
				Value:          d.Get("body").(string),
				Representation: "storage",
			},
		},
		Title: d.Get("title").(string),
	}
	version := d.Get("version").(int) // Get returns 0 if unset
	if version > 0 {
		result.Version = &Version{Number: version}
	}
	return result
}

func updateResourceDataFromContent(d *schema.ResourceData, content *Content, client *Client) error {
	d.SetId(content.Id)
	m := map[string]interface{}{
		"type":    content.Type,
		"space":   content.Space.Key,
		"body":    content.Body.Storage.Value,
		"title":   content.Title,
		"version": content.Version.Number,
		"url":     client.URL(content.Links.Context + content.Links.WebUI),
	}
	for k, v := range m {
		err := d.Set(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Body was showing as requiring changes when there weren't any. It appears there
// are some whitespace differences between the old and new. This supresses the
// false differences by comparing the trimmed strings
func resourceContentDiffBody(k, old, new string, d *schema.ResourceData) bool {
	return strings.TrimSpace(old) == strings.TrimSpace(new)
}
