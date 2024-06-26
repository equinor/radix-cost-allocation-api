package router

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/urfave/negroni/v3"
)

// NewMetricsHandler Constructor function
func NewMetricsHandler() http.Handler {
	serveMux := http.NewServeMux()
	serveMux.Handle("GET /metrics", promhttp.Handler())

	rec := negroni.NewRecovery()
	rec.PrintStack = false
	n := negroni.New(
		rec,
		NewZerologHandler(log.Logger),
	)
	n.UseHandler(serveMux)

	return n
}
