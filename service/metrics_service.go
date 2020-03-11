//Package service will have all the service and business logic definition
package service

import (
	"strings"

	"github.com/mdstella/true-tickets-challenge/dao"
	"github.com/mdstella/true-tickets-challenge/errors"
)

// MetricsService is the interface that will have all the definitions for
// the service layer
//go:generate mockery -name=MetricsService
type MetricsService interface {
	AddMetric(key string, value int) error
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
	if value <= 0 {
		return errors.NewBadParamError("Errors: metric value has to be higher than 0")
	}

	return srv.dao.StoreMetric(key, value)
}
