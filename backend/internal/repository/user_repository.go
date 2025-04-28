package repository

import (
    "database/sql"
    "github.com/liviaruegger/MAC0350/backend/internal/domain"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user domain.User) error {
    _, err := r.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
    return err
}

func (r *UserRepository) GetUserByID(id int) (domain.User, error) {
    var user domain.User
    row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    return user, err
}