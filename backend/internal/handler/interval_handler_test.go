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

		expectedInterval := domain.Interval{
			ActivityID: newIntervalReq.ActivityID,
			Duration:   newIntervalReq.Duration,
			Distance:   newIntervalReq.Distance,
			Type:       newIntervalReq.Type,
			Stroke:     newIntervalReq.Stroke,
			Notes:      newIntervalReq.Notes,
		}

		mockService.On("CreateInterval", expectedInterval).Return(nil)

		body, _ := json.Marshal(newIntervalReq)
		req, _ := http.NewRequest(http.MethodPost, "/intervals", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		mockService.AssertCalled(t, "CreateInterval", expectedInterval)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/intervals", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		// Clear previous expectations
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		newIntervalReq := CreateIntervalRequest{
			ActivityID: uuid.New(),
			Duration:   domain.DurationString((30 * time.Minute).String()),
			Distance:   1000,
			Type:       domain.IntervalType("swim"),
			Stroke:     domain.StrokeType("freestyle"),
			Notes:      "Test interval",
		}

		expectedInterval := domain.Interval{
			ActivityID: newIntervalReq.ActivityID,
			Duration:   newIntervalReq.Duration,
			Distance:   newIntervalReq.Distance,
			Type:       newIntervalReq.Type,
			Stroke:     newIntervalReq.Stroke,
			Notes:      newIntervalReq.Notes,
		}

		mockService.On("CreateInterval", mock.MatchedBy(func(i domain.Interval) bool {
			return i.ActivityID == expectedInterval.ActivityID &&
				i.Duration == expectedInterval.Duration &&
				i.Distance == expectedInterval.Distance &&
				i.Type == expectedInterval.Type &&
				i.Stroke == expectedInterval.Stroke &&
				i.Notes == expectedInterval.Notes
		})).Return(errors.New("service error"))

		body, _ := json.Marshal(newIntervalReq)
		req, _ := http.NewRequest(http.MethodPost, "/intervals", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertCalled(t, "CreateInterval", expectedInterval)
	})
}
