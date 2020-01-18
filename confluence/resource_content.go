package confluence

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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
	contentResponse, err := client.CreateContent(contentRequest)
	if err != nil {
		return err
	}
	d.SetId(contentResponse.Id)
	return resourceContentRead(d, m)
}

func resourceContentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	contentResponse, err := client.GetContent(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}
	return updateResourceDataFromContent(d, contentResponse, client)
}

func resourceContentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	contentRequest := contentFromResourceData(d)
	_, err := client.UpdateContent(contentRequest)
	if err != nil {
		d.SetId("")
		return err
	}
	return resourceContentRead(d, m)
}

func resourceContentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	err := client.DeleteContent(d.Id())
	if err != nil {
		return err
	}
	// d.SetId("") is automatically called assuming delete returns no errors
	return nil
}

func contentFromResourceData(d *schema.ResourceData) *Content {
	result := &Content{
		Id:   d.Id(),
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
