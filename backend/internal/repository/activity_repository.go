package repository

import (
	"database/sql"
	"time"

	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// ActivityRepository defines the interface for the activity repository
type ActivityRepository interface {
	Create(activity domain.Activity) error
	GetAll() ([]domain.Activity, error)
	GetAllByUser(userID int) ([]domain.Activity, error)
}

// PostgresActivityRepository is a concrete implementation of ActivityRepository using PostgreSQL
type PostgresActivityRepository struct {
	db *sql.DB
}

// NewActivityRepository creates a new PostgresActivityRepository
func NewActivityRepository(db *sql.DB) *PostgresActivityRepository {
	return &PostgresActivityRepository{db: db}
}

func (r *PostgresActivityRepository) Create(activity domain.Activity) error {
	_, err := r.db.Exec(
		`INSERT INTO activities (
			user_id, start, duration, distance, laps, pool_size, location_type
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		activity.UserID,
		activity.Start,
		int64(activity.Duration.Seconds()),
		activity.Distance,
		activity.Laps,
		activity.PoolSize,
		string(activity.LocationType),
	)
	return err
}

func (r *PostgresActivityRepository) GetAll() ([]domain.Activity, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type
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
		)
		if err != nil {
			return nil, err
		}

		a.Duration = time.Duration(durationSeconds) * time.Second
		a.LocationType = domain.LocationType(locationType)

		activities = append(activities, a)
	}

	return activities, nil
}

func (r *PostgresActivityRepository) GetAllByUser(userID int) ([]domain.Activity, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, start, duration, distance, laps, pool_size, location_type
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
		)
		if err != nil {
			return nil, err
		}

		a.Duration = time.Duration(durationSeconds) * time.Second
		a.LocationType = domain.LocationType(locationType)

		activities = append(activities, a)
	}

	return activities, nil
}
