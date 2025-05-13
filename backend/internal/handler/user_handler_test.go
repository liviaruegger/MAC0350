package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of app.UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetUserByID(id int) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/users", handler.CreateUser)

	t.Run("success", func(t *testing.T) {
		newUser := domain.User{ID: 1, Name: "John Doe"}
		mockService.On("CreateUser", newUser).Return(nil)

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		mockService.AssertCalled(t, "CreateUser", newUser)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		// Clear previous expectations
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		newUser := domain.User{ID: 1, Name: "John Doe"}
		mockService.On("CreateUser", mock.MatchedBy(func(u domain.User) bool {
			return u.ID == newUser.ID && u.Name == newUser.Name
		})).Return(errors.New("service error"))

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertCalled(t, "CreateUser", newUser)
	})
}

func TestGetUserByID(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users/:id", handler.GetUserByID)

	t.Run("success", func(t *testing.T) {
		user := domain.User{ID: 1, Name: "John Doe"}
		mockService.On("GetUserByID", 1).Return(user, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		mockService.AssertCalled(t, "GetUserByID", 1)
	})

	t.Run("invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/invalid", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		mockService.On("GetUserByID", 2).Return(domain.User{}, errors.New("user not found"))

		req, _ := http.NewRequest(http.MethodGet, "/users/2", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		mockService.AssertCalled(t, "GetUserByID", 2)
	})
}
