package domain

import (
	"fmt"
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
	Notes        string        `json:"notes"`         // Optional notes
}

// AvgPacePer100m returns the average pace in seconds per 100 meters
func (a Activity) AvgPacePer100m() float64 {
	if a.Distance == 0 {
		return 0
	}

	return a.Duration.Seconds() / a.Distance * 100
}

// AvgPaceFormatted returns the pace as a string in the format mm:ss per 100m;
// If distance is 0, it returns "N/A"
func (a Activity) AvgPaceFormatted() string {
	pace := a.AvgPacePer100m()
	if pace == 0 {
		return "N/A"
	}

	minutes := int(pace) / 60
	seconds := int(pace) % 60

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
