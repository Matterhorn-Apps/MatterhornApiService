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

type FoodRecord struct {

	// Calories in this food record
	Calories int32 `json:"calories"`

	// Optional name or label for this food
	Label string `json:"label,omitempty"`

	Timestamp string `json:"timestamp"`
}
