package main

import (
	"net/http"

	"github.com/devesh121/userAuth/internals/routes"
	"github.com/devesh121/userAuth/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()   // Load environment variables at first
	config.ConnectDB() // Connect to the database

	r := gin.Default()

	// Root route test
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// Health Check Endpoint
	r.GET("/health", healthCheck)

	// Register user routes
	routes.UserRoutes(r)

	println("✅ Server started at http://localhost:8080")
	r.Run(":8080")
}

func healthCheck(c *gin.Context) {
	db := config.DB
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to access database connection",
				"error":   err.Error(),
			})
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "down",
				"message": "Database connection failed",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"message": "Application and database are healthy",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "error",
		"message": "Database instance not initialized",
	})
}
