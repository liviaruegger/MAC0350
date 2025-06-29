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

	// Test GetUserByID
	retrieved, err := service.GetUserByID(1)
	if err != nil {
		t.Fatalf("expected to find user, got error: %v", err)
	}
	if retrieved.Name != "Ana" {
		t.Errorf("expected name 'Ana', got: %s", retrieved.Name)
	}
}
