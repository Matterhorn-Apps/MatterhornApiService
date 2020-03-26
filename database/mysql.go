package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationsDir string = "./migrations"

// DbConnect establishes a connection with MatterhornDb and returns a sql DB instance.
// On failure, logs error message and immediately exits program.
func DbConnect() (*sql.DB, error) {
	// Attempt to read DB connection information from environment variables
	var dbEndpoint, dbName, dbPassword, dbUser string
	var keyExists bool
	dbEndpoint, keyExists = os.LookupEnv("MATTERHORN_DB_ENDPOINT")
	if !keyExists {
		log.Fatalf("MATTERHORN_DB_ENDPOINT environment variable is unset!")
	}

	dbUser, keyExists = os.LookupEnv("MATTERHORN_DB_USERNAME")
	if !keyExists {
		log.Fatalf("MATTERHORN_DB_USERNAME environment variable is unset!")
	}

	dbPassword, keyExists = os.LookupEnv("MATTERHORN_DB_PASSWORD")
	if !keyExists {
		log.Fatalf("MATTERHORN_DB_PASSWORD environment variable is unset!")
	}

	dbName, keyExists = os.LookupEnv("MATTERHORN_DB_NAME")
	if !keyExists {
		log.Fatalf("MATTERHORN_DB_NAME environment variable is unset!")
	}

	// Create the MySQL DNS string for the DB connection
	// user:password@protocol(endpoint)/dbname?<params>
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true", dbUser, dbPassword, dbEndpoint, dbName)

	// Open db connection
	return sql.Open("mysql", dnsStr)
}

// Migrate executed migrations on the given database.
// On failure, logs error message and immediately exits program.
func Migrate(db *sql.DB) {
	log.Println("Migrating database...")

	// Run migrations
	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatalf("Could not start database migration: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsDir), // file://path/to/directory
		"mysql", driver)

	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database: %v", err)
	}

	log.Println("Database migrated.")
}

func TryExtractMySQLErrorCode(err error) (*uint16, bool) {
	// Attempt to map specific MySQL error to a status code
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		return &mysqlErr.Number, true
	} else {
		return nil, false
	}
}
