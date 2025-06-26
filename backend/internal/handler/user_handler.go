package handler

import (
	"github.com/liviaruegger/MAC0350/backend/internal/app"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"

	"net/http"
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
