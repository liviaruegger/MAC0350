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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetAllUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id uuid.UUID) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserService) GetUserByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/users", handler.CreateUser)

	t.Run("success", func(t *testing.T) {
		newUser := domain.User{
			Name:   "John Doe",
			Email:  "john@example.com",
			City:   "São Paulo",
			Phone:  "+55 11 91234-5678",
			Age:    30,
			Height: 175,
			Weight: 72.5,
		}

		mockService.On("CreateUser", mock.MatchedBy(func(u domain.User) bool {
			return u.Name == newUser.Name &&
				u.Email == newUser.Email &&
				u.City == newUser.City &&
				u.Phone == newUser.Phone &&
				u.Age == newUser.Age &&
				u.Height == newUser.Height &&
				u.Weight == newUser.Weight
		})).Return(nil)

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		newUser := domain.User{
			Name:   "John Doe",
			Email:  "john@example.com",
			City:   "São Paulo",
			Phone:  "+55 11 91234-5678",
			Age:    30,
			Height: 175,
			Weight: 72.5,
		}

		mockService.On("CreateUser", mock.Anything).Return(errors.New("service error"))

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestGetAllUsers(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", handler.GetAllUsers)

	t.Run("success", func(t *testing.T) {
		users := []domain.User{
			{
				ID:     uuid.New(),
				Name:   "John Doe",
				Email:  "john@example.com",
				City:   "São Paulo",
				Phone:  "+55 11 91234-5678",
				Age:    30,
				Height: 175,
				Weight: 72.5,
			},
			{
				ID:     uuid.New(),
				Name:   "Jane Doe",
				Email:  "jane@example.com",
				City:   "Rio de Janeiro",
				Phone:  "+55 21 98765-4321",
				Age:    28,
				Height: 165,
				Weight: 60.0,
			},
		}

		mockService.On("GetAllUsers").Return(users, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var returned []domain.User
		err := json.Unmarshal(resp.Body.Bytes(), &returned)
		assert.NoError(t, err)
		assert.Equal(t, users, returned)
	})

	t.Run("service error", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		mockService.On("GetAllUsers").Return([]domain.User{}, errors.New("service error"))

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestGetUserByID(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users/:id", handler.GetUserByID)

	t.Run("success", func(t *testing.T) {
		userID := uuid.New()
		user := domain.User{
			ID:     userID,
			Name:   "John Doe",
			Email:  "john@example.com",
			City:   "São Paulo",
			Phone:  "+55 11 91234-5678",
			Age:    30,
			Height: 175,
			Weight: 72.5,
		}
		mockService.On("GetUserByID", userID).Return(user, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/"+userID.String(), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var returned domain.User
		err := json.Unmarshal(resp.Body.Bytes(), &returned)
		assert.NoError(t, err)
		assert.Equal(t, user, returned)
	})

	t.Run("invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/invalid-id", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("not found", func(t *testing.T) {
		id := uuid.New()
		mockService.On("GetUserByID", id).Return(domain.User{}, errors.New("not found"))

		req, _ := http.NewRequest(http.MethodGet, "/users/"+id.String(), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}
