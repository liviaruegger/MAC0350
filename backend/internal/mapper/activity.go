package mapper

import (
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/entity"
)

// MapActivityToEntity maps a domain.Activity to an entity.Activity with intervals
func MapActivityToEntity(activity domain.Activity, intervals []domain.Interval) entity.Activity {
	mappedIntervals := make([]entity.Interval, len(intervals))
	for i, interval := range intervals {
		mappedIntervals[i] = MapIntervalToEntity(interval)
	}

	return entity.Activity{
		ID:           activity.ID,
		UserID:       activity.UserID,
		Start:        activity.Start,
		Duration:     activity.Duration,
		Distance:     activity.Distance,
		Laps:         activity.Laps,
		PoolSize:     activity.PoolSize,
		LocationType: entity.LocationType(activity.LocationType),
		Notes:        activity.Notes,
		Intervals:    mappedIntervals,
	}
}
