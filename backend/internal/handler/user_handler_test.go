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

func (m *MockUserService) GetAllUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id int) (domain.User, error) {
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

func (m *MockUserService) DeleteUser(id int) error {
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
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
			City:  "São Paulo",
			Phone: "+55 11 91234-5678",
		}
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
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		newUser := domain.User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
			City:  "São Paulo",
			Phone: "+55 11 91234-5678",
		}
		mockService.On("CreateUser", mock.MatchedBy(func(u domain.User) bool {
			return u.ID == newUser.ID &&
				u.Name == newUser.Name &&
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
		mockService.AssertCalled(t, "CreateUser", newUser)
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

		users := []domain.User{
			{
				ID:    1,
				Name:  "John Doe",
				Email: "john@example.com",
				City:  "São Paulo",
				Phone: "+55 11 91234-5678",
			},
			{
				ID:    2,
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
		user := domain.User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
			City:  "São Paulo",
			Phone: "+55 11 91234-5678",
		}
		mockService.On("GetUserByID", 1).Return(user, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		// Validate the response body
		var returned domain.User
		err := json.Unmarshal(resp.Body.Bytes(), &returned)
		assert.NoError(t, err)
		assert.Equal(t, user, returned)

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

func TestGetUserByEmail(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users/email/:email", handler.GetUserByEmail)

	t.Run("success", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		user := domain.User{
			ID:    1,
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

		// Use an invalid email format as parameter
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

		updatedUser := domain.User{
			ID:    1,
			Name:  "John Updated",
			Email: "john.updated@example.com",
			City:  "Campinas",
			Phone: "+55 19 91234-5678",
		}
		mockService.On("UpdateUser", updatedUser).Return(nil)

		body, _ := json.Marshal(updatedUser)
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(body))
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
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.Calls = nil

		updatedUser := domain.User{
			ID:    2,
			Name:  "Not Found",
			Email: "notfound@example.com",
			City:  "Nowhere",
			Phone: "000",
		}
		mockService.On("UpdateUser", updatedUser).Return(errors.New("user not found"))

		body, _ := json.Marshal(updatedUser)
		req, _ := http.NewRequest(http.MethodPut, "/users/2", bytes.NewBuffer(body))
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

		mockService.On("DeleteUser", 1).Return(nil)

		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNoContent, resp.Code)
		mockService.AssertCalled(t, "DeleteUser", 1)
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

		mockService.On("DeleteUser", 2).Return(errors.New("user not found"))

		req, _ := http.NewRequest(http.MethodDelete, "/users/2", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		mockService.AssertCalled(t, "DeleteUser", 2)
	})
}
