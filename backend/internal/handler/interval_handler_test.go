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
		newInterval := domain.Interval{
			ID:         1,
			ActivityID: 1,
			StartTime:  time.Date(2024, 6, 1, 10, 0, 0, 0, time.UTC),
			Duration:   time.Duration(30 * time.Minute),
			Distance:   1000,
			Type:       domain.IntervalType("swim"),
			Stroke:     domain.StrokeType("freestyle"),
			Notes:      "Test interval",
		}
		mockService.On("CreateInterval", newInterval).Return(nil)

		body, _ := json.Marshal(newInterval)
		req, _ := http.NewRequest(http.MethodPost, "/intervals", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		mockService.AssertCalled(t, "CreateInterval", newInterval)
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

		newInterval := domain.Interval{
			ID:         1,
			ActivityID: 1,
			StartTime:  time.Date(2024, 6, 1, 10, 0, 0, 0, time.UTC),
			Duration:   time.Duration(30 * time.Minute),
			Distance:   1000,
			Type:       domain.IntervalType("swim"),
			Stroke:     domain.StrokeType("freestyle"),
			Notes:      "Test interval",
		}
		mockService.On("CreateInterval", mock.MatchedBy(func(i domain.Interval) bool {
			return i.ID == newInterval.ID &&
				i.ActivityID == newInterval.ActivityID &&
				i.StartTime.Equal(newInterval.StartTime) &&
				i.Duration == newInterval.Duration &&
				i.Distance == newInterval.Distance &&
				i.Type == newInterval.Type &&
				i.Stroke == newInterval.Stroke &&
				i.Notes == newInterval.Notes
		})).Return(errors.New("service error"))

		body, _ := json.Marshal(newInterval)
		req, _ := http.NewRequest(http.MethodPost, "/intervals", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertCalled(t, "CreateInterval", newInterval)
	})
}
