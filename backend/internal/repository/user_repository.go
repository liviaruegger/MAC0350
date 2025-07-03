package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
)

// UserRepository defines the interface for the user repository
type UserRepository interface {
	CreateUser(user domain.User) error
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id uuid.UUID) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUser(id uuid.UUID) error
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
		`INSERT INTO users (id, name, email, city, phone, age, height, weight)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.ID, user.Name, user.Email, user.City, user.Phone, user.Age, user.Height, user.Weight,
	)
	return err
}

func (r *PostgresUserRepository) GetAllUsers() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, city, phone, age, height, weight FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.City, &user.Phone, &user.Age, &user.Height, &user.Weight); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *PostgresUserRepository) GetUserByID(id uuid.UUID) (domain.User, error) {
	var user domain.User
	row := r.db.QueryRow("SELECT id, name, email, city, phone, age, height, weight FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.City, &user.Phone, &user.Age, &user.Height, &user.Weight)
	return user, err
}

func (r *PostgresUserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	row := r.db.QueryRow("SELECT id, name, email, city, phone, age, height, weight FROM users WHERE email = $1", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.City, &user.Phone, &user.Age, &user.Height, &user.Weight)
	if err == sql.ErrNoRows {
		return user, nil // No user found with the given email
	}
	return user, err
}

func (r *PostgresUserRepository) UpdateUser(user domain.User) error {
	_, err := r.db.Exec(
		`UPDATE users 
		 SET name = $1, email = $2, city = $3, phone = $4, age = $5, height = $6, weight = $7 
		 WHERE id = $8`,
		user.Name, user.Email, user.City, user.Phone, user.Age, user.Height, user.Weight, user.ID,
	)
	return err
}

func (r *PostgresUserRepository) DeleteUser(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
