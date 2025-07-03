package mapper

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMapActivityToEntity(t *testing.T) {
	activity := domain.Activity{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		Start:        time.Now(),
		Duration:     domain.DurationString("1h30m"),
		Distance:     4000,
		Laps:         160,
		PoolSize:     25,
		LocationType: domain.LocationPool,
	}

	intervals := []domain.Interval{
		{
			ID:         uuid.New(),
			ActivityID: activity.ID,
			Duration:   domain.DurationString("30m"),
			Distance:   1000,
			Type:       domain.IntervalSwim,
			Stroke:     domain.StrokeFreestyle,
			Notes:      "Test interval",
		},
	}

	entity := MapActivityToEntity(activity, intervals)

	assert.Equal(t, activity.ID, entity.ID)
	assert.Equal(t, activity.UserID, entity.UserID)
	assert.Equal(t, activity.Start, entity.Start)
	assert.Equal(t, activity.Duration, entity.Duration)
	assert.Equal(t, activity.Distance, entity.Distance)
	assert.Equal(t, activity.Laps, entity.Laps)
	assert.Equal(t, activity.PoolSize, entity.PoolSize)
	assert.Equal(t, string(activity.LocationType), string(entity.LocationType))
	assert.Equal(t, activity.Notes, entity.Notes)
	assert.Len(t, entity.Intervals, len(intervals))
}
