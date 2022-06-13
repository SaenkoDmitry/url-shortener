package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	responseTimeMetric = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "url_shortener_response_time",
		Help:    "Average execution",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "status"})
	rpsMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "url_shortener_requests_per_seconds",
		Help: "Total amount of processed requests",
	}, []string{"method", "status"})
)

func HandleResponseTime(method, status string, start time.Time) {
	responseTimeMetric.WithLabelValues(method, status).Observe(float64(time.Since(start).Milliseconds()))
}

func HandleRPS(method, status string) {
	rpsMetric.WithLabelValues(method, status).Inc()
}
