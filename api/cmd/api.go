package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/stevenpg/kubutilization/api/pkg/client"
	"github.com/stevenpg/kubutilization/api/pkg/endpoint"
	"github.com/stevenpg/kubutilization/api/pkg/middleware"
)

// parseFlagUseExternalConnection ... Initializes are parses flags, returning the ptrs to caller
func parseFlagUseExternalConnection() *bool {
	externalConnection := flag.Bool("external", false, "a boolean selector for whether to search for a kubeconfig in ~/.kube/config")
	flag.Parse()
	return externalConnection
}

// registerEndpoints ... Add configured endpoints to Gin Engine
func registerEndpoints(engine *gin.Engine) {
	endpoint.RegisterUtilityEndpoints(engine)
	endpoint.RegisterNodeEndpoints(engine)
	endpoint.RegisterPodEndpoints(engine)
}

// main ... application entry point
func main() {
	if *parseFlagUseExternalConnection() {
		ginRouter := middleware.GinRouter(client.ExternalConnection())
		registerEndpoints(ginRouter)
		ginRouter.Run(":8080")
	} else {
		client.Connection()
	}
}
