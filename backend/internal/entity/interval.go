package entity

import (
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

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

// StrokeType defines the style of swimming stroke used
type StrokeType string

// Predefined stroke types
const (
	// StrokeFreestyle represents front crawl
	StrokeFreestyle StrokeType = "freestyle"
	// StrokeBackstroke represents backstroke
	StrokeBackstroke StrokeType = "backstroke"
	// StrokeBreaststroke represents breaststroke
	StrokeBreaststroke StrokeType = "breaststroke"
	// StrokeButterfly represents butterfly stroke
	StrokeButterfly StrokeType = "butterfly"
	// StrokeMedley represents individual medley (IM)
	StrokeMedley StrokeType = "medley"
	// StrokeUnknown is used when the stroke is not specified
	StrokeUnknown StrokeType = "unknown"
)

// Interval is the internal struct to represent a single segment of a swim session
type Interval struct {
	ID uuid.UUID `json:"id"`
	// Foreign key to the swim activity/session
	ActivityID uuid.UUID `json:"activity_id"`
	// Duration of the interval in string format, e.g., "1h30m"
	Duration domain.DurationString `json:"duration"`
	// Distance in meters
	Distance float64 `json:"distance"`
	// One of the predefined types
	Type IntervalType `json:"type"`
	// Type of swimming stroke
	Stroke StrokeType `json:"stroke"`
	// Optional notes like "felt strong", "used fins"
	Notes string `json:"notes"`
}
