package dao_test

import (
	"testing"
	"time"

	"github.com/mdstella/true-tickets-challenge/dao"
	"github.com/stretchr/testify/assert"
)

// In this file will show how can we invoke the dao unit tests
func Test_Complete_Flow(t *testing.T) {
	assert := assert.New(t)

	dao := dao.NewMetricDaoMemoryImpl(3) // Setting the TTL in 3 seconds for testing purpose

	// T0
	err := dao.StoreMetric("metric1", 1)
	assert.Nil(err)

	err = dao.StoreMetric("metric2", 3)
	assert.Nil(err)

	err = dao.StoreMetric("metric2", 1)
	assert.Nil(err)

	err = dao.StoreMetric("metric3", -2)
	assert.Nil(err)

	// Checking the metrics
	metrics, err := dao.GetMetricSumByKey("metric1")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 1)
	assert.Equal(1, metrics[0].Value)

	metrics, err = dao.GetMetricSumByKey("metric2")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 2)
	assert.Equal(3, metrics[0].Value)
	assert.Equal(1, metrics[1].Value)

	metrics, err = dao.GetMetricSumByKey("metric3")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 1)
	assert.Equal(-2, metrics[0].Value)

	// sleeping 3 secs
	time.Sleep(time.Second * 2)

	// T1 - adding more values to the metrics
	err = dao.StoreMetric("metric1", 10)
	assert.Nil(err)

	err = dao.StoreMetric("metric2", -2)
	assert.Nil(err)

	err = dao.StoreMetric("metric3", 10)
	assert.Nil(err)

	err = dao.StoreMetric("metric3", 22)
	assert.Nil(err)

	// Checking the metrics again
	metrics, err = dao.GetMetricSumByKey("metric1")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 2)
	assert.Equal(1, metrics[0].Value)
	assert.Equal(10, metrics[1].Value)

	metrics, err = dao.GetMetricSumByKey("metric2")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 3)
	assert.Equal(3, metrics[0].Value)
	assert.Equal(1, metrics[1].Value)
	assert.Equal(-2, metrics[2].Value)

	metrics, err = dao.GetMetricSumByKey("metric3")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 3)
	assert.Equal(-2, metrics[0].Value)
	assert.Equal(10, metrics[1].Value)
	assert.Equal(22, metrics[2].Value)

	// sleeping 3 more seconds and the metrics added on T0 should be removed
	time.Sleep(time.Second * 2)

	// T2 - Checking the metrics again
	metrics, err = dao.GetMetricSumByKey("metric1")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 1)
	assert.Equal(10, metrics[0].Value)

	metrics, err = dao.GetMetricSumByKey("metric2")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 1)
	assert.Equal(-2, metrics[0].Value)

	metrics, err = dao.GetMetricSumByKey("metric3")
	assert.Nil(err)
	assert.NotEmpty(metrics)
	assert.Len(metrics, 2)
	assert.Equal(10, metrics[0].Value)
	assert.Equal(22, metrics[1].Value)

	// sleeping 5 more seconds and there shouldn't be more metrics
	time.Sleep(time.Second * 3)

	metrics, err = dao.GetMetricSumByKey("metric1")
	assert.Nil(err)
	assert.Empty(metrics)

	metrics, err = dao.GetMetricSumByKey("metric2")
	assert.Nil(err)
	assert.Empty(metrics)

	metrics, err = dao.GetMetricSumByKey("metric3")
	assert.Nil(err)
	assert.Empty(metrics)
}

func Test_EmptyMetrics_NotStoredKey(t *testing.T) {
	assert := assert.New(t)

	dao := dao.NewMetricDaoMemoryImpl(3600)

	metrics, err := dao.GetMetricSumByKey("metric")
	assert.Nil(err)
	assert.Empty(metrics)
}
