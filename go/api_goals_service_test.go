package openapi

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCalorieGoals(t *testing.T) {
	// Set up DB mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Prepare test data
	userId := int64(1)
	calories := int64(100)
	expectedCalorieGoal := CalorieGoal{
		Calories: calories,
	}

	// Mock query
	query := fmt.Sprintf("SELECT Calories from CalorieGoals WHERE UserID=%d;", userId)
	mock.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows(
			[]string{"Calories"}).AddRow(calories))

	// Execute function
	service := GoalsApiService{
		db: db,
	}
	response, status, err := service.GetCalorieGoal(userId)

	// Validate result
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, *status)
	assert.Equal(t, expectedCalorieGoal, response)

	// Validate mocks
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPutCalorieGoal(t *testing.T) {
	// Set up DB mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Prepare test data
	userId := int64(1)
	calories := int64(100)
	calorieGoal := CalorieGoal{
		Calories: calories,
	}

	// Mock query
	query := fmt.Sprintf(
		"INSERT INTO CalorieGoals (UserID, Calories) VALUES (%d, %d) ON DUPLICATE KEY UPDATE Calories = %[2]d;",
		userId, calories)
	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute function
	service := GoalsApiService{
		db: db,
	}
	response, status, err := service.PutCalorieGoal(userId, calorieGoal)

	// Validate result
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNoContent, *status)
	assert.Equal(t, calorieGoal, response)

	// Validate mocks
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
