package main

import (
	"flag"
	"github.com/stevenpg/kubutilization/api/pkg/client"
)

// initFlags - Initializes are parses flags, returning the ptrs to caller
func initFlags() *bool {
	ExternalKubernetesPtr := flag.Bool("external", false, "a boolean selector for whether to search for a kubeconfig in ~/.kube/config")
	flag.Parse()

	return ExternalKubernetesPtr
}

func main() {
	if *initFlags() {
		client.ExternalConnection()
	} else {
		client.Connection()
	}
}
