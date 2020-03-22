package main

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables defined in .env* files
// associated with the current environment
func LoadEnv() {
	env := os.Getenv("MATTERHORN_ENV")
	if "" == env {
		// Default to "development"
		env = "dev"
	}

	// Environment-specific configuration
	godotenv.Load(".env." + env)

	// Global configuration (lower priority)
	godotenv.Load()
}
