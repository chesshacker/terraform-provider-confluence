package confluence

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccConfluenceAttachment_Created(t *testing.T) {
	if os.Getenv("RUN_ATTACHMENT_TEST") == "" {
		t.Skip("skipping test for now... set RUN_ATTACHMENT_TEST to enable test")
	} else {
		rName := acctest.RandomWithPrefix("resource-attachment-test")
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: testAccCheckConfluenceAttachmentDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccCheckConfluenceAttachmentConfigRequired(rName),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckConfluenceAttachmentExists("confluence_attachment.default"),
						resource.TestCheckResourceAttr(
							"confluence_attachment.default", "title", "file.txt"),
						resource.TestCheckResourceAttr(
							"confluence_attachment.default", "body", rName),
						resource.TestCheckResourceAttr(
							"confluence_attachment.default", "version", "1"),
					),
				},
			},
		})
	}
}

func testAccCheckConfluenceAttachmentDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)
	return confluenceAttachmentDestroyHelper(s, client)
}

func testAccCheckConfluenceAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Client)
		return confluenceAttachmentExistsHelper(s, client)
	}
}

func testAccCheckConfluenceAttachmentConfigRequired(rName string) string {
	return fmt.Sprintf(`
resource confluence_content "default" {
	title = "%s"
	body  = "Original value"
}
resource confluence_attachment "default" {
  title = "file.txt"
	data  = "%s"
	page = confluence_content.default.id
}
`, rName, rName)
}

func confluenceAttachmentDestroyHelper(s *terraform.State, client *Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		switch r.Type {
		case "confluence_content":
			_, err := client.GetContent(id)
			if err == nil {
				return fmt.Errorf("Content still exists. id: %s", id)
			}
		case "confluence_attachment":
			_, err := client.GetAttachment(id)
			if err == nil {
				return fmt.Errorf("Attachment still exists. id: %s", id)
			}
		default:
			return fmt.Errorf("Unknown resource: type = %s, id = %s", r.Type, id)
		}
	}
	return nil
}

func confluenceAttachmentExistsHelper(s *terraform.State, client *Client) error {
	for _, r := range s.RootModule().Resources {
		if r.Type == "confluence_attachment" {
			id := r.Primary.ID
			_, err := client.GetAttachment(id)
			if err != nil {
				return fmt.Errorf("Received an error retrieving attachment %s", err)
			}
		}
	}
	return nil
}
