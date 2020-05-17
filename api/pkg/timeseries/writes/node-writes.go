package writes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/stevenpg/kubutilization/api/pkg/timeseries/models"
	"k8s.io/client-go/kubernetes"
	"time"
)

var (
	// Utilized to generate unique IDs for node writes
	counter = 0
)

// PollAndPushMetrics ... wait every N seconds and push metrics
func PollAndPushMetrics(clientset *kubernetes.Clientset) {
	redisConnection := metricsInitHelper()

	var unlimitedRuntime = true
	for ok := true; ok; ok = unlimitedRuntime {

		// TODO - make time configurable via Viper config file
		time.Sleep(30 * time.Second)
		redisConnection.Ping()

		data, _ := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw(context.TODO())
		var rawMetrics models.NodeMetrics

		json.Unmarshal(data, &rawMetrics)
		redisEntities := buildEntities(rawMetrics)
		writeEntities(redisConnection, redisEntities)
	}
}

// Receives a list of node metrics and returns a list of Redis entities to write as JSON
func buildEntities(metrics models.NodeMetrics) []models.NodeMetricsEntity {
	var entities []models.NodeMetricsEntity
	for i := range metrics.Items {
		var entity models.NodeMetricsEntity
		metricsModel := metrics.Items[i]

		entity.Name = metricsModel.Metadata.Name
		entity.Timestamp = time.Now()
		entity.MetricsType = "Node"
		entity.Usage.CPU = metricsModel.Usage.CPU
		entity.Usage.Memory = metricsModel.Usage.Memory
		entities = append(entities, entity)
	}
	return entities
}

// Receives a list of Redis entities to convert
func writeEntities(connection *redis.Client, entities []models.NodeMetricsEntity) {
	for i := range entities {
		counter += 1
		entityKey := fmt.Sprintf("%s:%s:%d", entities[i].MetricsType, entities[i].Name, counter)

		jsonEntity, _ := json.Marshal(entities[i])
		// TODO - set global duration configuration property, default 6h
		connection.Set(entityKey, string(jsonEntity), time.Duration(time.Duration.Hours(6)))
	}
}

// Helper method that instantiates Redis client connection
func metricsInitHelper() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
