package confluence

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccConfluenceContent_Updated(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConfluenceContentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConfluenceContentConfigRequired,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfluenceContentExists("confluence_content.default"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "title", "Example Page"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "body", "<p>This page was built with Terraform<p>"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "version", "1"),
				),
			},
			{
				Config: testAccCheckConfluenceContentConfigUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfluenceContentExists("confluence_content.default"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "title", "Updated Page"),
					resource.TestCheckResourceAttr(
						"confluence_content.default", "body", "<p>This page was built with Terraform<p>"),
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

const testAccCheckConfluenceContentConfigRequired = `
resource confluence_content "default" {
  title = "Example Page"
  body  = "<p>This page was built with Terraform<p>"
}
`

const testAccCheckConfluenceContentConfigUpdated = `
resource confluence_content "default" {
  title = "Updated Page"
  body  = "<p>This page was built with Terraform<p>"
}
`

func confluenceContentDestroyHelper(s *terraform.State, client *Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		var contentResponse Content
		u := "/wiki/rest/api/content/" + id + "?expand=space,body.storage,version"
		err := client.Get(u, &contentResponse)
		if err == nil {
			return fmt.Errorf("Content still exists. id: %s", id)
		}
	}
	return nil
}

func confluenceContentExistsHelper(s *terraform.State, client *Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		var contentResponse Content
		u := "/wiki/rest/api/content/" + id + "?expand=space,body.storage,version"
		err := client.Get(u, &contentResponse)
		if err != nil {
			return fmt.Errorf("Received an error retrieving content %s", err)
		}
	}
	return nil
}
