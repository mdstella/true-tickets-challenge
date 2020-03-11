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
