package handler

// CreateUserRequest represents the request body for creating a new user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	City  string `json:"city" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
