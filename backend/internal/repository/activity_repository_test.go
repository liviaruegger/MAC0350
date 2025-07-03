package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func fakeActivity() domain.Activity {
	return domain.Activity{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Date:         "2023-10-01",
		Start:        time.Now(),
		Duration:     domain.DurationString((30 * time.Minute).String()),
		Distance:     1000,
		Laps:         20,
		PoolSize:     50,
		LocationType: domain.LocationPool,
		LocationName: "CEPE",
		Feeling:      domain.FeelingTired,
		HeartRateAvg: 120,
		HeartRateMax: 140,
		Notes:        "Test notes",
	}
}

func TestCreateActivity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)
	activity := fakeActivity()

	mock.ExpectExec(`INSERT INTO activities`).
		WithArgs(
			activity.ID,
			activity.UserID,
			activity.Date,
			activity.Start,
			int64(activity.Duration.Seconds()),
			activity.Distance,
			activity.Laps,
			activity.PoolSize,
			string(activity.LocationType),
			activity.LocationName,
			string(activity.Feeling),
			activity.HeartRateAvg,
			activity.HeartRateMax,
			activity.Notes,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateActivity(activity)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllActivities(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "date", "start", "duration", "distance", "laps", "pool_size",
		"location_type", "location_name", "feeling", "heart_rate_avg", "heart_rate_max", "notes",
	}).AddRow(
		uuid.New(), uuid.New(), "2023-10-01", now, int64(1800), 1000, 20, 50,
		"pool", "CEPE", "tired", 120, 140, "notes",
	)

	mock.ExpectQuery(`SELECT id, user_id, date, start, duration, distance, laps, pool_size, location_type, location_name, feeling, heart_rate_avg, heart_rate_max, notes FROM activities`).
		WillReturnRows(rows)

	activities, err := repo.GetAllActivities()
	assert.NoError(t, err)
	assert.Len(t, activities, 1)
	assert.Equal(t, "CEPE", activities[0].LocationName)
	assert.Equal(t, domain.FeelingTired, activities[0].Feeling)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetActivitiesByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)
	activity := fakeActivity()

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "date", "start", "duration", "distance", "laps", "pool_size",
		"location_type", "location_name", "feeling", "heart_rate_avg", "heart_rate_max", "notes",
	}).AddRow(
		activity.ID, activity.UserID, activity.Date, activity.Start, int64(activity.Duration.Seconds()), activity.Distance, activity.Laps, activity.PoolSize,
		string(activity.LocationType), activity.LocationName, string(activity.Feeling), activity.HeartRateAvg, activity.HeartRateMax, activity.Notes,
	)

	mock.ExpectQuery(`SELECT id, user_id, date, start, duration, distance, laps, pool_size, location_type, location_name, feeling, heart_rate_avg, heart_rate_max, notes FROM activities WHERE user_id = \$1`).
		WithArgs(activity.UserID).
		WillReturnRows(rows)

	result, err := repo.GetActivitiesByUser(activity.UserID)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, activity.ID, result[0].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetActivityByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)
	activity := fakeActivity()

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "date", "start", "duration", "distance", "laps", "pool_size",
		"location_type", "location_name", "feeling", "heart_rate_avg", "heart_rate_max", "notes",
	}).AddRow(
		activity.ID, activity.UserID, activity.Date, activity.Start, int64(activity.Duration.Seconds()), activity.Distance, activity.Laps, activity.PoolSize,
		string(activity.LocationType), activity.LocationName, string(activity.Feeling), activity.HeartRateAvg, activity.HeartRateMax, activity.Notes,
	)

	mock.ExpectQuery(`SELECT id, user_id, date, start, duration, distance, laps, pool_size, location_type, location_name, feeling, heart_rate_avg, heart_rate_max, notes FROM activities WHERE id = \$1`).
		WithArgs(activity.ID).
		WillReturnRows(rows)

	result, err := repo.GetActivityByID(activity.ID)
	assert.NoError(t, err)
	assert.Equal(t, activity.UserID, result.UserID)
	assert.Equal(t, activity.Feeling, result.Feeling)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateActivity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)
	activity := fakeActivity()

	mock.ExpectExec(`UPDATE activities SET`).
		WithArgs(
			activity.ID,
			activity.UserID,
			activity.Date,
			activity.Start,
			int64(activity.Duration.Seconds()),
			activity.Distance,
			activity.Laps,
			activity.PoolSize,
			string(activity.LocationType),
			activity.LocationName,
			string(activity.Feeling),
			activity.HeartRateAvg,
			activity.HeartRateMax,
			activity.Notes,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateActivity(activity)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteActivity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewActivityRepository(db)
	id := uuid.New()

	mock.ExpectExec(`DELETE FROM activities WHERE id = \$1`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteActivity(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
