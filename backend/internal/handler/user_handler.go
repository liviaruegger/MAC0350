package handler

import (
	"github.com/liviaruegger/MAC0350/backend/internal/app"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *app.UserService
}

func NewUserHandler(s *app.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var newUser domain.User
    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    if err := h.service.CreateUser(newUser); err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, newUser)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := h.service.GetUserByID(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.IndentedJSON(http.StatusOK, user)
}