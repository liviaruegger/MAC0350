package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// IntervalRepository defines the interface for the interval repository
type IntervalRepository interface {
	CreateInterval(interval domain.Interval) error
	GetIntervalsByActivity(activityID uuid.UUID) ([]domain.Interval, error)
}

// PostgresIntervalRepository is a concrete implementation of IntervalRepository using PostgreSQL
type PostgresIntervalRepository struct {
	db *sql.DB
}

// NewIntervalRepository creates a new PostgresIntervalRepository
func NewIntervalRepository(db *sql.DB) *PostgresIntervalRepository {
	return &PostgresIntervalRepository{db: db}
}

func (r *PostgresIntervalRepository) CreateInterval(interval domain.Interval) error {
	intervalID := uuid.New()

	_, err := r.db.Exec(`
		INSERT INTO intervals (
			id, activity_id, duration, distance, type, stroke, notes
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`,
		intervalID,
		interval.ActivityID,
		int64(interval.Duration.Seconds()),
		interval.Distance,
		string(interval.Type),
		string(interval.Stroke),
		interval.Notes,
	)
	return err
}

func (r *PostgresIntervalRepository) GetIntervalsByActivity(activityID uuid.UUID) ([]domain.Interval, error) {
	rows, err := r.db.Query(`
		SELECT id, activity_id, duration, distance, type, stroke, notes
		FROM intervals WHERE activity_id = $1
	`, activityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var intervals []domain.Interval
	for rows.Next() {
		var interval domain.Interval
		var durationSeconds int64
		if err := rows.Scan(
			&interval.ID,
			&interval.ActivityID,
			&durationSeconds,
			&interval.Distance,
			&interval.Type,
			&interval.Stroke,
			&interval.Notes,
		); err != nil {
			return nil, err
		}
		interval.Duration = domain.DurationString((time.Duration(durationSeconds) * time.Second).String())
		intervals = append(intervals, interval)
	}

	return intervals, nil
}
