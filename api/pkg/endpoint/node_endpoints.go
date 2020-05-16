package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

/**
* node_endpoints.go
* Contains all node endpoints
*
* Capabilities
- [X] Return all nodes in cluster
- [X] Return specific node from cluster
*/

// GetNodes ... returns basic node data in JSON format
func GetNodes(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	nodes, _ := connection.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": nodes,
	})
}

// GetNode ... get basic node data from input parameter
func GetNode(c *gin.Context) {
	nodeParam := c.Param("node")
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	node, _ := connection.CoreV1().Nodes().Get(context.TODO(), nodeParam, metav1.GetOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": node,
	})
}
