package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// HomeDir - returns the running user's home directory
func HomeDir() string {
	if linuxHome := os.Getenv("HOME"); linuxHome != "" {
		return linuxHome
	}
	return os.Getenv("USERPROFILE") // windows-fallback
}

// MarshalAndSetJson ... helper method that un-marshals JSON and returns to browser
func MarshalAndSetJson(c *gin.Context, data []byte) {
	var raw map[string]interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Something went wrong unmarshalling the data from kubernetes!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": raw,
		})
	}
}

// MarshalAndSetJson ... helper method that un-marshals JSON and returns to browser
func MarshalAndSetJsonSansGin(data []byte) (map[string]interface{}, error) {
	var raw map[string]interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return raw, err
	} else {
		return raw, err
	}
}
