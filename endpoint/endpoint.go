package endpoint

import (
	"context"
	"math"

	"github.com/go-kit/kit/endpoint"
	"github.com/mdstella/true-tickets-challenge/model"
	"github.com/mdstella/true-tickets-challenge/service"
)

// makeAddMetricEndpoint - endpoint invoked to store a new metric
func makeAddMetricEndpoint(srv service.MetricsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddMetricRequest)
		// rounding the input to the nearest int value
		nearestInt := int(math.Round(req.Value))
		err := srv.AddMetric(req.Key, nearestInt)
		return model.AddMetricResponse{
			Err: err,
		}, nil
	}
}

// makeSumMetricEndpoint - endpoint invoked to obtain the metric sum
func makeSumMetricEndpoint(srv service.MetricsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.SumMetricRequest)
		sum, err := srv.SumMetric(req.Key)
		return model.SumMetricResponse{
			Value: sum,
			Err:   err,
		}, nil
	}
}
