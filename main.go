package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	fireblocksprovider "github.com/liquid-collective/terraform-provider-fireblocks/src/provider"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fireblocksprovider.Provider,
	})
}
