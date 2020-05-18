package models

import "time"

type NodeMetrics struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`

	Metadata struct {
		SelfLink string `json:"selfLink"`
	} `json:"metadata"`

	Items []struct {
		Metadata struct {
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Name              string    `json:"name"`
			SelfLink          string    `json:"selfLink"`
		} `json:"metadata"`

		Timestamp time.Time `json:"timestamp"`

		Usage struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"usage"`

		Window string `json:"window"`
	} `json:"items"`
}
