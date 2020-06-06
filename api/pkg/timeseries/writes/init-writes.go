package writes

import (
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/go-redis/redis/v7"
	"k8s.io/client-go/kubernetes"
)

// StartNodeWrites ... Trigger goroutines to start writing node metrics data to Redis
func StartNodeWrites(clientset *kubernetes.Clientset, client *redis.Client, tsClient *redistimeseries.Client) {
	go PollAndPushMetrics(clientset, client, tsClient)
}

// StartPodWrites ... Trigger goroutines to start writing pod metrics data to Redis
func StartPodWrites() {

}
