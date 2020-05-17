package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

/**
* healthcheck_endpoint.go
* Return specific healthchecks for the application
*
* Capabilities
- [X] Provide health-check displaying the application is up and running
- [X] Provide health-check that API is successfully connected to a Kubernetes cluster
*/

// HealthCheck ... returns function that handles incoming health check requests with 200 OK
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}

// KubernetesConnectionHealthCheck ... returns function that
//		handles health check and pings Kubernetes to check connection
func KubernetesConnectionHealthCheck(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	result, _ := connection.CoreV1().Namespaces().Get(context.TODO(), "kube-system", metav1.GetOptions{})
	c.JSON(http.StatusOK, gin.H{
		"Status":              "OK",
		"operation":           "Checking if kube-system namespace exists...",
		"kube-system-details": result,
	})

}
