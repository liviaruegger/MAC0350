package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	t.Run("success", func(t *testing.T) {
		activity := domain.Activity{
			ID:           uuid.New(),
			UserID:       uuid.New(),
			Start:        time.Now(),
			Duration:     domain.DurationString((30 * time.Minute).String()),
			Distance:     1000,
			Laps:         20,
			PoolSize:     50,
			LocationType: domain.LocationType(domain.LocationPool),
			Notes:        "Test notes",
		}

		mock.ExpectExec(`INSERT INTO activities`).
			WithArgs(
				activity.ID,
				activity.UserID,
				activity.Start,
				int64(activity.Duration.Seconds()),
				activity.Distance,
				activity.Laps,
				activity.PoolSize,
				string(activity.LocationType),
				activity.Notes,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.CreateActivity(activity)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAllActivities(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	t.Run("success", func(t *testing.T) {
		id1 := uuid.New()
		userID1 := uuid.New()
		id2 := uuid.New()
		userID2 := uuid.New()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type", "notes"}).
			AddRow(id1, userID1, now, int64(1800), 1000, 20, 50, "pool", "Note A").
			AddRow(id2, userID2, now, int64(3600), 2000, 40, 25, "open_water", "Note B")

		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities`).
			WillReturnRows(rows)

		activities, err := repo.GetAllActivities()
		assert.NoError(t, err)
		assert.Len(t, activities, 2)
		assert.Equal(t, "Note A", activities[0].Notes)
		assert.Equal(t, "Note B", activities[1].Notes)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities`).
			WillReturnError(assert.AnError)

		activities, err := repo.GetAllActivities()
		assert.Error(t, err)
		assert.Empty(t, activities)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities`).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type", "notes"}).
				AddRow("invalid_uuid", "another_invalid_uuid", time.Now(), int64(1800), 1000, 20, 50, "pool", "Note"))

		activities, err := repo.GetAllActivities()
		assert.Error(t, err)
		assert.Empty(t, activities)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetActivitiesByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	t.Run("success", func(t *testing.T) {
		userID := uuid.New()
		id := uuid.New()
		now := time.Now()

		rows := sqlmock.NewRows([]string{"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type", "notes"}).
			AddRow(id, userID, now, int64(1800), 1000, 20, 50, "pool", "Note")

		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnRows(rows)

		activities, err := repo.GetActivitiesByUser(userID)
		assert.NoError(t, err)
		assert.Len(t, activities, 1)
		assert.Equal(t, "Note", activities[0].Notes)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		userID := uuid.New()
		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnError(assert.AnError)

		activities, err := repo.GetActivitiesByUser(userID)
		assert.Error(t, err)
		assert.Empty(t, activities)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		userID := uuid.New()
		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type", "notes"}).
				AddRow("invalid_id", userID, time.Now(), int64(1800), 1000, 20, 50, "pool", "Note"))

		activities, err := repo.GetActivitiesByUser(userID)
		assert.Error(t, err)
		assert.Empty(t, activities)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetActivityByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	t.Run("success", func(t *testing.T) {
		activityID := uuid.New()
		userID := uuid.New()
		now := time.Now()
		durationSeconds := int64(2700)
		distance := float64(1200)
		laps := 24
		poolSize := float64(25)
		locationType := "pool"
		notes := "Test GetActivityByID"

		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities WHERE id = \$1`).
			WithArgs(activityID).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type", "notes",
			}).AddRow(activityID, userID, now, durationSeconds, distance, laps, poolSize, locationType, notes))

		activity, err := repo.GetActivityByID(activityID)
		assert.NoError(t, err)
		assert.Equal(t, activityID, activity.ID)
		assert.Equal(t, userID, activity.UserID)
		assert.Equal(t, now.Unix(), activity.Start.Unix())
		assert.Equal(t, domain.DurationString((time.Duration(durationSeconds) * time.Second).String()), activity.Duration)
		assert.Equal(t, distance, activity.Distance)
		assert.Equal(t, laps, activity.Laps)
		assert.Equal(t, poolSize, activity.PoolSize)
		assert.Equal(t, domain.LocationType(locationType), activity.LocationType)
		assert.Equal(t, notes, activity.Notes)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		activityID := uuid.New()
		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities WHERE id = \$1`).
			WithArgs(activityID).
			WillReturnError(assert.AnError)

		activity, err := repo.GetActivityByID(activityID)
		assert.Error(t, err)
		assert.Equal(t, domain.Activity{}, activity)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		activityID := uuid.New()
		mock.ExpectQuery(`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities WHERE id = \$1`).
			WithArgs(activityID).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "user_id", "start", "duration", "distance", "laps", "pool_size", "location_type", "notes",
			}).AddRow("invalid_id", "invalid_user_id", time.Now(), int64(1800), 1000, 20, 50, "pool", "Note"))

		activity, err := repo.GetActivityByID(activityID)
		assert.Error(t, err)
		assert.Equal(t, domain.Activity{}, activity)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdateActivity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	t.Run("success", func(t *testing.T) {
		activity := domain.Activity{
			ID:           uuid.New(),
			UserID:       uuid.New(),
			Start:        time.Now(),
			Duration:     domain.DurationString((45 * time.Minute).String()),
			Distance:     1500,
			Laps:         30,
			PoolSize:     25,
			LocationType: domain.LocationType(domain.LocationPool),
			Notes:        "Updated notes",
		}

		mock.ExpectExec(`UPDATE activities SET`).
			WithArgs(
				activity.ID,
				activity.UserID,
				activity.Start,
				int64(activity.Duration.Seconds()),
				activity.Distance,
				activity.Laps,
				activity.PoolSize,
				string(activity.LocationType),
				activity.Notes,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.UpdateActivity(activity)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("update error", func(t *testing.T) {
		activity := domain.Activity{
			ID:           uuid.New(),
			UserID:       uuid.New(),
			Start:        time.Now(),
			Duration:     domain.DurationString((45 * time.Minute).String()),
			Distance:     1500,
			Laps:         30,
			PoolSize:     25,
			LocationType: domain.LocationType(domain.LocationPool),
			Notes:        "Updated notes",
		}

		mock.ExpectExec(`UPDATE activities SET`).
			WithArgs(
				activity.ID,
				activity.UserID,
				activity.Start,
				int64(activity.Duration.Seconds()),
				activity.Distance,
				activity.Laps,
				activity.PoolSize,
				string(activity.LocationType),
				activity.Notes,
			).
			WillReturnError(assert.AnError)

		err := repo.UpdateActivity(activity)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDeleteActivity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)

	t.Run("success", func(t *testing.T) {
		activityID := uuid.New()

		mock.ExpectExec(`DELETE FROM activities WHERE id = \$1`).
			WithArgs(activityID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.DeleteActivity(activityID)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("delete error", func(t *testing.T) {
		activityID := uuid.New()

		mock.ExpectExec(`DELETE FROM activities WHERE id = \$1`).
			WithArgs(activityID).
			WillReturnError(assert.AnError)

		err := repo.DeleteActivity(activityID)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
