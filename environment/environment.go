package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables defined in .env* files
// associated with the current environment
func LoadEnv(relPath string) {
	stage, ok := os.LookupEnv("MATTERHORN_ENV")
	if !ok {
		log.Fatalf("MATTERHORN_ENV environment variable must be set!")
	}

	switch stage {
	case "dev":
	case "prod":
		break
	default:
		log.Fatalf("Invalid value defined for MATTERHORN_ENV: %s", stage)
	}

	// Environment configuration files are loaded in order from HIGHEST to LOWEST priority.
	// .env.STAGE.local > .env.STAGE > .env.local > .env
	var err error
	err = godotenv.Load(fmt.Sprintf("%s/.env.%s.local", relPath, stage))
	if err != nil {
		log.Printf("Local stage configuration overrides not found.")
	}

	err = godotenv.Load(fmt.Sprintf("%s/.env.%s", relPath, stage))
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(fmt.Sprintf("%s/.env.local", relPath))
	if err != nil {
		log.Printf("Local global configuration overrides not found.")
	}

	err = godotenv.Load(fmt.Sprintf("%s/.env", relPath))
	if err != nil {
		log.Fatal(err)
	}
}
