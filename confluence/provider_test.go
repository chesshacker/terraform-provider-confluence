package confluence

import (
	"os"
	"testing"

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
