package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// responseWriter wraps gin.ResponseWriter to capture the response size
type responseWriter struct {
	gin.ResponseWriter
	responseSize int
}

func (w *responseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.responseSize += size
	return size, err
}

// MetricsMiddleware records metrics for each HTTP request
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Track in-flight requests
		RequestsInFlight.Inc()
		defer RequestsInFlight.Dec()

		// Wrap ResponseWriter to capture response size
		writer := &responseWriter{ResponseWriter: c.Writer}
		c.Writer = writer

		// Process request
		c.Next()

		// Record metrics after request is processed
		duration := time.Since(start).Seconds()
		status := c.Writer.Status()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		method := c.Request.Method

		// Record request duration
		HTTPRequestDuration.WithLabelValues(
			method,
			path,
		).Observe(duration)

		// Record total requests
		HTTPRequestsTotal.WithLabelValues(
			method,
			path,
			strconv.Itoa(status),
		).Inc()

		// Record failed requests (status >= 400)
		if status >= 400 {
			errorType := "client_error"
			if status >= 500 {
				errorType = "server_error"
			}
			RequestsFailed.WithLabelValues(
				method,
				path,
				errorType,
			).Inc()
		}

		// Record response size
		ResponseSize.WithLabelValues(
			method,
			path,
		).Observe(float64(writer.responseSize))
	}
}
