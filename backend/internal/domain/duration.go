package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

// DurationString represents a duration stored as a string,
// using Go's standard duration format, e.g., "1h30m0s".
type DurationString string

// MarshalJSON serializes the DurationString into JSON as a string,
// ensuring it remains in the Go duration format (e.g., "2h45m").
func (d DurationString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON deserializes a JSON string into a DurationString,
// validating that it conforms to the Go duration format.
func (d *DurationString) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if _, err := time.ParseDuration(s); err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	*d = DurationString(s)
	return nil
}

func (d DurationString) ToDuration() time.Duration {
	duration, err := time.ParseDuration(string(d))
	if err != nil {
		return time.Duration(0)
	}
	return duration
}

func (d DurationString) Seconds() float64 {
	return d.ToDuration().Seconds()
}
