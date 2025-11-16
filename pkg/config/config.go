package config

import (
	"os"
)

// Config holds the configuration for the application.
type Config struct {
	APIKey string
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		APIKey: getEnv("FRED_API_KEY", ""),
	}
}

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
