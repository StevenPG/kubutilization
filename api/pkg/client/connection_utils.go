package client

import "os"

// HomeDir - returns the running user's home directory
func HomeDir() string {
	if linuxHome := os.Getenv("HOME"); linuxHome != "" {
		return linuxHome
	}
	return os.Getenv("USERPROFILE") // windows-fallback
}
