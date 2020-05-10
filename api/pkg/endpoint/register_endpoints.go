package endpoint

import "github.com/gin-gonic/gin"

// RegisterUtilityEndpoints ... receive gin engine, apply utility endpoints (eg. healthcheck)
func RegisterUtilityEndpoints(gin *gin.Engine) {
	gin.GET("/api/health", HealthCheck)
}

// RegisterNodeEndpoints ... receive gin engine, apply node endpoints
func RegisterNodeEndpoints(gin *gin.Engine) {
	gin.GET("/api/v1/nodes", GetNodes)
	gin.GET("/api/v1/nodes/:node", GetNode)
}

// RegisterPodEndpoints ... receive gin engine, apply pod endpoints
func RegisterPodEndpoints(gin *gin.Engine) {
	gin.GET("/api/v1/pods", GetPods)
	gin.GET("/api/v1/pods/:pod", GetPod)
}
