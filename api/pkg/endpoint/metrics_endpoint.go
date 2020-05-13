package endpoint

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

// MetricsRoot ... Returns the JSON result of hitting the kubernetes metrics-server root endpoint
func MetricsRoot(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1").DoRaw(context.TODO())
	marshalAndSetJson(c, data)
}

func NodesEndpoint(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw(context.TODO())
	marshalAndSetJson(c, data)
}

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
