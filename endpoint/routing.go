// Package endpoint will have all the routing definition for the
// endpoints defined on our API
package endpoint

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/mdstella/true-tickets-challenge/errors"
	"github.com/mdstella/true-tickets-challenge/service"
)

// RegisterEndpoints will register all the endpoints from our API asociating them with the
// service layer obtained by parameter
func RegisterEndpoints(srv service.MetricsService) http.Handler {
	r := mux.NewRouter()

	// We generate the handler functions that will be executed once each endpoint is invoked. It will use a go-kit endpoint
	// (the one used by the framework selected to create the server), a decoder that will decode the message sent to the endpoint
	// into the model, and a common response encoder
	addMetricHandler := httptransport.NewServer(makeAddMetricEndpoint(srv), decodeAddMetricRequest, encodeResponse)
	// Define the method, path and the handler in the router to be able to dispatch requests to it.
	// swagger:route POST /metric/{key} metrics AddMetricRequest
	//
	// Status endpoint to check the app is up and running
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: StatusResponse
	r.Methods("POST").Path("/metric/{key}").Handler(addMetricHandler)

	// adding swagger endpoint to have the API doc available
	swaggerUrl := "/swagger-ui/"
	r.PathPrefix(swaggerUrl).Handler(http.StripPrefix(swaggerUrl, http.FileServer(http.Dir("./swagger-ui/"))))

	return cors.Default().Handler(r)
}

// errorer is implemented by all concrete response types that may contain errors.
type errorer interface {
	Error() error
}

// encodeResponse will encode the response message that will be sent to the client
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if e, ok := response.(errorer); ok && e.Error() != nil {
		EncodeHttpError(ctx, e.Error(), w)
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

// HttpCodeFrom retrieve the http status code, if it's not a core error will be consider an internal server error
func HttpCodeFrom(err error) int {
	if e, ok := err.(errors.CoreError); ok {
		return e.Code
	} else if e, ok := err.(*errors.CoreError); ok {
		return e.Code
	}
	return http.StatusInternalServerError
}

// EncodeHttpError encode the error response
func EncodeHttpError(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	httpStatus := HttpCodeFrom(err)
	w.WriteHeader(httpStatus)

	// gokit adds the prefix "Encode:" when encoding a msg and "Decode: " when error in the decoding phase
	message := strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(err.Error(), "Encode: "), "Decode: "), "Do: ")
	var errorResponse *errors.CoreError
	if nerr := errors.GetCoreError(err); nerr != nil {
		errorResponse = nerr
	} else {
		errorResponse = &errors.CoreError{Message: message}
	}

	json.NewEncoder(w).Encode(errorResponse)
}
