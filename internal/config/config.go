package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBDSN     string
	JWTSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbDSN := os.Getenv("DATABASE_URL")
	if dbDSN == "" {
		// Default local development postgres connection
		dbDSN = "host=localhost user=postgres password=postgres dbname=scd_event port=5432 sslmode=disable TimeZone=UTC"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super-secret-key-change-in-production"
	}

	return &Config{
		Port:      port,
		DBDSN:     dbDSN,
		JWTSecret: jwtSecret,
	}
}
