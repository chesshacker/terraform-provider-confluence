package confluence

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccConfluenceContent_Updated(t *testing.T) {
	rName := acctest.RandomWithPrefix("resource_content_test_")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConfluenceContentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConfluenceContentConfigRequired(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfluenceContentExists("confluence_content.default"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "title", rName),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "body", "Original value"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "version", "1"),
				),
			},
			{
				Config: testAccCheckConfluenceContentConfigUpdated(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfluenceContentExists("confluence_content.default"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "title", rName),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "body", "Updated value"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "version", "2"),
				),
			},
		},
	})
}

func testAccCheckConfluenceContentDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)
	return confluenceContentDestroyHelper(s, client)
}

func testAccCheckConfluenceContentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Client)
		return confluenceContentExistsHelper(s, client)
	}
}

func testAccCheckConfluenceContentConfigRequired(rName string) string {
	return fmt.Sprintf(`
resource confluence_content "default" {
  title = "%s"
  body  = "Original value"
}
`, rName)
}

func testAccCheckConfluenceContentConfigUpdated(rName string) string {
	return fmt.Sprintf(`
	resource confluence_content "default" {
		title = "%s"
		body  = "Updated value"
	}
	`, rName)
}

func confluenceContentDestroyHelper(s *terraform.State, client *Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		_, err := client.GetContent(id)
		if err == nil {
			return fmt.Errorf("Content still exists. id: %s", id)
		}
	}
	return nil
}

func confluenceContentExistsHelper(s *terraform.State, client *Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		_, err := client.GetContent(id)
		if err != nil {
			return fmt.Errorf("Received an error retrieving content %s", err)
		}
	}
	return nil
}
