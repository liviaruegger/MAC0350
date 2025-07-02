package mapper

import (
	"testing"

	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMapIntervalToEntity(t *testing.T) {
	interval := domain.Interval{
		ID:         uuid.New(),
		ActivityID: uuid.New(),
		Duration:   domain.DurationString("30m"),
		Distance:   1000,
		Type:       domain.IntervalSwim,
		Stroke:     domain.StrokeFreestyle,
		Notes:      "Test interval",
	}

	entity := MapIntervalToEntity(interval)

	assert.Equal(t, interval.ID, entity.ID)
	assert.Equal(t, interval.ActivityID, entity.ActivityID)
	assert.Equal(t, interval.Duration, entity.Duration)
	assert.Equal(t, interval.Distance, entity.Distance)
	assert.Equal(t, string(interval.Type), string(entity.Type))
	assert.Equal(t, string(interval.Stroke), string(entity.Stroke))
	assert.Equal(t, interval.Notes, entity.Notes)
}
