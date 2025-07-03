package handler

import "github.com/liviaruegger/MAC0350/backend/internal/entity"

// ErrorResponse represents an error message returned by the API.
// swagger:model
type ErrorResponse struct {
	// Error is a description of what went wrong.
	// Example: Service error
	Error string `json:"error"`
}

// MessageResponse is used for generic messages.
// swagger:model
type MessageResponse struct {
	// Message is a human-readable message.
	// Example: User created successfully
	Message string `json:"message"`
}

// GetActivitiesByUserResponse includes the list of all activities logged by the user
// swagger:model
type GetActivitiesByUserResponse struct {
	Activities []entity.Activity `json:"activities"`
}
