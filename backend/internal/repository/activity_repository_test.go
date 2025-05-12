package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	activity := domain.Activity{
		UserID:       1,
		Start:        time.Now(),
		Duration:     30 * time.Minute,
		Distance:     1000,
		Laps:         20,
		PoolSize:     50,
		LocationType: domain.LocationType("indoor"),
	}

	mock.ExpectExec(`INSERT INTO activities`).
		WithArgs(
			activity.UserID,
			activity.Start,
			int64(activity.Duration.Seconds()),
			activity.Distance,
			activity.Laps,
			activity.PoolSize,
			string(activity.LocationType),
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(activity)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllActivities(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type"}).
			AddRow(1, 1, time.Now(), int64(1800), 1000, 20, 50, "indoor").
			AddRow(2, 2, time.Now(), int64(3600), 2000, 40, 50, "outdoor")

		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type FROM activities`).
			WillReturnRows(rows)

		activities, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, activities, 2)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type FROM activities`).
			WillReturnError(assert.AnError)

		activities, err := repo.GetAll()
		assert.Error(t, err)
		assert.Nil(t, activities)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type"}).
			AddRow("invalid_id", 1, time.Now(), int64(1800), 1000, 20, 50, "indoor")

		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type FROM activities`).
			WillReturnRows(rows)

		activities, err := repo.GetAll()
		assert.Error(t, err)
		assert.Nil(t, activities)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAllActivitiesByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	t.Run("success", func(t *testing.T) {
		userID := 1
		rows := sqlmock.NewRows([]string{"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type"}).
			AddRow(1, userID, time.Now(), int64(1800), 1000, 20, 50, "indoor")

		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type FROM activities WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnRows(rows)

		activities, err := repo.GetAllByUser(userID)
		assert.NoError(t, err)
		assert.Len(t, activities, 1)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		userID := 1
		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type FROM activities WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnError(assert.AnError)

		activities, err := repo.GetAllByUser(userID)
		assert.Error(t, err)
		assert.Nil(t, activities)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		userID := 1
		rows := sqlmock.NewRows([]string{"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type"}).
			AddRow("invalid_id", userID, time.Now(), int64(1800), 1000, 20, 50, "indoor")

		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type FROM activities WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnRows(rows)

		activities, err := repo.GetAllByUser(userID)
		assert.Error(t, err)
		assert.Nil(t, activities)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
