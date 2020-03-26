package database

import (
	"testing"

	"github.com/Matterhorn-Apps/MatterhornApiService/environment"
)

func TestDbConnect(t *testing.T) {
	// Load environment variables
	environment.LoadEnv("../")

	// Attempt to connect to the database
	db, err := DbConnect()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}

	// Attempt to ping the database to make sure we actually established a connection
	err = db.Ping()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}
}
