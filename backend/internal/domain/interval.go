package domain

import "time"

// IntervalType defines the kind of interval (e.g., swim, rest, drill)
type IntervalType string

// Predefined interval types
const (
	// IntervalSwim represents a generic swim interval
	IntervalSwim IntervalType = "swim"
	// IntervalRest represents a rest interval (break between sets)
	IntervalRest IntervalType = "rest"
	// IntervalDrill represents a drill interval (technique-focused)
	IntervalDrill IntervalType = "drill"
	// IntervalKick represents a kick interval (legs only)
	IntervalKick IntervalType = "kick"
	// IntervalPull represents a pull interval (arms only)
	IntervalPull IntervalType = "pull"
	// IntervalWarmUp represents a warm-up interval (low intensity)
	IntervalWarmUp IntervalType = "warmup"
	// IntervalMainSet represents the main set of the workout
	IntervalMainSet IntervalType = "main_set"
	// IntervalCoolDown represents a cool-down interval (low intensity)
	IntervalCoolDown IntervalType = "cooldown"
)

// Interval represents a single segment of a swim session
type Interval struct {
	ID         uint          `json:"id"`
	ActivityID uint          `json:"activity_id"` // Foreign key to the swim activity/session
	StartTime  time.Time     `json:"start_time"`
	Duration   time.Duration `json:"duration"` // Duration of the interval
	Distance   float64       `json:"distance"` // Distance in meters
	Type       IntervalType  `json:"type"`     // One of the predefined types
	Notes      string        `json:"notes"`    // Optional notes like "felt strong", "used fins"
}
