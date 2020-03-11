package endpoint

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdstella/true-tickets-challenge/model"
)

//decodeAddMetricRequest - decode the addMetric request to generate the model request
func decodeAddMetricRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	key, ok := params["key"]
	if !ok {
		return nil, errors.New("Bad routing, metric key not provided")
	}
	var request model.AddMetricRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.Key = key
	return request, nil
}
