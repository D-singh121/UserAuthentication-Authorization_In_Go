package metrics

import (
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests by method and path",
		},
		[]string{"method", "path", "status"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path"},
	)

	DatabaseOperationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_operations_total",
			Help: "Total number of database operations",
		},
		[]string{"operation", "status"},
	)

	ActiveUsers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_users",
			Help: "Number of currently active users",
		},
	)

	RequestsFailed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_failed_total",
			Help: "Total number of failed HTTP requests",
		},
		[]string{"method", "path", "error_type"},
	)

	RequestsInFlight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Current number of requests being processed",
		},
	)

	ResponseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_size_bytes",
			Help:    "Size of HTTP responses in bytes",
			Buckets: []float64{100, 1000, 10000, 100000, 1000000}, // 100B, 1KB, 10KB, 100KB, 1MB
		},
		[]string{"method", "path"},
	)

	MemoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "process_memory_bytes",
			Help: "Memory usage by type",
		},
		[]string{"type"},
	)
)

func Initialize() {
	prometheus.MustRegister(
		HTTPRequestsTotal,
		HTTPRequestDuration,
		DatabaseOperationsTotal,
		ActiveUsers,
		RequestsFailed,
		RequestsInFlight,
		ResponseSize,
		MemoryUsage,
	)

	go collectRuntimeMetrics()
}

// collectRuntimeMetrics periodically updates runtime metrics
func collectRuntimeMetrics() {
	var memStats runtime.MemStats
	for {
		runtime.ReadMemStats(&memStats)

		// Update memory metrics
		MemoryUsage.WithLabelValues("heap").Set(float64(memStats.HeapAlloc))
		MemoryUsage.WithLabelValues("stack").Set(float64(memStats.StackInuse))
		MemoryUsage.WithLabelValues("system").Set(float64(memStats.Sys))

		// Sleep for a short duration before next collection
		time.Sleep(15 * time.Second)
	}
}
