package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Body struct {
	Storage *Storage `json:"storage,omitempty"`
}

type Content struct {
	Id      string        `json:"id,omitempty"`
	Type    string        `json:"type,omitempty"`
	Title   string        `json:"title,omitempty"`
	Space   *Space        `json:"space,omitempty"`
	Version *Version      `json:"version,omitempty"`
	Body    *Body         `json:"body,omitempty"`
	Links   *ContentLinks `json:"_links,omitempty"`
}

type ContentLinks struct {
	Context string `json:"context,omitempty"`
	WebUI   string `json:"webui,omitempty"`
}

type Space struct {
	Key string `json:"key,omitempty"`
}

type Storage struct {
	Value          string `json:"value,omitempty"`
	Representation string `json:"representation,omitempty"`
}

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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "page",
			},
			"space": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"body": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceContentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	contentRequest := Content{
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

	body, err := json.Marshal(contentRequest)
	if err != nil {
		return err
	}

	resp, err := client.Post("/wiki/rest/api/content", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP POST request error. Response code: %d", resp.StatusCode)
	}

	var contentResponse Content
	err = json.NewDecoder(resp.Body).Decode(&contentResponse)
	if err != nil {
		return err
	}

	// TODO: I feel like a lot of this response code checking and marshalling of
	// the response could be hoisted up to the client, but I'm not exactly sure
	// how I want it to look yet, so I am leaving it as-is for now.

	d.SetId(contentResponse.Id)
	return resourceContentRead(d, m)
}

func resourceContentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	resp, err := client.Get("/wiki/rest/api/content/" + d.Id() + "?expand=space,body.storage,version")
	if err != nil {
		d.SetId("")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		d.SetId("")
		return fmt.Errorf("HTTP GET request error. Response code: %d", resp.StatusCode)
	}

	var contentResponse Content
	err = json.NewDecoder(resp.Body).Decode(&contentResponse)
	if err != nil {
		d.SetId("")
		return err
	}
	d.SetId(contentResponse.Id)
	d.Set("type", contentResponse.Type)
	d.Set("space", contentResponse.Space.Key)
	d.Set("body", contentResponse.Body.Storage.Value)
	d.Set("title", contentResponse.Title)
	d.Set("version", contentResponse.Version.Number)
	d.Set("url", client.URL(contentResponse.Links.Context+contentResponse.Links.WebUI))
	return nil
}

func resourceContentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	contentRequest := Content{
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
		Version: &Version{
			Number: d.Get("version").(int) + 1,
		},
	}

	body, err := json.Marshal(contentRequest)
	if err != nil {
		return err
	}

	resp, err := client.Put("/wiki/rest/api/content/"+d.Id(), bytes.NewReader(body))
	if err != nil {
		d.SetId("")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP PUT request error. Response code: %d", resp.StatusCode)
	}

	return resourceContentRead(d, m)
}

func resourceContentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	resp, err := client.Delete("/wiki/rest/api/content/" + d.Id())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("HTTP DELETE request error. Response code: %d", resp.StatusCode)
	}
	// d.SetId("") is automatically called assuming delete returns no errors
	return nil
}
