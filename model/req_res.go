package model

// In this file we will keep the requests and response objects for the API

// AddMetric - Request / Response

// AddMetricRequest is the request that will store a new metric
// swagger:parameters AddMetricRequest
type AddMetricRequest struct {
	// The metric key
	// required: true
	// in: path
	Key string `json:"key"`

	// The metric value
	// in: body
	Value int `json:"value,omitempty"`
}

// AddMetricResponse just notify if there is any error adding the new metric
// swagger:response AddMetricResponse
type AddMetricResponse struct {
	//swagger:ignore
	Err error `json:"error,omitempty"`
}

// Implementing error method
func (r AddMetricResponse) Error() error { return r.Err }

// SumMetric - Request / Response

// SumMetricRequest is the request that will ask for a given metric sum
// swagger:parameters SumMetricRequest
type SumMetricRequest struct {
	// The metric key
	// required: true
	// in: path
	Key string `json:"key"`
}

// SumMetricResponse retrieves the sum of a given metric in the last 60 minutes
// swagger:response SumMetricResponse
type SumMetricResponse struct {
	// reprensents the sum of the metric
	Value int `json:"value"`
	//swagger:ignore
	Err error `json:"error,omitempty"`
}

// Implementing error method
func (r SumMetricResponse) Error() error { return r.Err }
