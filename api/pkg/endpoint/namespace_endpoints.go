package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

/**
* namespace_endpoints.go
* Contains all namespace endpoints
*
* Capabilities
- [X] Return all namespaces from kubernetes API
- [X] Return content for individual namespace
*/

// GetNamespaces ... return all namespaces in cluster
func GetNamespaces(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	namespaces, _ := connection.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": namespaces,
	})
}

// GetNamespace ... return namespace data from specific input namespace
func GetNamespace(c *gin.Context) {
	namespaceParam := c.Param("namespace")
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	namespace, _ := connection.CoreV1().Namespaces().Get(context.TODO(), namespaceParam, metav1.GetOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": namespace,
	})
}
