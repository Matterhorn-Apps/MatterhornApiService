/*
 * Matterhorn API
 *
 * Draft spec for the Matterhorn POC
 *
 * API version: 0.1.0
 * Contact: ozinoveva@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/Matterhorn-Apps/MatterhornApiService/go"
)

func main() {
	log.Printf("Server started")

	// Load environment variables
	LoadEnv()

	// Connect to database
	db, err := DbConnect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Migrate database if necessary
	Migrate(db)

	CountersApiService := openapi.NewCountersApiService(db)
	CountersApiController := openapi.NewCountersApiController(CountersApiService)

	ExerciseApiService := openapi.NewExerciseApiService(db)
	ExerciseApiController := openapi.NewExerciseApiController(ExerciseApiService)

	FoodApiService := openapi.NewFoodApiService()
	FoodApiController := openapi.NewFoodApiController(FoodApiService)

	GoalsApiService := openapi.NewGoalsApiService()
	GoalsApiController := openapi.NewGoalsApiController(GoalsApiService)

	MetricsApiService := openapi.NewMetricsApiService()
	MetricsApiController := openapi.NewMetricsApiController(MetricsApiService)

	router := openapi.NewRouter(CountersApiController, ExerciseApiController, FoodApiController, GoalsApiController, MetricsApiController)

	apiFs := http.FileServer(http.Dir("./api/"))
	router.PathPrefix("/api/").Handler(http.StripPrefix("/api/", apiFs))

	swaggerUiFs := http.FileServer(http.Dir("./swaggerui/"))
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", swaggerUiFs))

	log.Fatal(http.ListenAndServe(":5000", router))
}
