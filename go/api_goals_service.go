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
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	database "github.com/Matterhorn-Apps/MatterhornApiService/database"
)

// GoalsApiService is a service that implents the logic for the GoalsApiServicer
// This service should implement the business logic for every endpoint for the GoalsApi API.
// Include any external packages or services that will be required by this service.
type GoalsApiService struct {
	db *sql.DB
}

// NewGoalsApiService creates a default api service
func NewGoalsApiService(db *sql.DB) GoalsApiServicer {
	return &GoalsApiService{
		db: db,
	}
}

// GetCalorieGoal - Get the calorie goal for the user
func (s *GoalsApiService) GetCalorieGoal(userId int64) (interface{}, *int, error) {
	db := s.db

	// Query the database for matching exercise records
	query := fmt.Sprintf("SELECT Calories from CalorieGoals WHERE UserID=%d;", userId)
	readRows, readErr := db.Query(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User not found
				status := http.StatusNotFound
				return nil, &status, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, nil, readErr
	}
	defer readRows.Close()

	var calories int64
	readRows.Next()
	readErr = readRows.Scan(&calories)
	if readErr != nil {
		log.Printf("Failed to read row returned from query: %v", readErr)
		return nil, nil, readErr
	}

	record := CalorieGoal{
		Calories: calories,
	}

	status := http.StatusOK
	return record, &status, nil
}

// PutCalorieGoal - Update the calorie goal for the user
func (s *GoalsApiService) PutCalorieGoal(userId int64, calorieGoal CalorieGoal) (interface{}, *int, error) {
	db := s.db

	if calorieGoal.Calories < 0 {
		status := http.StatusBadRequest
		return nil, &status, errors.New("Invalid calorie goal provided")
	}

	// Query the database for matching exercise records
	query := fmt.Sprintf(
		"INSERT INTO CalorieGoals (UserID, Calories) VALUES (%d, %d) ON DUPLICATE KEY UPDATE Calories = %[2]d;",
		userId, calorieGoal.Calories)
	_, readErr := db.Exec(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User not found
				status := http.StatusNotFound
				return nil, &status, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		status := http.StatusNoContent
		return nil, &status, readErr
	}

	status := http.StatusNoContent
	return calorieGoal, &status, nil
}
