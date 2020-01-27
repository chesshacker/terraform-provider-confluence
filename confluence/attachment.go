package confluence

import (
	"fmt"
)

// Attachment is a primary resource in Confluence
type Attachment struct {
	Id       string           `json:"id,omitempty"`
	Metadata *Metadata        `json:"metadata,omitempty"`
	Title    string           `json:"title,omitempty"` // filename
	Type     string           `json:"type,omitempty"`  // always "attachment"
	Version  *Version         `json:"version,omitempty"`
	Links    *AttachmentLinks `json:"_links,omitempty"`
}

// Metadata is part of an Attachment
type Metadata struct {
	MediaType string `json:"mediaType,omitempty"`
}

// AttachmentLinks is part of Content
type AttachmentLinks struct {
	Context  string `json:"context,omitempty"`  // "/wiki"
	Download string `json:"download,omitempty"` // prefix with Context
}

func (c *Client) CreateAttachment(attachment *Attachment, data, pageId string) (*Attachment, error) {
	var response Attachment
	path := fmt.Sprintf("/wiki/rest/api/content/%s/child/attachment", pageId)
	if err := c.PostForm(path, attachment.Title, data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetAttachment(id string) (*Attachment, error) {
	var response Attachment
	path := fmt.Sprintf("/wiki/rest/api/content/%s?expand=version", id)
	if err := c.Get(path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) UpdateAttachment(attachment *Attachment, data, pageId string) (*Attachment, error) {
	var response Attachment
	attachment.Version.Number++
	path := fmt.Sprintf("/wiki/rest/api/content/%s/child/attachment/%s", pageId, attachment.Id)
	if err := c.PutForm(path, attachment.Title, data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteAttachment(id, pageId string) error {
	path := fmt.Sprintf("/wiki/rest/api/content/%s/child/attachment/%s", pageId, id)
	if err := c.Delete(path); err != nil {
		return err
	}
	return nil
}
