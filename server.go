package main

import (
	"fmt"
	"log"
	"net/http"
)

// Defines port to listen on for requests
const port = ":5000"

// Responds to an HTTP request with a friendly response message
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// Main entry point function for MatterhornApiService
func main() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(port, nil))
}
