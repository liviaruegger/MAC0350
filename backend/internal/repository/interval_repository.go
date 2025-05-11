package repository

import (
	"database/sql"

	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// IntervalRepository defines the interface for the interval repository
type IntervalRepository interface {
	Create(interval domain.Interval) error
}

// PostgresInvervalRepository is a concrete implementation of IntervalRepository using PostgreSQL
type PostgresIntervalRepository struct {
	db *sql.DB
}

// NewIntervalRepository creates a new PostgresIntervalRepository
func NewIntervalRepository(db *sql.DB) *PostgresIntervalRepository {
	return &PostgresIntervalRepository{db: db}
}

func (r *PostgresIntervalRepository) Create(interval domain.Interval) error {
	_, err := r.db.Exec(`
		INSERT INTO intervals (
			activity_id, start_time, duration, distance, type, stroke, notes
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		interval.ActivityID,
		interval.StartTime,
		int64(interval.Duration.Seconds()), // store duration in seconds
		interval.Distance,
		string(interval.Type),
		string(interval.Stroke),
		interval.Notes,
	)
	return err
}
