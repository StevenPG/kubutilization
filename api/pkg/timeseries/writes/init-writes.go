package writes

import (
	"k8s.io/client-go/kubernetes"
)

// StartNodeWrites ... Trigger goroutines to start writing node metrics data to Redis
func StartNodeWrites(clientset *kubernetes.Clientset) {
	go PollAndPushMetrics(clientset)
}

// StartPodWrites ... Trigger goroutines to start writing pod metrics data to Redis
func StartPodWrites() {

}
