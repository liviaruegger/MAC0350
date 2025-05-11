package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestPostgresIntervalRepository_Create(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewIntervalRepository(db)

	// Define the interval to be inserted
	interval := domain.Interval{
		ActivityID: 1,
		StartTime:  time.Now(),
		Duration:   time.Minute * 30,
		Distance:   1000,
		Type:       "swim",
		Stroke:     "freestyle",
		Notes:      "Test interval",
	}

	// Expect the SQL query to be executed with the correct parameters
	mock.ExpectExec(`INSERT INTO intervals`).
		WithArgs(
			interval.ActivityID,
			interval.StartTime,
			int64(interval.Duration.Seconds()),
			interval.Distance,
			string(interval.Type),
			string(interval.Stroke),
			interval.Notes,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the Create method
	err = repo.Create(interval)

	// Assert no errors occurred
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
