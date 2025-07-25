package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/app"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// ActivityHandler handles HTTP requests related to activities
type ActivityHandler struct {
	service app.ActivityService
}

// NewActivityHandler creates a new ActivityHandler
func NewActivityHandler(service app.ActivityService) *ActivityHandler {
	return &ActivityHandler{service: service}
}

// CreateActivity godoc
// @Summary Create a new activity
// @Description Creates a swim activity for a specific user
// @Tags activities
// @Accept json
// @Produce json
// @Param activity body handler.CreateActivityRequest true "Activity data"
// @Success 201 {object} domain.Activity "Activity successfully created"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /activities [post]
func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	var req CreateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON"})
		return
	}

	activity := domain.Activity{
		ID:           uuid.New(),
		UserID:       req.UserID,
		Date:         req.Date,
		Start:        time.Now(), // Field 'Start' is currently unused by the frontend
		Duration:     req.Duration,
		Distance:     req.Distance,
		Laps:         req.Laps,
		PoolSize:     req.PoolSize,
		LocationType: req.LocationType,
		LocationName: req.LocationName,
		Feeling:      req.Feeling,
		HeartRateAvg: req.HeartRateAvg,
		HeartRateMax: req.HeartRateMax,
		Notes:        req.Notes,
	}

	if err := h.service.CreateActivity(activity); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Error: "Service error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, activity)
}

// GetAllActivities godoc
// @Summary Get all activities
// @Description Retrieves all swim activities in the system
// @Tags activities
// @Accept json
// @Produce json
// @Success 200 {array} domain.Activity "List of all activities"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /activities [get]
func (h *ActivityHandler) GetAllActivities(c *gin.Context) {
	activities, err := h.service.GetAllActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to retrieve activities"})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// GetActivitiesByUser godoc
// @Summary Get all activities of a user
// @Description Retrieves all swim activities and their intervals for a given user ID
// @Tags activities
// @Accept json
// @Produce json
// @Param user_id path string true "User ID (UUID)"
// @Success 200 {object} GetActivitiesByUserResponse
// @Failure 400 {object} ErrorResponse "Invalid user ID"
// @Failure 404 {object} ErrorResponse "User not found or no activities"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users/{user_id}/activities [get]
func (h *ActivityHandler) GetActivitiesByUser(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	activities, err := h.service.GetActivitiesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to retrieve activities"})
		return
	}

	if len(activities) == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "No activities found for user"})
		return
	}

	response := GetActivitiesByUserResponse{
		Activities: activities,
	}

	c.JSON(http.StatusOK, response)
}
