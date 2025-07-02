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
// @Param interval body handler.CreateIntervalRequest true "Interval data"
// @Success 201 {object} domain.Interval "Interval successfully created"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /intervals [post]
func (h *IntervalHandler) CreateInterval(c *gin.Context) {
	var req CreateIntervalRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON or missing required fields"})
		return
	}

	interval := domain.Interval{
		ActivityID: req.ActivityID,
		Duration:   req.Duration,
		Distance:   req.Distance,
		Type:       req.Type,
		Stroke:     req.Stroke,
		Notes:      req.Notes,
	}

	if err := h.service.CreateInterval(interval); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Error: "Service error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, interval)
}
