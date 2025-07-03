package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateInterval(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewIntervalRepository(db)

	activityID := uuid.New()

	interval := domain.Interval{
		ActivityID: activityID,
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
			int64(interval.Duration.Seconds()),
			interval.Distance,
			string(interval.Type),
			string(interval.Stroke),
			interval.Notes,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateInterval(interval)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetIntervalsByActivity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewIntervalRepository(db)

	activityID := uuid.New()

	t.Run("success", func(t *testing.T) {
		intervals := []domain.Interval{
			{
				ID:         uuid.New(),
				ActivityID: activityID,
				Duration:   domain.DurationString((time.Minute * 30).String()),
				Distance:   1000,
				Type:       "swim",
				Stroke:     "freestyle",
				Notes:      "Test interval 1",
			},
			{
				ID:         uuid.New(),
				ActivityID: activityID,
				Duration:   domain.DurationString((time.Minute * 20).String()),
				Distance:   800,
				Type:       "swim",
				Stroke:     "backstroke",
				Notes:      "Test interval 2",
			},
		}

		rows := sqlmock.NewRows([]string{"id", "activity_id", "duration", "distance", "type", "stroke", "notes"})
		for _, interval := range intervals {
			rows.AddRow(
				interval.ID,
				interval.ActivityID,
				int64(interval.Duration.Seconds()),
				interval.Distance,
				string(interval.Type),
				string(interval.Stroke),
				interval.Notes,
			)
		}

		mock.ExpectQuery(`SELECT id, activity_id, duration, distance, type, stroke, notes FROM intervals WHERE activity_id = \$1`).
			WithArgs(activityID).
			WillReturnRows(rows)

		result, err := repo.GetIntervalsByActivity(activityID)
		assert.NoError(t, err)
		assert.Equal(t, intervals, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, activity_id, duration, distance, type, stroke, notes FROM intervals WHERE activity_id = \$1`).
			WithArgs(activityID).
			WillReturnError(assert.AnError)

		result, err := repo.GetIntervalsByActivity(activityID)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "activity_id", "duration", "distance", "type", "stroke", "notes"}).
			AddRow("invalid-uuid", activityID, 1800, 1000, "swim", "freestyle", "Test interval 1")

		mock.ExpectQuery(`SELECT id, activity_id, duration, distance, type, stroke, notes FROM intervals WHERE activity_id = \$1`).
			WithArgs(activityID).
			WillReturnRows(rows)

		result, err := repo.GetIntervalsByActivity(activityID)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
