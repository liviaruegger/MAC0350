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

// MockUserService is a mock implementation of app.UserService
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
			Name:  "John Doe",
			Email: "john@example.com",
			City:  "São Paulo",
			Phone: "+55 11 91234-5678",
		}
		mockService.On("CreateUser", mock.MatchedBy(func(u domain.User) bool {
			return u.Name == newUser.Name &&
				u.Email == newUser.Email &&
				u.City == newUser.City &&
				u.Phone == newUser.Phone
		})).Return(nil)

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		mockService.AssertCalled(t, "CreateUser", mock.MatchedBy(func(u domain.User) bool {
			return u.Name == newUser.Name &&
				u.Email == newUser.Email &&
				u.City == newUser.City &&
				u.Phone == newUser.Phone
		}))
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
			Name:  "John Doe",
			Email: "john@example.com",
			City:  "São Paulo",
			Phone: "+55 11 91234-5678",
		}
		mockService.On("CreateUser", mock.MatchedBy(func(u domain.User) bool {
			return u.Name == newUser.Name &&
				u.Email == newUser.Email &&
				u.City == newUser.City &&
				u.Phone == newUser.Phone
		})).Return(errors.New("service error"))

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertCalled(t, "CreateUser", mock.MatchedBy(func(u domain.User) bool {
			return u.Name == newUser.Name &&
				u.Email == newUser.Email &&
				u.City == newUser.City &&
				u.Phone == newUser.Phone
		}))
	})
}

func TestGetAllUsers(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", handler.GetAllUsers)

	t.Run("success", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		userID1 := uuid.New()
		userID2 := uuid.New()
		users := []domain.User{
			{
				ID:    userID1,
				Name:  "John Doe",
				Email: "john@example.com",
				City:  "São Paulo",
				Phone: "+55 11 91234-5678",
			},
			{
				ID:    userID2,
				Name:  "Jane Doe",
				Email: "jane@example.com",
				City:  "Rio de Janeiro",
				Phone: "+55 21 98765-4321",
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

		mockService.AssertCalled(t, "GetAllUsers")
	})

	t.Run("service error", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		mockService.On("GetAllUsers").Return([]domain.User{}, errors.New("service error"))

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertCalled(t, "GetAllUsers")
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
			ID:    userID,
			Name:  "John Doe",
			Email: "john@example.com",
			City:  "São Paulo",
			Phone: "+55 11 91234-5678",
		}
		mockService.On("GetUserByID", userID).Return(user, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/"+userID.String(), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		// Validate the response body
		var returned domain.User
		err := json.Unmarshal(resp.Body.Bytes(), &returned)
		assert.NoError(t, err)
		assert.Equal(t, user, returned)

		mockService.AssertCalled(t, "GetUserByID", userID)
	})

	t.Run("invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/invalid", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		notFoundID := uuid.New()
		mockService.On("GetUserByID", notFoundID).Return(domain.User{}, errors.New("user not found"))

		req, _ := http.NewRequest(http.MethodGet, "/users/"+notFoundID.String(), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		mockService.AssertCalled(t, "GetUserByID", notFoundID)
	})
}

func TestGetUserByEmail(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users/email/:email", handler.GetUserByEmail)

	t.Run("success", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		userID := uuid.New()
		user := domain.User{
			ID:    userID,
			Name:  "John Doe",
			Email: "john@example.com",
			City:  "São Paulo",
			Phone: "+55 11 91234-5678",
		}
		mockService.On("GetUserByEmail", "john@example.com").Return(user, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/email/john@example.com", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var returned domain.User
		err := json.Unmarshal(resp.Body.Bytes(), &returned)
		assert.NoError(t, err)
		assert.Equal(t, user, returned)

		mockService.AssertCalled(t, "GetUserByEmail", "john@example.com")
	})

	t.Run("invalid email", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		req, _ := http.NewRequest(http.MethodGet, "/users/email/invalid-email", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		mockService.On("GetUserByEmail", "notfound@example.com").Return(domain.User{}, errors.New("user not found"))

		req, _ := http.NewRequest(http.MethodGet, "/users/email/notfound@example.com", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		mockService.AssertCalled(t, "GetUserByEmail", "notfound@example.com")
	})
}

func TestUpdateUser(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/users/:id", handler.UpdateUser)

	t.Run("success", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		userID := uuid.New()
		updatedUser := domain.User{
			ID:    userID,
			Name:  "John Updated",
			Email: "john.updated@example.com",
			City:  "Campinas",
			Phone: "+55 19 91234-5678",
		}
		mockService.On("UpdateUser", updatedUser).Return(nil)

		body, _ := json.Marshal(updatedUser)
		req, _ := http.NewRequest(http.MethodPut, "/users/"+userID.String(), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		mockService.AssertCalled(t, "UpdateUser", updatedUser)
	})

	t.Run("invalid user ID", func(t *testing.T) {
		body, _ := json.Marshal(domain.User{
			Name:  "Invalid ID",
			Email: "invalid@example.com",
			City:  "City",
			Phone: "123",
		})
		req, _ := http.NewRequest(http.MethodPut, "/users/invalid", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		userID := uuid.New()
		req, _ := http.NewRequest(http.MethodPut, "/users/"+userID.String(), bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		userID := uuid.New()
		updatedUser := domain.User{
			ID:    userID,
			Name:  "Not Found",
			Email: "notfound@example.com",
			City:  "Nowhere",
			Phone: "000",
		}
		mockService.On("UpdateUser", updatedUser).Return(errors.New("user not found"))

		body, _ := json.Marshal(updatedUser)
		req, _ := http.NewRequest(http.MethodPut, "/users/"+userID.String(), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		mockService.AssertCalled(t, "UpdateUser", updatedUser)
	})
}

func TestDeleteUser(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/users/:id", handler.DeleteUser)

	t.Run("success", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		userID := uuid.New()
		mockService.On("DeleteUser", userID).Return(nil)

		req, _ := http.NewRequest(http.MethodDelete, "/users/"+userID.String(), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNoContent, resp.Code)
		mockService.AssertCalled(t, "DeleteUser", userID)
	})

	t.Run("invalid user ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/invalid", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		notFoundID := uuid.New()
		mockService.On("DeleteUser", notFoundID).Return(errors.New("user not found"))

		req, _ := http.NewRequest(http.MethodDelete, "/users/"+notFoundID.String(), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		mockService.AssertCalled(t, "DeleteUser", notFoundID)
	})
}
