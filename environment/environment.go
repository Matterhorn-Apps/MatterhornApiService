package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables defined in .env* files
// associated with the current environment
func LoadEnv(relPath string) {
	env := os.Getenv("MATTERHORN_ENV")
	if "" == env {
		// Default to "local"
		env = "dev"
	}

	// Environment-specific configuration
	err := godotenv.Load(relPath + "/.env." + env)
	if err != nil {
		log.Fatalf("Failed to load configuration from .env.%s: %v", env, err)
	}

	// Global configuration (lower priority)
	err = godotenv.Load(relPath + "/.env")
	if err != nil {
		log.Fatalf("Failed to load configuration from .env: %v", err)
	}
}
