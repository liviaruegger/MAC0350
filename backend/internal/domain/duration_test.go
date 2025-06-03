package domain

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDurationString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		duration DurationString
		want     string
	}{
		{"zero", "0s", `"0s"`},
		{"seconds", "5s", `"5s"`},
		{"minutes", "2m0s", `"2m0s"`},
		{"hours", "1h30m0s", `"1h30m0s"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.duration)
			if err != nil {
				t.Fatalf("MarshalJSON() error = %v", err)
			}
			if string(got) != tt.want {
				t.Errorf("MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestDurationString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    DurationString
		wantErr bool
	}{
		{"zero", `"0s"`, "0s", false},
		{"seconds", `"10s"`, "10s", false},
		{"minutes", `"2m0s"`, "2m0s", false},
		{"hours", `"1h30m0s"`, "1h30m0s", false},
		{"invalid format", `"abc"`, "", true},
		{"not a string", `123`, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d DurationString
			err := json.Unmarshal([]byte(tt.input), &d)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && d != tt.want {
				t.Errorf("UnmarshalJSON() = %v, want %v", d, tt.want)
			}
		})
	}
}

func TestDurationString_RoundTrip(t *testing.T) {
	orig := DurationString("42m7s")
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("MarshalJSON() error = %v", err)
	}
	var decoded DurationString
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("UnmarshalJSON() error = %v", err)
	}
	if decoded != orig {
		t.Errorf("RoundTrip: got %v, want %v", decoded, orig)
	}
}

func TestDurationString_ToDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration DurationString
		expected time.Duration
	}{
		{"zero", "0s", time.Duration(0)},
		{"seconds", "5s", 5 * time.Second},
		{"minutes", "2m0s", 2 * time.Minute},
		{"hours", "1h30m0s", 1*time.Hour + 30*time.Minute},
		{"invalid format", "abc", 0}, // should return 0 on invalid format
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.duration.ToDuration()
			if got != tt.expected {
				t.Errorf("ToDuration() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestDurationString_Seconds(t *testing.T) {
	tests := []struct {
		name     string
		duration DurationString
		expected float64
	}{
		{"zero", "0s", 0},
		{"seconds", "5s", 5},
		{"minutes", "2m0s", 120},
		{"hours", "1h30m0s", 5400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.duration.Seconds()
			if got != tt.expected {
				t.Errorf("Seconds() = %f, want %f", got, tt.expected)
			}
		})
	}
}
