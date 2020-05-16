package endpoint

import "github.com/gin-gonic/gin"

// RegisterUtilityEndpoints ... receive gin engine, apply utility endpoints (eg. healthcheck)
func RegisterUtilityEndpoints(gin *gin.Engine) {
	gin.GET("/api/health", HealthCheck)
	gin.GET("/api/kubehealth", KubernetesConnectionHealthCheck)
}

// RegisterMetricsEndpoints ... return the root of the metrics endpoint
func RegisterMetricsEndpoints(gin *gin.Engine) {
	gin.GET("/api/v1/metrics", MetricsRoot)
	gin.GET("/api/v1/metrics/nodes", NodesRoot)
	gin.GET("/api/v1/metrics/nodes/:node", SpecificNode)
	gin.GET("/api/v1/metrics/pods", PodsRoot)

}

func RegisterNamespaceEndpoints(gin *gin.Engine) {
	gin.GET("/api/v1/namespaces", GetNamespaces)
	gin.GET("/api/v1/namespaces/:namespace", GetNamespace)
}

// RegisterNodeEndpoints ... receive gin engine, apply node endpoints
func RegisterNodeEndpoints(gin *gin.Engine) {
	gin.GET("/api/v1/nodes", GetNodes)
	gin.GET("/api/v1/nodes/:node", GetNode)
}

// RegisterPodEndpoints ... receive gin engine, apply pod endpoints
func RegisterPodEndpoints(gin *gin.Engine) {
	gin.GET("/api/v1/pods", GetPods)
	gin.GET("/api/v1/namespaces/:namespace/pods", GetPodsWithNamespace)
	gin.GET("/api/v1/namespaces/:namespace/pods/:pod", GetPodWithNamespace)
}
