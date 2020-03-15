package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/mdstella/true-tickets-challenge/model"

	"github.com/mdstella/true-tickets-challenge/dao"
	m_dao "github.com/mdstella/true-tickets-challenge/dao/mocks"
	"github.com/mdstella/true-tickets-challenge/service"
	"github.com/stretchr/testify/assert"
)

// In this file will show how can we invoke the dao and run an integration test over the 2 layers and also how to mock the dao
// layer to run unit tests
func Test_Complete_Flow(t *testing.T) {
	assert := assert.New(t)

	dao := dao.NewMetricDaoMemoryImpl(3) // Setting the TTL in 3 seconds for testing purpose
	srv := service.NewMetricsServiceImpl(dao)

	// T0
	err := srv.AddMetric("metric1", 1)
	assert.Nil(err)

	err = srv.AddMetric("metric2", 3)
	assert.Nil(err)

	err = srv.AddMetric("metric2", 1)
	assert.Nil(err)

	err = srv.AddMetric("metric3", -2)
	assert.Nil(err)

	// Checking the metrics
	val, err := srv.SumMetric("metric1")
	assert.Nil(err)
	assert.Equal(1, val)

	val, err = srv.SumMetric("metric2")
	assert.Nil(err)
	assert.Equal(4, val)

	val, err = srv.SumMetric("metric3")
	assert.Nil(err)
	assert.Equal(-2, val)

	// sleeping 3 secs
	time.Sleep(time.Second * 2)

	// T1 - adding more values to the metrics
	err = srv.AddMetric("metric1", 10)
	assert.Nil(err)

	err = srv.AddMetric("metric2", -2)
	assert.Nil(err)

	err = srv.AddMetric("metric3", 10)
	assert.Nil(err)

	err = srv.AddMetric("metric3", 22)
	assert.Nil(err)

	// Checking the metrics again
	val, err = srv.SumMetric("metric1")
	assert.Nil(err)
	assert.Equal(11, val)

	val, err = srv.SumMetric("metric2")
	assert.Nil(err)
	assert.Equal(2, val)

	val, err = srv.SumMetric("metric3")
	assert.Nil(err)
	assert.Equal(30, val)

	// sleeping 3 more seconds and the metrics added on T0 should be removed
	time.Sleep(time.Second * 2)

	// T2 - Checking the metrics again
	val, err = srv.SumMetric("metric1")
	assert.Nil(err)
	assert.Equal(10, val)

	val, err = srv.SumMetric("metric2")
	assert.Nil(err)
	assert.Equal(-2, val)

	val, err = srv.SumMetric("metric3")
	assert.Nil(err)
	assert.Equal(32, val)

	// sleeping 5 more seconds and there shouldn't be more metrics
	time.Sleep(time.Second * 3)

	val, err = srv.SumMetric("metric1")
	assert.Nil(err)
	assert.Equal(0, val)

	val, err = srv.SumMetric("metric2")
	assert.Nil(err)
	assert.Equal(0, val)

	val, err = srv.SumMetric("metric3")
	assert.Nil(err)
	assert.Equal(0, val)
}

func Test_AddMetric_Error_No_Key(t *testing.T) {
	assert := assert.New(t)

	dao := dao.NewMetricDaoMemoryImpl(3600)
	srv := service.NewMetricsServiceImpl(dao)

	err := srv.AddMetric("", 1)
	assert.NotNil(err)
}

func Test_AddMetric_Error_Zero_Value(t *testing.T) {
	assert := assert.New(t)

	dao := dao.NewMetricDaoMemoryImpl(3600)
	srv := service.NewMetricsServiceImpl(dao)

	err := srv.AddMetric("metric", 0)
	assert.NotNil(err)
}

func Test_SumMetric_Error_No_Key(t *testing.T) {
	assert := assert.New(t)

	dao := dao.NewMetricDaoMemoryImpl(3600)
	srv := service.NewMetricsServiceImpl(dao)

	// T0
	err := srv.AddMetric("metric1", 1)
	assert.Nil(err)

	_, err = srv.SumMetric("")
	assert.NotNil(err)
}

// Adding a test case mocking the DAO to be able to test a failure no the Sum method, as it is on memory
// we are only retrieving nil, but mocking it, we will be able to test that failure scenario
func Test_SumMetric_Error_MockingDao(t *testing.T) {
	assert := assert.New(t)

	dao := &m_dao.MetricDao{}
	srv := service.NewMetricsServiceImpl(dao)

	// mocking StoreMetric method
	dao.On("StoreMetric", "metric", 1).Return(nil)
	err := srv.AddMetric("metric", 1)
	assert.Nil(err)

	// mocking GetMetricSumByKey method
	dao.On("GetMetricSumByKey", "metric").Return([]model.MetricDto{}, errors.New("SOME ERROR"))
	_, err = srv.SumMetric("metric")
	assert.NotNil(err)

	dao.AssertExpectations(t)
	dao.AssertCalled(t, "StoreMetric", "metric", 1)
	dao.AssertCalled(t, "GetMetricSumByKey", "metric")
}

/*
	GetMetricSumByKey(key string) ([]model.MetricDto, error)
*/
