package handler

import (
	"time"

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

// CreateIntervalRequest represents the request body for creating a new interval
type CreateIntervalRequest struct {
	// ActivityID is the ID of the associated activity/session
	ActivityID uuid.UUID `json:"activity_id" binding:"required"`
	// StartTime is the start time of the interval
	StartTime time.Time `json:"start_time" binding:"required"`
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
