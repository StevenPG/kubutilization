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
		httpEngine := middleware.GinEngine(client.ExternalConnection())
		endpoint.RegisterUtilityEndpoints(httpEngine)
		endpoint.RegisterNodeEndpoints(httpEngine)
		endpoint.RegisterPodEndpoints(httpEngine)
		httpEngine.Run(":8080")
	} else {
		client.Connection()
	}
}
