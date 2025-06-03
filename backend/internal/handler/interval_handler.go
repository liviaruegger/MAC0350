package handler

import (
	"github.com/liviaruegger/MAC0350/backend/internal/app"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)

type IntervalHandler struct {
	service app.IntervalService
}

func NewIntervalHandler(s app.IntervalService) *IntervalHandler {
	return &IntervalHandler{service: s}
}

// CreateInterval godoc
// @Summary Create a new interval
// @Description Creates an interval with the data provided in the request body
// @Tags intervals
// @Accept json
// @Produce json
// @Param interval body domain.Interval true "Interval data"
// @Success 201 {object} domain.Interval "Interval successfully created"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /intervals [post]
func (h *IntervalHandler) CreateInterval(c *gin.Context) {
	var newInterval domain.Interval
	if err := c.BindJSON(&newInterval); err != nil {
		return
	}

	if err := h.service.CreateInterval(newInterval); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Service error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newInterval)
}
