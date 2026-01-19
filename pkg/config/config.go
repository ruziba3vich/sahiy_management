package config

import (
	"os"
	"strconv"

	"github.com/ruziba3vich/sahiy_management/pkg/database"
)

type Config struct {
	database.DBConfig
	JWTSecret      string
	JWTExpiryHours int
	AppPort        string
}

func LoadConfig() *Config {
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	jwtExpiryHours, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))

	dbConfig := database.DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     dbPort,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "sahiy_management"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	jwtSecret := getEnv("JWT_SECRET", "aG4gRG9lIiwiYWRtaW4iOnRydW###UsImlhdCI6MTUxNjIzOTAyMn0")
	port := getEnv("PORT", "8080")

	return &Config{
		JWTSecret:      jwtSecret,
		DBConfig:       dbConfig,
		JWTExpiryHours: jwtExpiryHours,
		AppPort:        port,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
