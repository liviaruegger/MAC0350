package repository

import (
	"database/sql"

	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// UserRepository defines the interface for the user repository
type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserByID(id int) (domain.User, error)
}

// PostgresUserRepository is a concrete implementation of UserRepository using a PostgreSQL database connection
type PostgresUserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new PostgresUserRepository
func NewUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(user domain.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (name, email, city, phone) VALUES ($1, $2, $3, $4)",
		user.Name, user.Email, user.City, user.Phone,
	)
	return err
}

func (r *PostgresUserRepository) GetUserByID(id int) (domain.User, error) {
	var user domain.User
	row := r.db.QueryRow("SELECT id, name, email, city, phone FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.City, &user.Phone)
	return user, err
}
