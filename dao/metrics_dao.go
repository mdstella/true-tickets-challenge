//Package dao will represent the persistent layer, for the challenge purpose
//we will do all the persistence in memory, but with the interface defined
// easily we can add any storage/database
package dao

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/mdstella/true-tickets-challenge/model"
)

// MetricDao is the interface that will have all the definitions for
// the dao layer
//go:generate mockery -name=MetricDao
type MetricDao interface {
	StoreMetric(key string, value int) error
	GetMetricSumByKey(key string) ([]model.MetricDto, error)
}

// MetricDaoMemoryImpl will be the implementation for the MetricDao
type MetricDaoMemoryImpl struct {
	// representation of the memory storage for metrics
	// will be a map that will have the metric key as and an array with MetricDto as value.
	metrics map[string][]model.MetricDto
	// mutex to access the map on a sync way
	metricsMutex *sync.RWMutex
	// the TTL that will be used to know if we need to retrieve or not a metric from the storage
	ttl int64
}

// NewMetricDaoMemoryImpl is used as a constructor, invoked at service startup time will retrieve an in memory DAO
func NewMetricDaoMemoryImpl(TTL int64) MetricDao {
	return &MetricDaoMemoryImpl{
		metrics:      make(map[string][]model.MetricDto),
		metricsMutex: &sync.RWMutex{},
		ttl:          TTL,
	}
}

// StoreMetric will be used to store the metrics in this case in the memory map
func (dao *MetricDaoMemoryImpl) StoreMetric(key string, value int) error {
	// creating the new metric that will be stored
	metric := model.MetricDto{
		Key:       key,
		Value:     value,
		CreatedAt: time.Now().UTC().Unix(),
	}
	dao.metricsMutex.Lock()
	defer dao.metricsMutex.Unlock()
	if metrics, ok := dao.metrics[key]; ok {
		metrics = append(metrics, metric)
		dao.metrics[key] = metrics
	} else {
		dao.metrics[key] = []model.MetricDto{metric}
	}

	// Adding this print line just to see on the console
	// the information inside the map
	fmt.Println(dao.metrics)
	return nil
}

// GetMetricSumByKey will retrieve the metric sum by the given metric key
func (dao *MetricDaoMemoryImpl) GetMetricSumByKey(key string) ([]model.MetricDto, error) {
	emptyResult := []model.MetricDto{}
	// locking as we are going to manipulate the metrics in the memory map
	// this won't be needed on a real environment with a DB as we will be able to
	// filter the metrics when querying the DB using the TTL
	dao.metricsMutex.Lock()
	defer dao.metricsMutex.Unlock()
	metrics, ok := dao.metrics[key]

	if ok && len(metrics) > 0 {
		// key found we have metrics to retrieve/evaluate
		filteredMetrics := make([]model.MetricDto, 0)
		now := time.Now().UTC().Unix()
		for _, metric := range metrics {

			if math.Abs(float64(now-dao.ttl)) <= float64(metric.CreatedAt) {
				filteredMetrics = append(filteredMetrics, metric)
			} else {
				fmt.Println(fmt.Sprintf("Discarding metric: %s", metric))
			}
		}
		// adding it again to the map, discarding the ones that are expired
		dao.metrics[key] = filteredMetrics
		return filteredMetrics, nil
	}
	// no metrics to retrieve
	return emptyResult, nil
}
