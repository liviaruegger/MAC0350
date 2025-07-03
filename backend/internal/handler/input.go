package handler

import (
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// CreateUserRequest represents the request body for creating a new user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	City  string `json:"city" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

// CreateActivityRequest represents the request body for creating a new activity
type CreateActivityRequest struct {
	// ID of the user who performed the activity
	UserID uuid.UUID `json:"user_id" binding:"required"`
	// Date in ISO 8601 format, e.g., "2023-10-01"
	Date string `json:"date" binding:"required"`
	// Start time of the activity
	// Start time.Time `json:"start"` // TODO - must implement format handling
	// Duration of the activity in a string format, e.g., "1h30m"
	Duration domain.DurationString `json:"duration" binding:"required"`
	// Total distance in meters
	Distance float64 `json:"distance" binding:"required"`
	// Number of pool laps
	Laps int `json:"laps" binding:"required"`
	// Pool size in meters (0 if open water)
	PoolSize float64 `json:"pool_size" binding:"required"`
	// "pool" or "open_water"
	LocationType domain.LocationType `json:"location_type" binding:"required"`
	// Optional name for the location, e.g., "CEPE"
	LocationName string `json:"location_name,omitempty"`
	// Optional feeling after the swim, e.g., "tired"
	Feeling domain.FeelingType `json:"feeling,omitempty"`
	// Average heart rate during the activity
	HeartRateAvg int `json:"heart_rate_avg,omitempty"`
	// Maximum heart rate during the activity
	HeartRateMax int `json:"heart_rate_max,omitempty"`
	// Optional notes
	Notes string `json:"notes"`
}

// GetActivitiesByUserRequest represents the request parameters for fetching activities by user ID
type GetActivitiesByUserRequest struct {
	// UserID is the ID of the user whose activities are being requested
	UserID uuid.UUID `json:"user_id" binding:"required"`
}

// CreateIntervalRequest represents the request body for creating a new interval
type CreateIntervalRequest struct {
	// ActivityID is the ID of the associated activity/session
	ActivityID uuid.UUID `json:"activity_id" binding:"required"`
	// Duration of the interval in string format, e.g., "1h30m"
	Duration domain.DurationString `json:"duration" binding:"required"`
	// Distance in meters
	Distance float64 `json:"distance" binding:"required"`
	// Type is one of the predefined interval types like "swim", "rest", etc.
	Type domain.IntervalType `json:"type" binding:"required"`
	// Stroke is the swimming stroke type like "freestyle", "backstroke", etc.
	Stroke domain.StrokeType `json:"stroke" binding:"required"`
	// Notes are optional remarks such as "felt strong", "used fins"
	Notes string `json:"notes"`
}
