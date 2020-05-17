package writes

import (
	"context"
	"encoding/json"
	"fmt"
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/go-redis/redis/v7"
	"github.com/stevenpg/kubutilization/api/pkg/timeseries/models"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"strings"
	"time"
)

// PollAndPushMetrics ... wait every N seconds and push metrics
func PollAndPushMetrics(clientset *kubernetes.Clientset, redisConnection *redis.Client, timeSeriesConnection *redistimeseries.Client) {

	var unlimitedRuntime = true
	for ok := true; ok; ok = unlimitedRuntime {

		// TODO - make time configurable via Viper config file
		time.Sleep(30 * time.Second)
		redisConnection.Ping()

		data, _ := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw(context.TODO())
		var rawMetrics models.NodeMetrics

		json.Unmarshal(data, &rawMetrics)
		redisEntities := buildEntities(rawMetrics)
		writeEntities(timeSeriesConnection, redisEntities)
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
func writeEntities(connection *redistimeseries.Client, entities []models.NodeMetricsEntity) {
	CheckAndInitTSKeys(connection, entities)
	for i := range entities {
		cpuEntityKey := generateCpuKey(entities[i].MetricsType, entities[i].Name)
		memEntityKey := generateMemKey(entities[i].MetricsType, entities[i].Name)

		jsonEntity, _ := json.Marshal(entities[i])
		// TODO - set global duration configuration property, default 6h

		fmt.Println(fmt.Sprintf("Writing cpu metrics for -> %s", cpuEntityKey))
		fmt.Println(fmt.Sprintf("Writing mem metrics for -> %s", memEntityKey))

		cpuUsage := strings.TrimSuffix(entities[i].Usage.CPU, "n")
		memUsage := strings.TrimSuffix(entities[i].Usage.Memory, "Ki")

		cpuAsFloat, err := strconv.ParseFloat(cpuUsage, 64)
		if err != nil {
			fmt.Println(err)
		}
		memAsFloat, err := strconv.ParseFloat(memUsage, 64)
		if err != nil {
			fmt.Println(err)
		}

		connection.AddAutoTs(cpuEntityKey, cpuAsFloat)
		connection.AddAutoTs(memEntityKey, memAsFloat)

		fmt.Println(string(jsonEntity))
	}
}

// CheckAndInitKeys ... Receives a redis time series connection, verifies and creates keys if they don't exist
func CheckAndInitTSKeys(connection *redistimeseries.Client, entities []models.NodeMetricsEntity) {
	for i := range entities {
		cpuEntityKey := generateCpuKey(entities[i].MetricsType, entities[i].Name)
		_, cpuIsPresent := connection.Info(cpuEntityKey)
		if cpuIsPresent != nil {
			// TODO - make timeouts configurable
			connection.CreateKeyWithOptions(cpuEntityKey, redistimeseries.CreateOptions{RetentionMSecs: 6 * time.Hour})
			connection.CreateKeyWithOptions(cpuEntityKey+"_avg", redistimeseries.CreateOptions{RetentionMSecs: 6 * time.Hour})
			connection.CreateRule(cpuEntityKey, redistimeseries.AvgAggregation, 60, cpuEntityKey+"_avg")
		}

		memEntityKey := generateMemKey(entities[i].MetricsType, entities[i].Name)
		_, memIsPresent := connection.Info(memEntityKey)
		if memIsPresent != nil {
			// TODO - make timeouts configurable
			connection.CreateKeyWithOptions(memEntityKey, redistimeseries.CreateOptions{RetentionMSecs: 6 * time.Hour})
			connection.CreateKeyWithOptions(memEntityKey+"_avg", redistimeseries.CreateOptions{RetentionMSecs: 6 * time.Hour})
			connection.CreateRule(memEntityKey, redistimeseries.AvgAggregation, 60, memEntityKey+"_avg")
		}
	}
}

// Helper functions to keep functions aligned when generating CPU key
func generateCpuKey(metricsType string, name string) string {
	return fmt.Sprintf("%s:%s:CPU", metricsType, name)
}

// Helper functions to keep functions aligned when generating Mem key
func generateMemKey(metricsType string, name string) string {
	return fmt.Sprintf("%s:%s:Memory", metricsType, name)
}
