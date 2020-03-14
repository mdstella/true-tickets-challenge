package model

import "fmt"

// DTO MODEL -- Internal model that will be used for persistence and logic on the application

// MetricDto will have the metric information
type MetricDto struct {
	// metric key
	Key string
	// metric value
	Value int
	// Created at value that is an integer value
	CreatedAt int64
}

// ToString - metric string representation
func (m MetricDto) String() string {
	return fmt.Sprintf("Metric--> key: %s - value: %d", m.Key, m.Value)
}
