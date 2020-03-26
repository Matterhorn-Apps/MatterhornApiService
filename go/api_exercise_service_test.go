package openapi

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetExerciseRecords(t *testing.T) {
	// Set up DB mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Prepare test data
	userId := int64(1)
	startDateTime := "2020-03-25 00:00:00"
	endDateTime := "2020-03-26 00:00:00"
	calories := int32(100)
	label := "exercise-label"
	timestamp := "2020-03-25 12:00:00"
	expectedExerciseRecord := ExerciseRecord{
		Calories:  calories,
		Label:     label,
		Timestamp: timestamp,
	}

	// Mock query
	query := fmt.Sprintf(
		"SELECT Calories, Label, Timestamp from ExerciseRecords WHERE UserID=%d AND Timestamp BETWEEN '%s' AND '%s';",
		userId, startDateTime, endDateTime)
	mock.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows(
			[]string{"Calories", "Label", "Timestamp"}).AddRow(calories, label, timestamp))

	// Execute function
	service := ExerciseApiService{
		db: db,
	}
	response, status, err := service.GetExerciseRecords(userId, startDateTime, endDateTime)

	// Validate result
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, *status)
	assert.IsType(t, []ExerciseRecord{}, response)
	assert.Equal(t, expectedExerciseRecord, response.([]ExerciseRecord)[0])

	// Validate mocks
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostExerciseRecord(t *testing.T) {
	// Set up DB mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Prepare test data
	userId := int64(1)
	calories := int32(100)
	label := "exercise-label"
	timestamp := "2020-03-25 12:00:00"
	exerciseRecord := ExerciseRecord{
		Calories:  calories,
		Label:     label,
		Timestamp: timestamp,
	}

	// Mock query
	query := fmt.Sprintf(
		"INSERT INTO ExerciseRecords (UserID, Calories, Label, Timestamp) VALUES (%d, %d, '%s', '%s');",
		userId, calories, label, timestamp)
	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute function
	service := ExerciseApiService{
		db: db,
	}
	response, status, err := service.PostExerciseRecord(userId, exerciseRecord)

	// Validate result
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, *status)
	assert.Equal(t, exerciseRecord, response)

	// Validate mocks
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
