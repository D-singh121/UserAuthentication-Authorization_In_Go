package metrics

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var pathSanitizer = []*regexp.Regexp{
	regexp.MustCompile(`/\d+`),            // replace numeric IDs
	regexp.MustCompile(`/[a-f0-9-]{24,}`), // replace MongoDB-like IDs
}

func sanitizePath(path string) string {
	for _, re := range pathSanitizer {
		path = re.ReplaceAllString(path, "/:id")
	}
	return path
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method

		RequestsInFlight.Inc()
		c.Next()
		RequestsInFlight.Dec()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())

		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		path = sanitizePath(path)

		fmt.Printf("✅ METRICS | %s %s [%s] duration=%.3fs\n", method, path, status, duration)

		HTTPRequestsTotal.WithLabelValues(method, path, status).Inc()
		HTTPRequestDuration.WithLabelValues(method, path, status).Observe(duration)

		if length := c.Writer.Size(); length >= 0 {
			ResponseSize.WithLabelValues(method, path).Observe(float64(length))
		}

		if c.Writer.Status() >= 400 {
			errorType := "client_error"
			if c.Writer.Status() >= 500 {
				errorType = "server_error"
			}
			RequestsFailed.WithLabelValues(method, path, errorType).Inc()
		}
	}
}
