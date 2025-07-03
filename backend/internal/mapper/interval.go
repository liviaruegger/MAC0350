package mapper

import (
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/entity"
)

// MapIntervalToEntity maps a domain.Interval to an entity.Interval
func MapIntervalToEntity(interval domain.Interval) entity.Interval {
	return entity.Interval{
		ID:         interval.ID,
		ActivityID: interval.ActivityID,
		Duration:   interval.Duration,
		Distance:   interval.Distance,
		Type:       entity.IntervalType(interval.Type),
		Stroke:     entity.StrokeType(interval.Stroke),
		Notes:      interval.Notes,
	}
}
