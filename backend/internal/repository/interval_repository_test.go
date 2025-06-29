package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestPostgresIntervalRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewIntervalRepository(db)

	activityID := uuid.New()
	start := time.Now()

	interval := domain.Interval{
		ActivityID: activityID,
		StartTime:  start,
		Duration:   domain.DurationString((time.Minute * 30).String()),
		Distance:   1000,
		Type:       "swim",
		Stroke:     "freestyle",
		Notes:      "Test interval",
	}

	mock.ExpectExec(`INSERT INTO intervals`).
		WithArgs(
			sqlmock.AnyArg(), // id (generated UUID)
			interval.ActivityID,
			interval.StartTime,
			int64(interval.Duration.Seconds()),
			interval.Distance,
			string(interval.Type),
			string(interval.Stroke),
			interval.Notes,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(interval)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
