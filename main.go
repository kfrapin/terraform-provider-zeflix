package main

import (
	// utilisation du SDKv2
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/kfrapin/terraform-provider-zeflix/zeflix"
)

var (
	version = "dev"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: zeflix.New(version),
	})
}
