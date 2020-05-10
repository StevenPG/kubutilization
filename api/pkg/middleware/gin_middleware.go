package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"time"
)

// GinContext ... Returns a gin context for usage by caller
func GinRouter(clientset *kubernetes.Clientset) *gin.Engine {
	ginEngine := gin.Default()

	ginEngine.Use(kubeContextMiddleware(*clientset))

	ginEngine.Use(static.Serve("/", static.LocalFile("../../../ui/dist", false)))
	ginEngine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET"},
		AllowHeaders:    []string{"Origin"},
		MaxAge:          12 * time.Hour,
	}))

	return ginEngine
}

//kubeContextMiddleware ... add kubernetes clientset to context
func kubeContextMiddleware(clientset kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("clientset", clientset)
		c.Next()
	}
}
