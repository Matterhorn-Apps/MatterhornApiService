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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A FoodApiController binds http requests to an api service and writes the service results to the http response
type FoodApiController struct {
	service FoodApiServicer
}

// NewFoodApiController creates a default api controller
func NewFoodApiController(s FoodApiServicer) Router {
	return &FoodApiController{ service: s }
}

// Routes returns all of the api route for the FoodApiController
func (c *FoodApiController) Routes() Routes {
	return Routes{ 
		{
			"GetFoodRecords",
			strings.ToUpper("Get"),
			"/user/{userId}/food",
			c.GetFoodRecords,
		},
		{
			"PostFoodRecord",
			strings.ToUpper("Post"),
			"/user/{userId}/food",
			c.PostFoodRecord,
		},
	}
}

// GetFoodRecords - Get food records for a user and a given time range
func (c *FoodApiController) GetFoodRecords(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	query := r.URL.Query()
	userId, err := parseIntParameter(params["userId"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	startDateTime := query.Get("startDateTime")
	endDateTime := query.Get("endDateTime")
	result, err := c.service.GetFoodRecords(userId, startDateTime, endDateTime)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// PostFoodRecord - Add a new food record
func (c *FoodApiController) PostFoodRecord(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	userId, err := parseIntParameter(params["userId"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	foodRecord := &FoodRecord{}
	if err := json.NewDecoder(r.Body).Decode(&foodRecord); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.PostFoodRecord(userId, *foodRecord)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}