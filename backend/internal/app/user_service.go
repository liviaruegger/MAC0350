package app

import (
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"
)

type UserService interface {
	CreateUser(user domain.User) error
	GetUserByID(id int) (domain.User, error)
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

func (s *userService) GetUserByID(id int) (domain.User, error) {
	return s.repo.GetUserByID(id)
}
