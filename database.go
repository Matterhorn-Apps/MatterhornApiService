package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/go-sql-driver/mysql"
)

// DbConnect establishes a connection with MatterhornDb and returns a sql DB instance.
func DbConnect() (*sql.DB, error) {
	// Attempt to read DB connection information from environment variables
	var dbEndpoint, dbPassword, dbName string
	var keyExists bool
	dbEndpoint, keyExists = os.LookupEnv("MATTERHORN_DB_ENDPOINT")
	if !keyExists {
		panic("MATTERHORN_DB_ENDPOINT environment variable is unset!")
	}

	dbPassword, keyExists = os.LookupEnv("MATTERHORN_DB_PASSWORD")
	if !keyExists {
		panic("MATTERHORN_DB_PASSWORD environment variable is unset!")
	}

	dbName, keyExists = os.LookupEnv("MATTERHORN_DB_NAME")
	if !keyExists {
		panic("MATTERHORN_DB_NAME environment variable is unset!")
	}

	dbUser := "admin"

	// Create the MySQL DNS string for the DB connection
	// user:password@protocol(endpoint)/dbname?<params>
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbEndpoint, dbName)

	// Open db connection
	return sql.Open("mysql", dnsStr)
}
