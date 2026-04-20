package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "dopc_requests_total",
			Help: "Total number of requests to the delivery price endpoint",
		},
		[]string{"status_code"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "dopc_request_duration_seconds",
			Help:    "Duration of requests to the delivery price endpoint",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"status_code"},
	)
)

func Register() {
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(RequestDuration)
}
