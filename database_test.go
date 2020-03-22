package main

import (
	"testing"
)

func TestDbConnect(t *testing.T) {
	// Load environment variables
	LoadEnv()

	// Attempt to connect to the database
	db, err := DbConnect()
	if err != nil {
		t.Errorf("Failed to connect to database: %s", err.Error())
	}

	// Attempt to ping the database to make sure we actually established a connection
	err = db.Ping()
	if err != nil {
		t.Errorf("Failed to connect to database: %s", err.Error())
	}
}
