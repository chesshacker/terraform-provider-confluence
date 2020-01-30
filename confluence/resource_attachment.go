package confluence

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAttachmentCreate,
		Read:   resourceAttachmentRead,
		Update: resourceAttachmentUpdate,
		Delete: resourceAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"media_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "text/plain",
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	attachmentRequest := attachmentFromResourceData(d)
	pageId := d.Get("page").(string)
	data := d.Get("data").(string)
	attachmentResponse, err := client.CreateAttachment(attachmentRequest, data, pageId)
	if err != nil {
		return err
	}
	d.SetId(attachmentResponse.Id)
	return resourceAttachmentRead(d, m)
}

func resourceAttachmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	attachmentResponse, err := client.GetAttachment(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}
	attachmentData, err := client.GetAttachmentBody(attachmentResponse)
	if err != nil {
		d.SetId("")
		return err
	}
	err = d.Set("data", attachmentData)
	if err != nil {
		return err
	}
	return updateResourceDataFromAttachment(d, attachmentResponse, client)
}

func resourceAttachmentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	attachmentRequest := attachmentFromResourceData(d)
	pageId := d.Get("page").(string)
	data := d.Get("data").(string)
	_, err := client.UpdateAttachment(attachmentRequest, data, pageId)
	if err != nil {
		d.SetId("")
		return err
	}
	return resourceAttachmentRead(d, m)
}

func resourceAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	pageId := d.Get("page").(string)
	err := client.DeleteAttachment(d.Id(), pageId)
	if err != nil {
		return err
	}
	// d.SetId("") is automatically called assuming delete returns no errors
	return nil
}

func attachmentFromResourceData(d *schema.ResourceData) *Attachment {
	result := &Attachment{
		Id:   d.Id(),
		Type: "attachment",
		Metadata: &Metadata{
			MediaType: d.Get("media_type").(string),
		},
		Title: d.Get("title").(string),
	}
	version := d.Get("version").(int) // Get returns 0 if unset
	if version > 0 {
		result.Version = &Version{Number: version}
	}
	return result
}

func updateResourceDataFromAttachment(d *schema.ResourceData, attachment *Attachment, client *Client) error {
	d.SetId(attachment.Id)
	m := map[string]interface{}{
		"title":      attachment.Title,
		"version":    attachment.Version.Number,
		"media_type": attachment.Metadata.MediaType,
	}
	for k, v := range m {
		err := d.Set(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
