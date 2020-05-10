package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck ... returns function that handles incoming health check requests with 200 OK
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}
