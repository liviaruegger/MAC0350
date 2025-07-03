package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// ActivityRepository defines the interface for the activity repository
type ActivityRepository interface {
	CreateActivity(activity domain.Activity) error
	GetAllActivities() ([]domain.Activity, error)
	GetActivitiesByUser(userID uuid.UUID) ([]domain.Activity, error)
	GetActivityByID(activityID uuid.UUID) (domain.Activity, error)
	UpdateActivity(activity domain.Activity) error
	DeleteActivity(activityID uuid.UUID) error
}

type PostgresActivityRepository struct {
	db *sql.DB
}

func NewActivityRepository(db *sql.DB) *PostgresActivityRepository {
	return &PostgresActivityRepository{db: db}
}

func (r *PostgresActivityRepository) CreateActivity(activity domain.Activity) error {
	_, err := r.db.Exec(
		`INSERT INTO activities (
			id, user_id, date, start, duration, distance, laps, pool_size,
			location_type, location_name, feeling, heart_rate_avg, heart_rate_max, notes
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
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
	)

	return err
}

func (r *PostgresActivityRepository) GetAllActivities() ([]domain.Activity, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, date, start, duration, distance, laps, pool_size,
		        location_type, location_name, feeling, heart_rate_avg, heart_rate_max, notes
		 FROM activities`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []domain.Activity
	for rows.Next() {
		var a domain.Activity
		var durationSeconds int64
		var locationType, feeling string

		err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Date,
			&a.Start,
			&durationSeconds,
			&a.Distance,
			&a.Laps,
			&a.PoolSize,
			&locationType,
			&a.LocationName,
			&feeling,
			&a.HeartRateAvg,
			&a.HeartRateMax,
			&a.Notes,
		)
		if err != nil {
			return nil, err
		}

		a.Duration = domain.DurationString((time.Duration(durationSeconds) * time.Second).String())
		a.LocationType = domain.LocationType(locationType)
		a.Feeling = domain.FeelingType(feeling)

		activities = append(activities, a)
	}

	return activities, nil
}

func (r *PostgresActivityRepository) GetActivitiesByUser(userID uuid.UUID) ([]domain.Activity, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, date, start, duration, distance, laps, pool_size,
		        location_type, location_name, feeling, heart_rate_avg, heart_rate_max, notes
		 FROM activities
		 WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []domain.Activity
	for rows.Next() {
		var a domain.Activity
		var durationSeconds int64
		var locationType, feeling string

		err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Date,
			&a.Start,
			&durationSeconds,
			&a.Distance,
			&a.Laps,
			&a.PoolSize,
			&locationType,
			&a.LocationName,
			&feeling,
			&a.HeartRateAvg,
			&a.HeartRateMax,
			&a.Notes,
		)
		if err != nil {
			return nil, err
		}

		a.Duration = domain.DurationString((time.Duration(durationSeconds) * time.Second).String())
		a.LocationType = domain.LocationType(locationType)
		a.Feeling = domain.FeelingType(feeling)

		activities = append(activities, a)
	}

	return activities, nil
}

func (r *PostgresActivityRepository) GetActivityByID(activityID uuid.UUID) (domain.Activity, error) {
	var a domain.Activity
	var durationSeconds int64
	var locationType, feeling string

	err := r.db.QueryRow(
		`SELECT id, user_id, date, start, duration, distance, laps, pool_size,
		        location_type, location_name, feeling, heart_rate_avg, heart_rate_max, notes
		 FROM activities
		 WHERE id = $1`,
		activityID,
	).Scan(
		&a.ID,
		&a.UserID,
		&a.Date,
		&a.Start,
		&durationSeconds,
		&a.Distance,
		&a.Laps,
		&a.PoolSize,
		&locationType,
		&a.LocationName,
		&feeling,
		&a.HeartRateAvg,
		&a.HeartRateMax,
		&a.Notes,
	)
	if err != nil {
		return a, err
	}

	a.Duration = domain.DurationString((time.Duration(durationSeconds) * time.Second).String())
	a.LocationType = domain.LocationType(locationType)
	a.Feeling = domain.FeelingType(feeling)

	return a, nil
}

func (r *PostgresActivityRepository) UpdateActivity(activity domain.Activity) error {
	_, err := r.db.Exec(
		`UPDATE activities SET
			user_id = $2,
			date = $3,
			start = $4,
			duration = $5,
			distance = $6,
			laps = $7,
			pool_size = $8,
			location_type = $9,
			location_name = $10,
			feeling = $11,
			heart_rate_avg = $12,
			heart_rate_max = $13,
			notes = $14
		WHERE id = $1`,
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
	)

	return err
}

func (r *PostgresActivityRepository) DeleteActivity(activityID uuid.UUID) error {
	_, err := r.db.Exec(
		`DELETE FROM activities WHERE id = $1`,
		activityID,
	)

	return err
}
