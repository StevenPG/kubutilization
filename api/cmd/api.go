package main

import (
	"flag"
	"github.com/stevenpg/kubutilization/api/pkg/client"
)

// parseFlagUseExternalConnection - Initializes are parses flags, returning the ptrs to caller
func parseFlagUseExternalConnection() *bool {
	externalConnection := flag.Bool("external", false, "a boolean selector for whether to search for a kubeconfig in ~/.kube/config")
	flag.Parse()
	return externalConnection
}

func main() {
	if *parseFlagUseExternalConnection() {
		client.ExternalConnection()
	} else {
		client.Connection()
	}
}
