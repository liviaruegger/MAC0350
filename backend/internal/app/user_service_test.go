package app

import (
	"errors"
	"testing"

	"github.com/liviaruegger/MAC0350/backend/internal/domain"

	"github.com/google/uuid"
)

// Mock repository
type mockUserRepo struct {
	users map[uuid.UUID]domain.User
}

func (m *mockUserRepo) CreateUser(user domain.User) error {
	if _, exists := m.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepo) GetAllUsers() ([]domain.User, error) {
	var userList []domain.User
	for _, user := range m.users {
		userList = append(userList, user)
	}
	return userList, nil
}

func (m *mockUserRepo) GetUserByID(id uuid.UUID) (domain.User, error) {
	user, exists := m.users[id]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (m *mockUserRepo) GetUserByEmail(email string) (domain.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}

func (m *mockUserRepo) UpdateUser(user domain.User) error {
	if _, exists := m.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepo) DeleteUser(id uuid.UUID) error {
	if _, exists := m.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(m.users, id)
	return nil
}

func TestUserService(t *testing.T) {
	repo := &mockUserRepo{users: make(map[uuid.UUID]domain.User)}
	service := NewUserService(repo)

	userID := uuid.New()
	user := domain.User{
		ID:    userID,
		Name:  "Ana",
		Email: "ana@example.com",
	}

	// Test CreateUser
	err := service.CreateUser(user)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Test GetAllUsers
	users, err := service.GetAllUsers()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 user, got: %d", len(users))
	}

	// Test GetUserByID
	retrieved, err := service.GetUserByID(userID)
	if err != nil {
		t.Fatalf("expected to find user, got error: %v", err)
	}
	if retrieved.Name != "Ana" {
		t.Errorf("expected name 'Ana', got: %s", retrieved.Name)
	}

	// Test GetUserByEmail
	retrievedByEmail, err := service.GetUserByEmail("ana@example.com")
	if err != nil {
		t.Fatalf("expected to find user by email, got error: %v", err)
	}
	if retrievedByEmail.ID != userID {
		t.Errorf("expected user ID %v, got: %v", userID, retrievedByEmail.ID)
	}

	// Test UpdateUser
	user.Name = "Ana Paula"
	err = service.UpdateUser(user)
	if err != nil {
		t.Fatalf("expected no error on update, got: %v", err)
	}
	updated, _ := service.GetUserByID(userID)
	if updated.Name != "Ana Paula" {
		t.Errorf("expected updated name 'Ana Paula', got: %s", updated.Name)
	}

	// Test DeleteUser
	err = service.DeleteUser(userID)
	if err != nil {
		t.Fatalf("expected no error on delete, got: %v", err)
	}
	_, err = service.GetUserByID(userID)
	if err == nil {
		t.Errorf("expected error after deleting user, got nil")
	}
}
