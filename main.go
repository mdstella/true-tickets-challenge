// Package classification true-tickets challenge
//
// the purpose of this application is create metrics system
//
//     Schemes: http
//     Host: localhost:9091
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/mdstella/true-tickets-challenge/dao"
	"github.com/mdstella/true-tickets-challenge/service"

	"github.com/go-kit/kit/log"
	"github.com/mdstella/true-tickets-challenge/endpoint"
)

//go:generate swagger generate spec -o swagger-ui/swagger.json
func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	port := os.Getenv("PORT")

	if port == "" {
		logger.Log("$PORT must be set")
		port = ":9091"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	// service, dao and endpoints intialization
	dao := dao.NewMetricDaoMemoryImpl()
	srv := service.NewMetricsServiceImpl(dao)
	handler := endpoint.RegisterEndpoints(srv)

	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, handler))

}
