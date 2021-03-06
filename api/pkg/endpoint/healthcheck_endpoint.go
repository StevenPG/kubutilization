package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
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
- [ ] Add healthcheck for connection to Redis
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

// RedisHealthcheck ... Perform a new client connection to redis and perform ping
func RedisHealthcheck(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		// TODO - external this with Viper, pull universal configuration from gin context
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, _ := client.Ping().Result()

	c.JSON(http.StatusOK, gin.H{
		"Ping": pong,
	})
}
