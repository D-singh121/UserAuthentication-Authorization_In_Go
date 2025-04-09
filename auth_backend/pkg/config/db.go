// Database Connection Handler
package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devesh121/userAuth/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB // Global DB instance accessible across the app

// ConnectDB initializes and opens a DB connection using the configuration from .env
func ConnectDB() {

	LoadEnv() // Load environment variables from .env file

	dbConfig, err := GetDBConfig() // Get database config values
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	// Create the PostgreSQL DSN (Data Source Name) string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DB_Name, dbConfig.SSLMode)

	// Configure GORM with connection pooling and logging (for development/debugging)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level (adjust for production)
			Colorful:      true,        // Disable color in production
		},
	)
	// Open a new connection using GORM and the PostgreSQL driver
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true, // Recommended for performance in some cases
		PrepareStmt:                              true, // Improves performance for repeated queries
		DisableForeignKeyConstraintWhenMigrating: true, // Consider your needs for FK constraints in production
		// Uncomment and adjust these for production connection pooling
		// ConnPool: &sql.Pool{
		// 	MaxIdleConns: 10,
		// 	MaxOpenConns: 100,
		// 	ConnMaxLifetime: time.Hour,
		// },
	})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err) // Exit if connection fails
	}

	log.Println("✅ Database connection successful")

	//Auto migrating the models for table creation on psql database
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Failed to auto migrate models: %v", err)
	}
	log.Println("✅ Database migration completed")
}
