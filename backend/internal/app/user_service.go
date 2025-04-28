package app

import "github.com/liviaruegger/MAC0350/backend/internal/domain"

type UserService struct {
    repo interface {
        CreateUser(user domain.User) error
        GetUserByID(id int) (domain.User, error)
    }
}

func NewUserService(r interface {
    CreateUser(user domain.User) error
    GetUserByID(id int) (domain.User, error)
}) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) CreateUser(user domain.User) error {
    return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (domain.User, error) {
    return s.repo.GetUserByID(id)
}