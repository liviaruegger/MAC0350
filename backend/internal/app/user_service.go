package app

import (
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user domain.User) error
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id uuid.UUID) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUser(id uuid.UUID) error
}

// UserService provides user-related operations
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(r repository.UserRepository) *userService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(user domain.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id uuid.UUID) (domain.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) GetUserByEmail(email string) (domain.User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *userService) UpdateUser(user domain.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
