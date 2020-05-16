package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

// GetNodes ... returns basic node data in JSON format
func GetPods(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	namespaces, _ := connection.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": namespaces,
	})
}

func GetPod(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	nodes, _ := connection.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": nodes,
	})
}
