package handler

import (
	"github.com/liviaruegger/MAC0350/backend/internal/app"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"

	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service app.UserService
}

func NewUserHandler(s app.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Creates a user with the provided name, email, city, and phone
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User data."
// @Success 201 {object} domain.User "User successfully created"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser domain.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON"})
		return
	}

	if err := h.service.CreateUser(newUser); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Error: "Service error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Returns a list of all users with their name, email, city, and phone
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} domain.User "List of users"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Error: "Service error"})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Returns the user with name, email, city, and phone for the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User "User found"
// @Failure 400 {object} ErrorResponse "Invalid user ID"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// GetUserByEmail godoc
// @Summary Get user by email
// @Description Returns the user with name, email, city, and phone for the specified email
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} domain.User "User found"
// @Failure 400 {object} ErrorResponse "Invalid email"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /users/email/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	// Simple email validation
	if !isValidEmail(email) {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid email format"})
		return
	}

	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// isValidEmail checks if the email has a basic valid format.
func isValidEmail(email string) bool {
	// Very simple regex for demonstration; consider using a more robust one in production
	re := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(re, email)
	return matched
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Updates the user with the provided ID, name, email, city, and phone
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body domain.User true "Updated user data"
// @Success 200 {object} domain.User "User successfully updated"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	var updatedUser domain.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON"})
		return
	}
	updatedUser.ID = uint(id)

	if err := h.service.UpdateUser(updatedUser); err != nil {
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes the user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "User successfully deleted"
// @Failure 400 {object} ErrorResponse "Invalid user ID"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
