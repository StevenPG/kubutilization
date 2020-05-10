package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetNodes ... returns basic node data in JSON format
func GetPods(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}

func GetPod(c *gin.Context) {
	pod := c.Param("pod")
	c.JSON(http.StatusOK, gin.H{
		"Status":      "OK",
		"selectedPod": pod,
	})
}
