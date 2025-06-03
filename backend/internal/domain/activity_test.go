package domain

import (
	"testing"
	"time"
)

func TestAvgPacePer100m(t *testing.T) {
	tests := []struct {
		name     string
		activity Activity
		expected float64
	}{
		{
			name: "Valid distance and duration",
			activity: Activity{
				Duration: DurationString((time.Duration(1500) * time.Second).String()), // 25 minutes
				Distance: 1000,                                                         // 1 km
			},
			expected: 150, // 150 seconds per 100m
		},
		{
			name: "Zero distance",
			activity: Activity{
				Duration: DurationString((time.Duration(1500) * time.Second).String()),
				Distance: 0,
			},
			expected: 0,
		},
		{
			name: "Zero duration",
			activity: Activity{
				Duration: DurationString((0 * time.Second).String()),
				Distance: 1000,
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.activity.AvgPacePer100m()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAvgPaceFormatted(t *testing.T) {
	tests := []struct {
		name     string
		activity Activity
		expected string
	}{
		{
			name: "Valid distance and duration",
			activity: Activity{
				Duration: DurationString((time.Duration(1500) * time.Second).String()), // 25 minutes
				Distance: 1000,                                                         // 1 km
			},
			expected: "02:30", // 2 minutes 30 seconds per 100m
		},
		{
			name: "Zero distance",
			activity: Activity{
				Duration: DurationString((time.Duration(1500) * time.Second).String()),
				Distance: 0,
			},
			expected: "N/A",
		},
		{
			name: "Zero duration",
			activity: Activity{
				Duration: DurationString((0 * time.Second).String()),
				Distance: 1000,
			},
			expected: "N/A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.activity.AvgPaceFormatted()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
