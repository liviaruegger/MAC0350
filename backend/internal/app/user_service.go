package app

import (
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/liviaruegger/MAC0350/backend/internal/repository"
)

// UserService provides user-related operations
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(user domain.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (domain.User, error) {
	return s.repo.GetUserByID(id)
}
