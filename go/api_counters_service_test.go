package openapi

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCounter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT Value from Counters WHERE ID='1';").WillReturnRows(sqlmock.NewRows([]string{"Value"}).AddRow(66))
	mock.ExpectQuery("UPDATE Counters SET Value='67' WHERE ID='1';").WillReturnRows(sqlmock.NewRows([]string{"ID", "Value"}).AddRow(1, 67))

	service := CountersApiService{
		db: db,
	}

	response, err := service.GetCounter()
	assert.Nil(t, err)
	assert.IsType(t, Counter{}, response)
	assert.Equal(t, int64(66), response.(Counter).Value)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
