package middleware

import (
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"

	"k8s.io/client-go/kubernetes"
	"time"
)

// GinContext ... Returns a gin context for usage by caller
func GinRouter(clientset *kubernetes.Clientset, client *redis.Client, tsClient *redistimeseries.Client) *gin.Engine {
	ginEngine := gin.Default()

	ginEngine.Use(kubeContextMiddleware(*clientset))
	ginEngine.Use(redisContextMiddleware(client, tsClient))

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

// rediscontextMiddleware ... add redis standard and TS client to gin context
func redisContextMiddleware(redis *redis.Client, tsRedis *redistimeseries.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("redis", redis)
		c.Set("tsRedis", tsRedis)
		c.Next()
	}
}
