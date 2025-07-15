package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ListeningPort    string
	ExpectedPassword string
	ExpectedUser     string
}

// For inject using ldflags in compile time
var (
	ListeningPort    string
	ExpectedPassword string
	ExpectedUser     string
)

// LoadConfig loads configuration values from ldflags (in prod) or env vars (in dev).
func LoadConfig() (*Config, error) {
	godotenv.Load()
	cfg := &Config{
		ListeningPort:    os.Getenv("LISTENING_PORT"),
		ExpectedPassword: os.Getenv("EXPECTED_PASSWORD"),
		ExpectedUser:     os.Getenv("EXPECTED_USER"),
	}

	// Fallback/default port
	if cfg.ListeningPort == "" {
		cfg.ListeningPort = "8080"
	}

	return cfg, nil
}
