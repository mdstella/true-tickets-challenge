package model

import (
	"time"
)

// DTO MODEL -- Internal model that will be used for persistence and logic on the application

// MetricDto will have the metric information
type MetricDto struct {
	// metric key
	Key string
	// metric value
	Value int
	// time to live / time to consider the metric (should be 60 minutes)
	TTL time.Duration
	// Created at value that is an integer value
	CreatedAt int64
}
