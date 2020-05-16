package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

// GetNodes ... returns basic node data in JSON format
func GetNamespaces(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	namespaces, _ := connection.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": namespaces,
	})
}

// GetNode ... get basic node data from input parameter
func GetNamespace(c *gin.Context) {
	namespaceParam := c.Param("namespace")
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	namespace, _ := connection.CoreV1().Namespaces().Get(context.TODO(), namespaceParam, metav1.GetOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": namespace,
	})
}
