package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetNodes ... returns basic node data in JSON format
func GetNodes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}
