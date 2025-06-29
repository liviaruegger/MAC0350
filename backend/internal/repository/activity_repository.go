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
	GetAllActivitiesByUser(userID uuid.UUID) ([]domain.Activity, error)
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
			id, user_id, start, duration, distance, laps, pool_size, location_type, notes
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		activity.ID,
		activity.UserID,
		activity.Start,
		int64(activity.Duration.Seconds()),
		activity.Distance,
		activity.Laps,
		activity.PoolSize,
		string(activity.LocationType),
		activity.Notes,
	)

	return err
}

func (r *PostgresActivityRepository) GetAllActivities() ([]domain.Activity, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes FROM activities`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []domain.Activity
	for rows.Next() {
		var a domain.Activity
		var durationSeconds int64
		var locationType string

		err = rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Start,
			&durationSeconds,
			&a.Distance,
			&a.Laps,
			&a.PoolSize,
			&locationType,
			&a.Notes,
		)
		if err != nil {
			return nil, err
		}

		a.Duration = domain.DurationString((time.Duration(durationSeconds) * time.Second).String())
		a.LocationType = domain.LocationType(locationType)

		activities = append(activities, a)
	}

	return activities, nil
}

func (r *PostgresActivityRepository) GetAllActivitiesByUser(userID uuid.UUID) ([]domain.Activity, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type, notes
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
		var locationType string

		err = rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Start,
			&durationSeconds,
			&a.Distance,
			&a.Laps,
			&a.PoolSize,
			&locationType,
			&a.Notes,
		)
		if err != nil {
			return nil, err
		}

		a.Duration = domain.DurationString((time.Duration(durationSeconds) * time.Second).String())
		a.LocationType = domain.LocationType(locationType)

		activities = append(activities, a)
	}

	return activities, nil
}
