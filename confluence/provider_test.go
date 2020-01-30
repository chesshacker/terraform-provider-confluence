package confluence

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"confluence": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CONFLUENCE_SITE"); v == "" {
		t.Fatal("CONFLUENCE_SITE must be set for acceptance tests")
	}
	if v := os.Getenv("CONFLUENCE_USER"); v == "" {
		t.Fatal("CONFLUENCE_USER must be set for acceptance tests")
	}
	if v := os.Getenv("CONFLUENCE_TOKEN"); v == "" {
		t.Fatal("CONFLUENCE_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("CONFLUENCE_SPACE"); v == "" {
		t.Fatal("CONFLUENCE_SPACE must be set for acceptance tests")
	}
}
func testAccCheckConfluenceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)
	return confluenceDestroyHelper(s, client)
}

func testAccCheckConfluenceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Client)
		return confluenceExistsHelper(s, client)
	}
}

func confluenceDestroyHelper(s *terraform.State, client *Client) error {
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

func confluenceExistsHelper(s *terraform.State, client *Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		switch r.Type {
		case "confluence_content":
			_, err := client.GetContent(id)
			if err != nil {
				return fmt.Errorf("Received an error retrieving content %s", err)
			}
		case "confluence_attachment":
			_, err := client.GetAttachment(id)
			if err != nil {
				return fmt.Errorf("Received an error retrieving attachment %s", err)
			}
		default:
			return fmt.Errorf("Unknown resource: type = %s, id = %s", r.Type, id)
		}
	}
	return nil
}
