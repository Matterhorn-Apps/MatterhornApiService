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

// A MetricsApiController binds http requests to an api service and writes the service results to the http response
type MetricsApiController struct {
	service MetricsApiServicer
}

// NewMetricsApiController creates a default api controller
func NewMetricsApiController(s MetricsApiServicer) Router {
	return &MetricsApiController{service: s}
}

// Routes returns all of the api route for the MetricsApiController
func (c *MetricsApiController) Routes() Routes {
	return Routes{
		{
			"GetMetrics",
			strings.ToUpper("Get"),
			"/user/{userId}/metrics",
			c.GetMetrics,
		},
		{
			"PutMetrics",
			strings.ToUpper("Put"),
			"/user/{userId}/metrics",
			c.PutMetrics,
		},
	}
}

// GetMetrics - Get the body metrics for the user
func (c *MetricsApiController) GetMetrics(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := parseIntParameter(params["userId"])
	if err != nil {
		// Bad Request
		w.WriteHeader(400)
		return
	}

	result, status, err := c.service.GetMetrics(userId)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, status, w)
}

// PutMetrics - Update body metrics for the user
func (c *MetricsApiController) PutMetrics(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := parseIntParameter(params["userId"])
	if err != nil {
		w.WriteHeader(500)
		return
	}

	bodyMetrics := &BodyMetrics{}
	if err := json.NewDecoder(r.Body).Decode(&bodyMetrics); err != nil {
		w.WriteHeader(500)
		return
	}

	result, status, err := c.service.PutMetrics(userId, *bodyMetrics)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, status, w)
}
