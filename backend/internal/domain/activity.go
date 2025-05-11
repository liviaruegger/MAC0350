package domain

import (
	"time"
)

// LocationType defines the kind of environment for the swim
type LocationType string

// Predefined location types
const (
	LocationPool      LocationType = "pool"
	LocationOpenWater LocationType = "open_water"
)

// Activity represents a full swim session
type Activity struct {
	ID           uint          `json:"id"`
	UserID       uint          `json:"user_id"`
	Start        time.Time     `json:"start"`         // Start time of the activity
	Duration     time.Duration `json:"duration"`      // Duration of the activity
	Distance     float64       `json:"distance"`      // Total distance in meters
	Laps         int           `json:"laps"`          // Number of pool laps
	PoolSize     float64       `json:"pool_size"`     // Pool length in meters (0 if open water)
	LocationType LocationType  `json:"location_type"` // "pool" or "open_water"
	Intervals    []Interval    `json:"intervals"`     // Breakdown of the swim into segments
}
