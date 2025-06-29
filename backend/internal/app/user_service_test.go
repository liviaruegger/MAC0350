package app

import (
	"errors"
	"testing"

	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// Mock repository
type mockUserRepo struct {
	users map[int]domain.User
}

func (m *mockUserRepo) CreateUser(user domain.User) error {
	if _, exists := m.users[int(user.ID)]; exists {
		return errors.New("user already exists")
	}
	m.users[int(user.ID)] = user
	return nil
}

func (m *mockUserRepo) GetAllUsers() ([]domain.User, error) {
	var userList []domain.User
	for _, user := range m.users {
		userList = append(userList, user)
	}
	return userList, nil
}

func (m *mockUserRepo) GetUserByID(id int) (domain.User, error) {
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
	if _, exists := m.users[int(user.ID)]; !exists {
		return errors.New("user not found")
	}
	m.users[int(user.ID)] = user
	return nil
}

func (m *mockUserRepo) DeleteUser(id int) error {
	if _, exists := m.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(m.users, id)
	return nil
}

func TestUserService(t *testing.T) {
	repo := &mockUserRepo{users: make(map[int]domain.User)}
	service := NewUserService(repo)

	user := domain.User{
		ID:    1,
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
	retrieved, err := service.GetUserByID(1)
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
	if retrievedByEmail.ID != 1 {
		t.Errorf("expected user ID 1, got: %d", retrievedByEmail.ID)
	}

	// Test UpdateUser
	user.Name = "Ana Paula"
	err = service.UpdateUser(user)
	if err != nil {
		t.Fatalf("expected no error on update, got: %v", err)
	}
	updated, _ := service.GetUserByID(1)
	if updated.Name != "Ana Paula" {
		t.Errorf("expected updated name 'Ana Paula', got: %s", updated.Name)
	}

	// Test DeleteUser
	err = service.DeleteUser(1)
	if err != nil {
		t.Fatalf("expected no error on delete, got: %v", err)
	}
	_, err = service.GetUserByID(1)
	if err == nil {
		t.Errorf("expected error after deleting user, got nil")
	}
}
