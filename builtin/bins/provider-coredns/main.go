package main

import (
	"github.com/hashicorp/terraform/builtin/providers/coredns"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: coredns.Provider,
	})
}
