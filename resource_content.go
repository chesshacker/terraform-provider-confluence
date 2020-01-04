package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceContent() *schema.Resource {
	return &schema.Resource{
		Create: resourceContentCreate,
		Read:   resourceContentRead,
		Update: resourceContentUpdate,
		Delete: resourceContentDelete,

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "page",
			},
			"space": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"body": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceContentCreate(d *schema.ResourceData, m interface{}) error {
	title := d.Get("title").(string)
	d.SetId(title)
	return resourceContentRead(d, m)
}

func resourceContentRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceContentUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceContentRead(d, m)
}

func resourceContentDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
