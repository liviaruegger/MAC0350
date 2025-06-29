package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	City  string    `json:"city"`
	Phone string    `json:"phone"`
}
