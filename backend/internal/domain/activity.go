package domain

import (
	"time"
)

type Activity struct {
	ID       uint      `json:"id"`
	Distance string    `json:"distance"`
	Start    time.Time `json:"start"`
	Finish   time.Time `json:"finish"`
	Size     time.Time `json:"size"`
	Laps     int       `json:"laps"`
	UserID   uint 	   `json:"user_id"`
}