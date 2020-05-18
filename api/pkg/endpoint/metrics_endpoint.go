package endpoint

import (
	"context"
	"fmt"
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/gin-gonic/gin"
	"github.com/stevenpg/kubutilization/api/pkg/client"
	"github.com/stevenpg/kubutilization/api/pkg/timeseries/models"
	"github.com/stevenpg/kubutilization/api/pkg/timeseries/writes"
	"k8s.io/client-go/kubernetes"
	"math"
	"net/http"
)

/**
* metrics_endpoint.go
* Contains all metrics endpoints
*
* Capabilities
- [X] Provide metrics root at apis/metrics.k8s.io/v1beta1
- [X] Provide metrics for nodes
- [X] Provide metrics for specific node
- [ ] Provide metrics for pods in supplied namespace
- [X] Provide metrics for pods across all namespaces
- [ ] Provide metrics for a specific pod within a given namespace

- [ ] Return all historical metrics from Redis for node utilization data
- [ ] Return all historical metrics from Redis for pod utilization data
*/

// MetricsRoot ... Returns the JSON result of hitting the kubernetes metrics-server root endpoint
func MetricsRoot(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1").DoRaw(context.TODO())
	client.MarshalAndSetJson(c, data)
}

// NodesRoot ... Return the nodes endpoint
func NodesRoot(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw(context.TODO())
	client.MarshalAndSetJson(c, data)
}

// SpecificNode ... Returns the metrics results for a specific node
func SpecificNode(c *gin.Context) {
	node := c.Param("node")
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes/" + node).DoRaw(context.TODO())
	client.MarshalAndSetJson(c, data)
}

// PodsRoot ... Return pods metrics endpoint from all namespaces
func PodsRoot(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.TODO())
	client.MarshalAndSetJson(c, data)
}

// TODO
// HistoricalNodeCPU ... returns Node CPU data in time-series format for given Node
func HistoricalNodeCPU(c *gin.Context) {
	node := c.Param("node")
	c.JSON(http.StatusInternalServerError, gin.H{
		"result": node,
	})
}

// TODO
// AverageNodeCPU ... Gives average over last 60s of time series data for input node CPU
func AverageNodeCPU(c *gin.Context) {
	tsRedis := c.MustGet("tsRedis").(*redistimeseries.Client)
	node := c.Param("node")
	endpoint := fmt.Sprintf("%s_avg", writes.GenerateCpuKey("Node", node))
	dataPoint, _ := tsRedis.Get(endpoint)
	c.JSON(http.StatusOK, gin.H{
		"node":          node,
		"raw_format":    "nanocpu",
		"raw_result":    *dataPoint,
		"value_in_mCpu": convertNanoCPUToMicroCPU(dataPoint.Value),
	})
}

//TODO
// HistoricalNodeMem ... returns Node Mem data in time-series format for given Node
func HistoricalNodeMem(c *gin.Context) {
	node := c.Param("node")
	var timeRange models.TimeRangeURLInput
	if c.Bind(&timeRange) == nil {
		if timeRange.From == 0 || timeRange.To == 0 {
			// Ignore Inputs
			c.JSON(http.StatusOK, gin.H{
				"result":        node,
				"single-result": "stuff and things",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": node,
				"from":   timeRange.From,
				"to":     timeRange.To,
			})
		}

	}
}

// AverageNodeMem ... Gives average over last 60s of time series data for input node Mem
func AverageNodeMem(c *gin.Context) {
	tsRedis := c.MustGet("tsRedis").(*redistimeseries.Client)
	node := c.Param("node")
	endpoint := fmt.Sprintf("%s_avg", writes.GenerateMemKey("Node", node))
	dataPoint, _ := tsRedis.Get(endpoint)
	c.JSON(http.StatusOK, gin.H{
		"node":        node,
		"raw_format":  "Ki",
		"raw_result":  *dataPoint,
		"value_in_Mi": convertKiToMiHelper(dataPoint.Value),
		"value_in_MB": convertKiToMBHelper(dataPoint.Value),
	})
}

func convertNanoCPUToMicroCPU(nanoCpu float64) float64 {
	return math.Round(nanoCpu / 1_000_000)
}

func convertKiToMiHelper(kibibytes float64) float64 {
	return math.Round(kibibytes / 1024)
}

func convertKiToMBHelper(kibibytes float64) float64 {
	return math.Round(kibibytes / 977)
}
