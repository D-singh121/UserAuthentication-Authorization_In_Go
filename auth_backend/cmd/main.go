package main

import (
	"net/http"

	"github.com/devesh121/userAuth/internals/routes"
	"github.com/devesh121/userAuth/monitoring/metrics"
	"github.com/devesh121/userAuth/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()
	routes.UserRoutes(r)

	metrics.Initialize()
	r.Use(metrics.MetricsMiddleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/health", healthCheck)

	println("✅ Server started at http://localhost:8080")
	r.Run("0.0.0.0:8080")
}

func healthCheck(c *gin.Context) {
	status := "up"
	message := "Application and database are healthy"
	isHealthy := true

	db := config.DB
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			status = "error"
			message = "Failed to access database connection"
			isHealthy = false
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  status,
				"message": message,
				"error":   err.Error(),
			})
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			status = "down"
			message = "Database connection failed"
			isHealthy = false
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  status,
				"message": message,
				"error":   err.Error(),
			})
			return
		}
	} else {
		status = "error"
		message = "Database instance not initialized"
		isHealthy = false
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  status,
			"message": message,
		})
		return
	}

	// Only record metrics if everything is healthy
	if isHealthy {
		metrics.ActiveUsers.Set(1) // Example metric
		c.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": message,
		})
	}
}
