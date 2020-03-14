//Package service will have all the service and business logic definition
package service

import (
	"fmt"
	"strings"

	"github.com/mdstella/true-tickets-challenge/dao"
	"github.com/mdstella/true-tickets-challenge/errors"
)

// MetricsService is the interface that will have all the definitions for
// the service layer
//go:generate mockery -name=MetricsService
type MetricsService interface {
	AddMetric(key string, value int) error
	SumMetric(key string) (int, error)
}

// MetricsServiceImpl will be the implementation for the MetricService
type MetricsServiceImpl struct {
	dao dao.MetricDao
}

// NewMetricsServiceImpl is used as a constructor, invoked at service startup time
func NewMetricsServiceImpl(dao dao.MetricDao) MetricsService {
	return &MetricsServiceImpl{
		dao: dao,
	}
}

// AddMetric service implementation
func (srv *MetricsServiceImpl) AddMetric(key string, value int) error {
	// Input validations cosidering the keys should be empty and that the metric value should be positive
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.NewBadParamError("Errors: metric key can't be empty")
	}

	// Keeping the validation here to show that we can add validation or logic. Also the numbers validations can be done
	// using a golang validator adding parameters on the request object and using a validation library like: github.com/go-playground/validator/v10
	if value == 0 {
		return errors.NewBadParamError("Errors: metric value has to be higher or less than 0")
	}

	return srv.dao.StoreMetric(key, value)
}

// SumMetric will retrieve the sum for the given metric key
func (srv *MetricsServiceImpl) SumMetric(key string) (int, error) {
	// Input validations cosidering the keys should be empty and that the metric value should be positive
	key = strings.TrimSpace(key)
	if key == "" {
		return 0, errors.NewBadParamError("Errors: metric key can't be empty")
	}

	// Ask the DAO the metrics by the given key
	metrics, err := srv.dao.GetMetricSumByKey(key)
	if err != nil {
		return 0, err
	}

	fmt.Println("Metrics obtained by the dao: ", metrics)

	value := 0
	for _, metric := range metrics {
		value += metric.Value
	}
	return value, nil
}
