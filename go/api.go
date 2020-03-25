/*
 * Matterhorn API
 *
 * Draft spec for the Matterhorn POC
 *
 * API version: 0.1.0
 * Contact: ozinoveva@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
)

// CountersApiRouter defines the required methods for binding the api requests to a responses for the CountersApi
// The CountersApiRouter implementation should parse necessary information from the http request,
// pass the data to a CountersApiServicer to perform the required actions, then write the service results to the http response.
type CountersApiRouter interface {
	GetCounter(http.ResponseWriter, *http.Request)
}

// ExerciseApiRouter defines the required methods for binding the api requests to a responses for the ExerciseApi
// The ExerciseApiRouter implementation should parse necessary information from the http request,
// pass the data to a ExerciseApiServicer to perform the required actions, then write the service results to the http response.
type ExerciseApiRouter interface {
	GetExerciseRecords(http.ResponseWriter, *http.Request)
	PostExerciseRecord(http.ResponseWriter, *http.Request)
}

// FoodApiRouter defines the required methods for binding the api requests to a responses for the FoodApi
// The FoodApiRouter implementation should parse necessary information from the http request,
// pass the data to a FoodApiServicer to perform the required actions, then write the service results to the http response.
type FoodApiRouter interface {
	GetFoodRecords(http.ResponseWriter, *http.Request)
	PostFoodRecord(http.ResponseWriter, *http.Request)
}

// GoalsApiRouter defines the required methods for binding the api requests to a responses for the GoalsApi
// The GoalsApiRouter implementation should parse necessary information from the http request,
// pass the data to a GoalsApiServicer to perform the required actions, then write the service results to the http response.
type GoalsApiRouter interface {
	GetCalorieGoal(http.ResponseWriter, *http.Request)
	PutCalorieGoal(http.ResponseWriter, *http.Request)
}

// MetricsApiRouter defines the required methods for binding the api requests to a responses for the MetricsApi
// The MetricsApiRouter implementation should parse necessary information from the http request,
// pass the data to a MetricsApiServicer to perform the required actions, then write the service results to the http response.
type MetricsApiRouter interface {
	GetMetrics(http.ResponseWriter, *http.Request)
	PutMetrics(http.ResponseWriter, *http.Request)
}

// CountersApiServicer defines the api actions for the CountersApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type CountersApiServicer interface {
	GetCounter() (interface{}, error)
}

// ExerciseApiServicer defines the api actions for the ExerciseApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ExerciseApiServicer interface {
	GetExerciseRecords(int64, string, string) (interface{}, *int, error)
	PostExerciseRecord(int64, ExerciseRecord) (interface{}, *int, error)
}

// FoodApiServicer defines the api actions for the FoodApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type FoodApiServicer interface {
	GetFoodRecords(int64, string, string) (interface{}, *int, error)
	PostFoodRecord(int64, FoodRecord) (interface{}, *int, error)
}

// GoalsApiServicer defines the api actions for the GoalsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type GoalsApiServicer interface {
	GetCalorieGoal(int64) (interface{}, *int, error)
	PutCalorieGoal(int64, CalorieGoal) (interface{}, *int, error)
}

// MetricsApiServicer defines the api actions for the MetricsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type MetricsApiServicer interface {
	GetMetrics(int64) (interface{}, *int, error)
	PutMetrics(int64, BodyMetrics) (interface{}, *int, error)
}
