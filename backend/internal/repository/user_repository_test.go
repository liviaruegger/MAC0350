package repository

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	user := domain.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Name, user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	expectedUser := domain.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email)

	mock.ExpectQuery("SELECT id, name, email FROM users WHERE id = \\$1").
		WithArgs(expectedUser.ID).
		WillReturnRows(rows)

	user, err := repo.GetUserByID(int(expectedUser.ID))
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}
