package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// LocationType defines the kind of environment for the swim
type LocationType string

// Predefined location types
const (
	LocationPool      LocationType = "pool"
	LocationOpenWater LocationType = "open_water"
)

// Activity is the internal struct to represent a swimming activity or session
type Activity struct {
	// ID is the unique identifier for the activity (PK)
	ID uuid.UUID `json:"id"`
	// UserID is the ID of the user who performed the activity (FK)
	UserID uuid.UUID `json:"user_id"`
	// Start time of the activity
	Start time.Time `json:"start"`
	// Duration of the activity in string format, e.g., "1h30m"
	Duration domain.DurationString `json:"duration"`
	// Total distance in meters
	Distance float64 `json:"distance"`
	// Number of pool laps
	Laps int `json:"laps"`
	// Pool length in meters (0 if open water)
	PoolSize float64 `json:"pool_size"`
	// "pool" or "open_water"
	LocationType LocationType `json:"location_type"`
	// Optional notes
	Notes string `json:"notes"`
	// Intervals are the segments of the swim session
	Intervals []Interval `json:"intervals"`
}
