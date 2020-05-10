package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetNodes ... returns basic node data in JSON format
func GetPods(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", "")
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
		"test":   namespace,
	})
}

func GetPod(c *gin.Context) {
	pod := c.Param("pod")
	c.JSON(http.StatusOK, gin.H{
		"Status":      "OK",
		"selectedPod": pod,
	})
}
