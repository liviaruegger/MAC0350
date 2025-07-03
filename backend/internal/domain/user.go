package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	City   string    `json:"city"`
	Phone  string    `json:"phone"`
	Age    int       `json:"age"`
	Height int       `json:"height"`
	Weight float64   `json:"weight"`
}
