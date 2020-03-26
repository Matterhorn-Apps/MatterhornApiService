package openapi

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetMetrics(t *testing.T) {
	// Set up DB mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Prepare test data
	userId := int64(1)
	height := int64(70)
	weight := float32(160.2)
	age := int32(30)
	sex := "Male"
	expectedBodyMetrics := BodyMetrics{
		Height: height,
		Weight: weight,
		Age:    age,
		Sex:    sex,
	}

	// Mock query
	query := fmt.Sprintf("SELECT Height, Weight, Sex, Age from Users WHERE UserID=%d;", userId)
	mock.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows(
			[]string{"Height", "Weight", "Sex", "Age"}).AddRow(height, weight, sex, age))

	// Execute function
	service := MetricsApiService{
		db: db,
	}
	response, status, err := service.GetMetrics(userId)

	// Validate result
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, *status)
	assert.Equal(t, expectedBodyMetrics, response)

	// Validate mocks
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPutMetrics(t *testing.T) {
	// Set up DB mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Prepare test data
	userId := int64(1)
	height := int64(70)
	weight := float32(160.2)
	age := int32(30)
	sex := "Male"
	bodyMetrics := BodyMetrics{
		Height: height,
		Weight: weight,
		Age:    age,
		Sex:    sex,
	}

	// Mock query
	query := fmt.Sprintf(
		"UPDATE Users SET Height = %[2]d, Weight = %[3]f, Sex = '%[4]s', Age = %[5]d WHERE UserID = %[1]d",
		userId, height, weight, sex, age)
	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute function
	service := MetricsApiService{
		db: db,
	}
	response, status, err := service.PutMetrics(userId, bodyMetrics)

	// Validate result
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNoContent, *status)
	assert.Equal(t, bodyMetrics, response)

	// Validate mocks
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
