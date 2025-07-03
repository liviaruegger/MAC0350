package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockIntervalService is a mock implementation of app.IntervalService
type MockIntervalService struct {
	mock.Mock
}

func (m *MockIntervalService) CreateInterval(interval domain.Interval) error {
	args := m.Called(interval)
	return args.Error(0)
}

func (m *MockIntervalService) GetIntervalsByActivity(activityID uuid.UUID) ([]domain.Interval, error) {
	args := m.Called(activityID)
	if raw := args.Get(0); raw != nil {
		return raw.([]domain.Interval), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateInterval(t *testing.T) {
	mockService := new(MockIntervalService)
	handler := NewIntervalHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/intervals", handler.CreateInterval)

	t.Run("success", func(t *testing.T) {
		newIntervalReq := CreateIntervalRequest{
			ActivityID: uuid.New(),
			Duration:   domain.DurationString((30 * time.Minute).String()),
			Distance:   1000,
			Type:       domain.IntervalType("swim"),
			Stroke:     domain.StrokeType("freestyle"),
			Notes:      "Test interval",
		}

		// Usamos MatchedBy para ignorar o ID aleatório gerado
		mockService.On("CreateInterval", mock.MatchedBy(func(i domain.Interval) bool {
			return i.ActivityID == newIntervalReq.ActivityID &&
				i.Duration == newIntervalReq.Duration &&
				i.Distance == newIntervalReq.Distance &&
				i.Type == newIntervalReq.Type &&
				i.Stroke == newIntervalReq.Stroke &&
				i.Notes == newIntervalReq.Notes
		})).Return(nil)

		body, _ := json.Marshal(newIntervalReq)
		req, _ := http.NewRequest(http.MethodPost, "/intervals", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/intervals", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		newIntervalReq := CreateIntervalRequest{
			ActivityID: uuid.New(),
			Duration:   domain.DurationString((30 * time.Minute).String()),
			Distance:   1000,
			Type:       domain.IntervalType("swim"),
			Stroke:     domain.StrokeType("freestyle"),
			Notes:      "Test interval",
		}

		mockService.On("CreateInterval", mock.MatchedBy(func(i domain.Interval) bool {
			return i.ActivityID == newIntervalReq.ActivityID &&
				i.Duration == newIntervalReq.Duration &&
				i.Distance == newIntervalReq.Distance &&
				i.Type == newIntervalReq.Type &&
				i.Stroke == newIntervalReq.Stroke &&
				i.Notes == newIntervalReq.Notes
		})).Return(errors.New("service error"))

		body, _ := json.Marshal(newIntervalReq)
		req, _ := http.NewRequest(http.MethodPost, "/intervals", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertExpectations(t)
	})
}
