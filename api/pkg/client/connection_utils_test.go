package client_test

import (
	"fmt"
	"github.com/stevenpg/kubutilization/api/pkg/client"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestHomeDir - test homedir function for returning variable value
func TestHomeDir(t *testing.T) {

	assert := assert.New(t)
	homeDir := client.HomeDir()
	fmt.Printf("Retrieving %s\n", homeDir)
	assert.True(homeDir == os.Getenv("USERPROFILE") || homeDir == os.Getenv("HOME"))
}
