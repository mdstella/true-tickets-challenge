package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mdstella/true-tickets-challenge/model"
	"github.com/mdstella/true-tickets-challenge/service"
)

// makeAddMetricEndpoint - endpoint invoked to store a new metric
func makeAddMetricEndpoint(srv service.MetricsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddMetricRequest)
		err := srv.AddMetric(req.Key, req.Value)
		return model.AddMetricResponse{
			Err: err,
		}, nil
	}
}
