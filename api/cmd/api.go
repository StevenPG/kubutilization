package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/stevenpg/kubutilization/api/pkg/client"
	"github.com/stevenpg/kubutilization/api/pkg/endpoint"
	"github.com/stevenpg/kubutilization/api/pkg/middleware"
	"github.com/stevenpg/kubutilization/api/pkg/timeseries/writes"
)

// parseFlagUseExternalConnection ... Initializes are parses flags, returning the ptrs to caller
func parseInputFlags() (*bool, *bool) {
	externalConnection := flag.Bool("external", false, "a boolean selector for whether to search for a kubeconfig in ~/.kube/config")
	enableRedis := flag.Bool("timeseries", false, "a boolean selector for whether to connect to redis and write time-series data")
	flag.Parse()
	return externalConnection, enableRedis
}

// registerEndpoints ... Add configured endpoints to Gin Engine
func registerEndpoints(engine *gin.Engine) {
	endpoint.RegisterUtilityEndpoints(engine)
	endpoint.RegisterMetricsEndpoints(engine)
	endpoint.RegisterNamespaceEndpoints(engine)
	endpoint.RegisterNodeEndpoints(engine)
	endpoint.RegisterPodEndpoints(engine)
}

// main ... application entry point
func main() {
	useExternal, enableTS := parseInputFlags()
	if *useExternal {
		connection := client.ExternalConnection()

		redisClient := client.RedisClient()
		redisTSClient := client.RedisTSClient()

		ginRouter := middleware.GinRouter(connection, redisClient, redisTSClient)
		registerEndpoints(ginRouter)

		if *enableTS {
			writes.StartNodeWrites(connection, redisClient, redisTSClient)
		}

		ginRouter.Run(":8080")
	} else {
		client.Connection()
	}
}
