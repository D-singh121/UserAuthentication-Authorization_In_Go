// Environment Loader
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig struct holds all necessary database environment variables
type DBConfig struct {
	Host     string // Database host, e.g., localhost
	Port     string // Port where PostgreSQL is running (default: 5432)
	User     string // Username for PostgreSQL
	Password string // Password for PostgreSQL
	DB_Name  string // Database name (was DB_Name)
	SSLMode  string // SSL mode setting, e.g., disable
}

// LoadEnv loads the .env file into the application
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, relying on system env variables")
	}
}

// GetDBConfig returns a populated DBConfig struct using values from the environment
func GetDBConfig() (DBConfig, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	if host == "" || port == "" || user == "" || password == "" || dbName == "" {
		return DBConfig{}, fmt.Errorf("one or more critical database environment variables are not set (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)")
	}

	return DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DB_Name:  dbName,
		SSLMode:  sslMode,
	}, nil

}
