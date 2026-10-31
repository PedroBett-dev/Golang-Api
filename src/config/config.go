package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("erro ao carregar .env: %w", err)
	}

	cfg := &Config{
		ApiConfig: getApiConfig(),
	}

	return cfg, nil
}

func getApiConfig() apiConfig {
	return apiConfig{
		url:  getEnv("API_URL", "localhost"),
		port: getEnv("API_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
