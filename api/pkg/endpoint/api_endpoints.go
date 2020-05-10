package endpoint

import "github.com/gin-gonic/gin"

// RegisterNodeEndpoints ... receive gin engine, apply node endpoints
func RegisterNodeEndpoints(gin *gin.Engine) {
	gin.GET("/api/v1/nodes", nil)
}

// RegisterPodEndpoints ... receive gin engine, apply pod endpoints
func RegisterPodEndpoints(gin *gin.Engine) {
	gin.GET("/api/v1/pods", nil)
}
