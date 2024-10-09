package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	uuidRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "uuid_requests_total",
			Help: "Total number of UUID requests",
		},
	)
)

func init() {
	prometheus.MustRegister(uuidRequests)
}

func getUUID(w http.ResponseWriter, r *http.Request) {
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
