package domain

// import (
// 	"time"
// )

type Activity struct {
	ID       uint      `json:"id"`
	Distance string    `json:"distance"`
	Start    int 	   `json:"start"`
	Finish   int	   `json:"finish"`
	Size     int	   `json:"size"`
	Laps     int       `json:"laps"`
	UserID   uint 	   `json:"user_id"`
}