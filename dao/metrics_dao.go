//Package dao will represent the persistent layer, for the challenge purpose
//we will do all the persistence in memory, but with the interface defined
// easily we can add any storage/database
package dao

import (
	"fmt"
	"sync"
	"time"

	"github.com/mdstella/true-tickets-challenge/model"
)

const (
	// added as a constant, but on a productive environment this
	// should be a property value externalized to avoid changing
	// code if we need to increase/decrease this value
	defaultTTL = 60 * time.Minute
)

// MetricDao is the interface that will have all the definitions for
// the dao layer
//go:generate mockery -name=MetricDao
type MetricDao interface {
	StoreMetric(key string, value int) error
}

// MetricDaoMemoryImpl will be the implementation for the MetricDao
type MetricDaoMemoryImpl struct {
	// representation of the memory storage for metrics
	// will be a map that will have the metric key as and an array with MetricDto as value.
	metrics map[string][]model.MetricDto
	// mutex to access the map on a sync way
	metricsMutex *sync.RWMutex
}

// NewMetricDaoMemoryImpl is used as a constructor, invoked at service startup time will retrieve an in memory DAO
func NewMetricDaoMemoryImpl() MetricDao {
	return &MetricDaoMemoryImpl{
		metrics:      make(map[string][]model.MetricDto),
		metricsMutex: &sync.RWMutex{},
	}
}

// StoreMetric will be used to store the metrics in this case in the memory map
func (dao *MetricDaoMemoryImpl) StoreMetric(key string, value int) error {
	// creating the new metric that will be stored
	metric := model.MetricDto{
		Key:       key,
		Value:     value,
		TTL:       defaultTTL,
		CreatedAt: time.Now().UTC().Unix(),
	}
	dao.metricsMutex.Lock()
	if metrics, ok := dao.metrics[key]; ok {
		metrics = append(metrics, metric)
		dao.metrics[key] = metrics
	} else {
		dao.metrics[key] = []model.MetricDto{metric}
	}

	// Adding this print line just to see on the console
	// the information inside the map
	fmt.Println(dao.metrics)

	dao.metricsMutex.Unlock()
	return nil
}
