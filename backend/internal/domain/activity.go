package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
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
	// ID is the unique identifier for the activity (PK)
	ID uuid.UUID `json:"id"`
	// UserID is the ID of the user who performed the activity (FK)
	UserID uuid.UUID `json:"user_id"`
	// Start time of the activity
	Start time.Time `json:"start"`
	// Duration of the activity in string format, e.g., "1h30m"
	Duration DurationString `json:"duration"`
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
