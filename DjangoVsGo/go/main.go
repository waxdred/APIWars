package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	uuidRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "uuid_requests_total",
			Help: "Total number of UUID requests",
		},
	)
	uuidLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_golang_request_duration_seconds",
			Help:    "Histogram for the duration in seconds of UUID requests.",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	prometheus.MustRegister(uuidRequests)
	prometheus.MustRegister(uuidLatency)
}

func getUUID(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		uuidLatency.Observe(duration)
	}()

	uuidRequests.Inc()
	uuid := uuid.New()
	json.NewEncoder(w).Encode(uuid)
}

func main() {
	http.HandleFunc("/uuid", getUUID)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
