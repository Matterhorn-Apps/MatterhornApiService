package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"database/sql"
)

// Defines port to listen on for requests
const port = ":5000"

// Responds to an HTTP request with a friendly response message
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

type counterResponse struct {
	ID    int
	Value int
}

// Responds to an HTTP request by displaying the latest counter value and incrementing it in the database
func counterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Defensive check of DB instance passed to function
	if db == nil {
		fmt.Fprintf(w, "Db is nil :(")
		return
	}

	// Query the database for the current counter value
	readRows, readRrr := db.Query("SELECT Value from Counters WHERE ID='1';")
	if readRrr != nil {
		panic("Failed to query database. Error: " + readRrr.Error())
	}
	defer readRows.Close()

	// Read value from row response
	var value int
	readRows.Next()
	readRows.Scan(&value)

	// Query the database to update the counter value
	updateRows, updateErr := db.Query(fmt.Sprintf("UPDATE Counters SET Value='%d' WHERE ID='%d'", value+1, 1))
	if updateErr != nil {
		panic("Failed to update counter value. Error: " + updateErr.Error())
	}
	defer updateRows.Close()

	// Construct response data object
	var data counterResponse = counterResponse{
		ID:    1,
		Value: value,
	}
	jData, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		panic(fmt.Sprintf("Failed to marshal json response data: %s", jsonErr.Error()))
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

// Main entry point function for MatterhornApiService
func main() {
	run()
	os.Exit(1)
}

func run() {
	db, err := DbConnect()
	if err != nil {
		panic("Failed to connect to DB! Error: " + err.Error())
	}

	if db == nil {
		panic("DB is nil immediately after assignment!")
	}

	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { helloWorld(w, r) })
	http.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) { counterHandler(w, r, db) })
	http.ListenAndServe(port, nil)
}
