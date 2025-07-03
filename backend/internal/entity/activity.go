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

// FeelingType defines options for how a swimmer feels after a session
type FeelingType string

// Predefined feelings
const (
	FeelingExcellent FeelingType = "excellent"
	FeelingGood      FeelingType = "good"
	FeelingRegular   FeelingType = "regular"
	FeelingTired     FeelingType = "tired"
	FeelingBad       FeelingType = "bad"
)

// Activity is the internal struct to represent a swimming activity or session
type Activity struct {
	// ID is the unique identifier for the activity (PK)
	ID uuid.UUID `json:"id"`
	// UserID is the ID of the user who performed the activity (FK)
	UserID uuid.UUID `json:"user_id"`
	// Date in ISO 8601 format, e.g., "2023-10-01"
	Date string `json:"date"`
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
	// Optional name for the location, e.g., "CEPE"
	LocationName string `json:"location_name,omitempty"`
	// Optional feeling after the swim, e.g., "tired"
	Feeling FeelingType `json:"feeling,omitempty"`
	// Average heart rate during the activity
	HeartRateAvg int `json:"heart_rate_avg,omitempty"`
	// Maximum heart rate during the activity
	HeartRateMax int `json:"heart_rate_max,omitempty"`
	// Average pace in seconds per 100 meters, formatted mm:ss
	AvgPacePer100m string `json:"avg_pace_per_100m,omitempty"`
	// Optional notes
	Notes string `json:"notes"`
	// Intervals are the segments of the swim session
	Intervals []Interval `json:"intervals"`
}
