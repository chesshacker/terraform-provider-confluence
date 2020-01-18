package main

import (
	"github.com/chesshacker/terraform-provider-confluence/confluence"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: confluence.Provider})
}
