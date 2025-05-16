package metrics

import (
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// HTTPRequestsTotal tracks total number of HTTP requests
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests by method and path",
		},
		[]string{"method", "path", "status"},
	)

	// HTTPRequestDuration tracks request duration
	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)

	// DatabaseOperationsTotal tracks database operations
	DatabaseOperationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_operations_total",
			Help: "Total number of database operations",
		},
		[]string{"operation", "status"},
	)

	// ActiveUsers tracks current active users
	ActiveUsers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_users",
			Help: "Number of currently active users",
		},
	)

	// RequestsFailed tracks failed requests
	RequestsFailed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_failed_total",
			Help: "Total number of failed HTTP requests",
		},
		[]string{"method", "path", "error_type"},
	)

	// RequestsInFlight tracks concurrent requests
	RequestsInFlight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Current number of requests being processed",
		},
	)

	// ResponseSize tracks response sizes
	ResponseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_size_bytes",
			Help:    "Size of HTTP responses in bytes",
			Buckets: []float64{100, 1000, 10000, 100000, 1000000},
		},
		[]string{"method", "path"},
	)

	// MemoryUsage tracks memory usage
	MemoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "process_memory_bytes",
			Help: "Memory usage by type",
		},
		[]string{"type"},
	)
)

// Initialize registers all metrics with Prometheus
func Initialize() {
	// Register all metrics
	prometheus.MustRegister(HTTPRequestsTotal)
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(DatabaseOperationsTotal)
	prometheus.MustRegister(ActiveUsers)
	prometheus.MustRegister(RequestsFailed)
	prometheus.MustRegister(RequestsInFlight)
	prometheus.MustRegister(ResponseSize)
	prometheus.MustRegister(MemoryUsage)

	// Start runtime metrics collection
	go collectRuntimeMetrics()
}

// collectRuntimeMetrics periodically updates runtime metrics
func collectRuntimeMetrics() {
	var memStats runtime.MemStats
	for {
		runtime.ReadMemStats(&memStats)

		MemoryUsage.WithLabelValues("heap").Set(float64(memStats.HeapAlloc))
		MemoryUsage.WithLabelValues("stack").Set(float64(memStats.StackInuse))
		MemoryUsage.WithLabelValues("system").Set(float64(memStats.Sys))

		time.Sleep(15 * time.Second)
	}
}
