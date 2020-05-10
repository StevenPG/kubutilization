package main

import (
	"flag"
	"github.com/stevenpg/kubutilization/api/pkg/client"
	"github.com/stevenpg/kubutilization/api/pkg/endpoint"
	"github.com/stevenpg/kubutilization/api/pkg/middleware"
)

// parseFlagUseExternalConnection - Initializes are parses flags, returning the ptrs to caller
func parseFlagUseExternalConnection() *bool {
	externalConnection := flag.Bool("external", false, "a boolean selector for whether to search for a kubeconfig in ~/.kube/config")
	flag.Parse()
	return externalConnection
}

func main() {
	if *parseFlagUseExternalConnection() {
		server := middleware.GinEngine(client.ExternalConnection())
		endpoint.RegisterNodeEndpoints(server)
		endpoint.RegisterPodEndpoints(server)
		server.Run(":8080")
	} else {
		client.Connection()
	}
}
