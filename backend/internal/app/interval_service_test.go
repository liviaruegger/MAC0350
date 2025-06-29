package app

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// mockIntervalRepository is a mock implementation of IntervalRepository
type mockIntervalRepository struct {
	createFunc func(domain.Interval) error
}

func (m *mockIntervalRepository) Create(interval domain.Interval) error {
	if m.createFunc != nil {
		return m.createFunc(interval)
	}
	return nil
}

func TestNewIntervalService(t *testing.T) {
	mockRepo := &mockIntervalRepository{}
	service := NewIntervalService(mockRepo)
	if service == nil {
		t.Fatal("expected non-nil service")
	}
}

func TestCreateInterval(t *testing.T) {
	tests := []struct {
		name        string
		createFunc  func(domain.Interval) error
		interval    domain.Interval
		expectedErr error
	}{
		{
			name: "Success",
			createFunc: func(interval domain.Interval) error {
				return nil
			},
			interval: domain.Interval{
				ID:         uuid.New(),
				ActivityID: uuid.New(),
				StartTime:  time.Date(2024, 6, 1, 10, 0, 0, 0, time.UTC),
				Duration:   domain.DurationString((30 * time.Minute).String()),
				Distance:   1000,
				Type:       domain.IntervalType("swim"),
				Stroke:     domain.StrokeType("freestyle"),
				Notes:      "Test interval",
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			createFunc: func(interval domain.Interval) error {
				return errors.New("repo error")
			},
			interval: domain.Interval{
				ID:         uuid.New(),
				ActivityID: uuid.New(),
			},
			expectedErr: errors.New("repo error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockIntervalRepository{
				createFunc: tc.createFunc,
			}
			service := NewIntervalService(mockRepo)
			err := service.CreateInterval(tc.interval)
			if tc.expectedErr == nil && err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if tc.expectedErr != nil && (err == nil || err.Error() != tc.expectedErr.Error()) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
