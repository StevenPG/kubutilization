package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

/**
* pod_endpoints.go
* Contains all pod endpoints
*
* Capabilities
- [X] Return all pods within a namespace
- [X] Return all pods in cluster
- [X] Return specific input pod within a given namespace
*/

// GetPods ... returns basic node data in JSON format
func GetPods(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	namespaces, _ := connection.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	var podsListArray []v1.PodList
	for namespaceIndex := range namespaces.Items {
		pods, _ := connection.CoreV1().Pods(namespaces.Items[namespaceIndex].Name).List(context.TODO(), metav1.ListOptions{})
		podsListArray = append(podsListArray, *pods)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": podsListArray,
	})
}

// GetPodsWithNamespace ... get all pods from within the supplied namespace
func GetPodsWithNamespace(c *gin.Context) {
	namespace := c.Param("namespace")
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	pods, _ := connection.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": pods,
	})
}

// GetPodWithNamespace ... get specific pod from within supplied namespace
func GetPodWithNamespace(c *gin.Context) {
	namespace := c.Param("namespace")
	pod := c.Param("pod")
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	result, _ := connection.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
