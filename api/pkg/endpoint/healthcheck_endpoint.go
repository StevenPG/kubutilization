package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* healthcheck_endpoint.go
* Return specific healthchecks for the application
*
* Capabilities
- [X] Provide health-check displaying the application is up and running
- [ ] Provide health-check that API is successfully connected to a Kubernetes cluster
*/

// HealthCheck ... returns function that handles incoming health check requests with 200 OK
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}
