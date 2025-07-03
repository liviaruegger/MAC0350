package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock ActivityService ---

type MockActivityService struct {
	mock.Mock
}

func (m *MockActivityService) CreateActivity(activity domain.Activity) error {
	args := m.Called(activity)
	return args.Error(0)
}

func (m *MockActivityService) GetAllActivities() ([]domain.Activity, error) {
	args := m.Called()
	return args.Get(0).([]domain.Activity), args.Error(1)
}

func (m *MockActivityService) GetActivitiesByUser(userID uuid.UUID) ([]entity.Activity, error) {
	args := m.Called(userID)

	var activities []entity.Activity
	if raw := args.Get(0); raw != nil {
		activities = raw.([]entity.Activity)
	} else {
		activities = []entity.Activity{}
	}

	return activities, args.Error(1)
}

func (m *MockActivityService) GetActivityByID(id uuid.UUID) (domain.Activity, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Activity), args.Error(1)
}

func (m *MockActivityService) UpdateActivity(activity domain.Activity) error {
	args := m.Called(activity)
	return args.Error(0)
}

func (m *MockActivityService) DeleteActivity(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

// --- Tests ---

func TestCreateActivityHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockActivityService)
	handler := NewActivityHandler(mockService)

	router := gin.Default()
	router.POST("/activities", handler.CreateActivity)

	t.Run("success", func(t *testing.T) {
		reqBody := CreateActivityRequest{
			UserID:       uuid.New(),
			Duration:     domain.DurationString("30m"),
			Distance:     1500,
			Laps:         30,
			PoolSize:     50,
			LocationType: domain.LocationPool,
			Notes:        "Morning swim",
		}

		body, _ := json.Marshal(reqBody)

		mockService.On("CreateActivity", mock.MatchedBy(func(a domain.Activity) bool {
			return a.UserID == reqBody.UserID && a.Distance == reqBody.Distance
		})).Return(nil)

		req, _ := http.NewRequest(http.MethodPost, "/activities", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/activities", bytes.NewBuffer([]byte("invalid")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		reqBody := CreateActivityRequest{
			UserID:       uuid.New(),
			Duration:     domain.DurationString("1h"),
			Distance:     2000,
			Laps:         40,
			PoolSize:     50,
			LocationType: domain.LocationOpenWater,
			Notes:        "",
		}

		body, _ := json.Marshal(reqBody)

		mockService.On("CreateActivity", mock.Anything).Return(errors.New("internal error"))

		req, _ := http.NewRequest(http.MethodPost, "/activities", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestGetActivitiesByUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockActivityService)
	handler := NewActivityHandler(mockService)

	router := gin.Default()
	router.GET("/users/:id/activities", handler.GetActivitiesByUser)

	t.Run("success", func(t *testing.T) {
		userID := uuid.New()
		mockService.On("GetActivitiesByUser", userID).Return([]entity.Activity{
			{ID: uuid.New(), Distance: 1500},
		}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/"+userID.String()+"/activities", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "1500") // Check that the response contains the expected activity
		mockService.AssertExpectations(t)
	})

	t.Run("invalid UUID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/not-a-uuid/activities", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("no activities", func(t *testing.T) {
		userID := uuid.New()
		mockService.On("GetActivitiesByUser", userID).Return([]entity.Activity{}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/"+userID.String()+"/activities", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		userID := uuid.New()
		mockService.On("GetActivitiesByUser", userID).Return(nil, errors.New("db error"))

		req, _ := http.NewRequest(http.MethodGet, "/users/"+userID.String()+"/activities", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
