package confluence

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccConfluenceContent_Updated(t *testing.T) {
	rName := acctest.RandomWithPrefix("resource-content-test")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConfluenceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConfluenceContentConfigRequired(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfluenceExists("confluence_content.default"),
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
					testAccCheckConfluenceExists("confluence_content.default"),
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

func TestAccConfluenceContent_Parent(t *testing.T) {
	rName := acctest.RandomWithPrefix("resource-content-test")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConfluenceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConfluenceContentConfigParent(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfluenceExists("confluence_content.parent"),
					testAccCheckConfluenceExists("confluence_content.child"),
					resource.TestCheckResourceAttrPair("confluence_content.child", "parent",
						"confluence_content.parent", "id"),
				),
			},
		},
	})
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

func testAccCheckConfluenceContentConfigParent(rName string) string {
	return fmt.Sprintf(`
	resource confluence_content "parent" {
		title = "%s-parent"
		body  = "parent"
	}
	resource confluence_content "child" {
		title  = "%s-child"
		body   = "child"
		parent = confluence_content.parent.id
	}
	`, rName, rName)
}
