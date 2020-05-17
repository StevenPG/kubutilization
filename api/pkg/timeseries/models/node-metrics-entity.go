package models

import "time"

type NodeMetricsEntity struct {
	Timestamp   time.Time `json:"timestamp"`
	MetricsType string    `json:"type"`
	Name        string    `json:"name"`

	Usage struct {
		CPU    string `json:"cpu"`
		Memory string `json:"memory"`
	} `json:"usage"`
}
