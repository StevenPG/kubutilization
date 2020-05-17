package endpoint

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

/**
* metrics_endpoint.go
* Contains all metrics endpoints
*
* Capabilities
- [X] Provide metrics root at apis/metrics.k8s.io/v1beta1
- [X] Provide metrics for nodes
- [X] Provide metrics for specific node
- [ ] Provide metrics for pods in supplied namespace
- [X] Provide metrics for pods across all namespaces
- [ ] Provide metrics for a specific pod within a given namespace
*/

// MetricsRoot ... Returns the JSON result of hitting the kubernetes metrics-server root endpoint
func MetricsRoot(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1").DoRaw(context.TODO())
	marshalAndSetJson(c, data)
}

// NodesRoot ... Return the nodes endpoint
func NodesRoot(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw(context.TODO())
	marshalAndSetJson(c, data)
}

// SpecificNode ... Returns the metrics results for a specific node
func SpecificNode(c *gin.Context) {
	node := c.Param("node")
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes/" + node).DoRaw(context.TODO())
	marshalAndSetJson(c, data)
}

// PodsRoot ... Return pods metrics endpoint from all namespaces
func PodsRoot(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.TODO())
	marshalAndSetJson(c, data)
}

// helper method that un-marshals JSON and returns to browser
func marshalAndSetJson(c *gin.Context, data []byte) {
	var raw map[string]interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Something went wrong unmarshalling the data from kubernetes!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": raw,
		})
	}
}
