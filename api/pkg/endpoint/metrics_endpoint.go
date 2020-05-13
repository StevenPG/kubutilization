package endpoint

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func CallRootMetricsApi(c *gin.Context) {
	connection := c.MustGet("clientset").(kubernetes.Clientset)
	data, _ := connection.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1").DoRaw(context.TODO())
	fmt.Println(data)
}
