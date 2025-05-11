package domain

import (
	"testing"
	"time"
)

func TestPacePer100m(t *testing.T) {
	tests := []struct {
		name     string
		interval Interval
		expected float64
	}{
		{
			name: "Valid swim interval",
			interval: Interval{
				Duration: 120 * time.Second,
				Distance: 100,
				Type:     IntervalSwim,
			},
			expected: 120,
		},
		{
			name: "Rest interval",
			interval: Interval{
				Duration: 60 * time.Second,
				Distance: 0,
				Type:     IntervalRest,
			},
			expected: 0,
		},
		{
			name: "Zero distance",
			interval: Interval{
				Duration: 60 * time.Second,
				Distance: 0,
				Type:     IntervalSwim,
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.interval.PacePer100m()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPaceFormatted(t *testing.T) {
	tests := []struct {
		name     string
		interval Interval
		expected string
	}{
		{
			name: "Valid swim interval",
			interval: Interval{
				Duration: 125 * time.Second,
				Distance: 100,
				Type:     IntervalSwim,
			},
			expected: "02:05",
		},
		{
			name: "Rest interval",
			interval: Interval{
				Duration: 60 * time.Second,
				Distance: 0,
				Type:     IntervalRest,
			},
			expected: "N/A",
		},
		{
			name: "Zero distance",
			interval: Interval{
				Duration: 60 * time.Second,
				Distance: 0,
				Type:     IntervalSwim,
			},
			expected: "N/A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.interval.PaceFormatted()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
