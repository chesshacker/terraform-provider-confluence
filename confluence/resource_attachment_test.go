package confluence

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccConfluenceAttachment_Created(t *testing.T) {
	rName := acctest.RandomWithPrefix("resource-attachment-test")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConfluenceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConfluenceAttachmentConfigRequired(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfluenceExists("confluence_attachment.default"),
					resource.TestCheckResourceAttr(
						"confluence_attachment.default", "title", "file.txt"),
					resource.TestCheckResourceAttr(
						"confluence_attachment.default", "data", rName),
					resource.TestCheckResourceAttr(
						"confluence_attachment.default", "version", "1"),
				),
			},
		},
	})
}

func testAccCheckConfluenceAttachmentConfigRequired(rName string) string {
	time.Sleep(time.Second)
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
