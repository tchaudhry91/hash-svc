package main

import (
	"flag"
	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	hashservice "github.com/tchaudhry91/hash-svc/pkg"
	"net/http"
	"os"
)

func main() {
	var serverAddr string
	flag.StringVar(
		&serverAddr,
		"serverAddr",
		":8080",
		"Fully qualified server address like 0.0.0.0:8080",
	)
	flag.Parse()

	logger := kitlog.NewJSONLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "microservices",
		Subsystem: "hash_service",
		Name:      "api_request_total",
		Help:      "Total API requests, partitioned by method and error",
	},
		fieldKeys,
	)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "microservices",
		Subsystem: "hash_service",
		Name:      "request_processing_latency",
		Help:      "Time taken per request, partitioned by method and error",
	},
		fieldKeys,
	)

	svc := hashservice.NewHashService()
	svc = hashservice.NewLoggingMiddleware(logger, svc)
	svc = hashservice.NewInstrumentingMiddleware(requestCount, requestLatency, svc)
	endpoint := hashservice.MakeHashSHA256Endpoint(svc)
	transportHandler := hashservice.MakeHashSHA256Handler(endpoint)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", transportHandler)
	logger.Log("msg", "Started HTTP Server", "addr", serverAddr)
	logger.Log("err", http.ListenAndServe(serverAddr, nil))
}
