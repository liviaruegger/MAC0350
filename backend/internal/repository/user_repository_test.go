package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	user := domain.User{
		ID:     uuid.New(),
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		City:   "São Paulo",
		Phone:  "+5511999999999",
		Age:    30,
		Height: 170,
		Weight: 65.5,
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.ID, user.Name, user.Email, user.City, user.Phone, user.Age, user.Height, user.Weight).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	expectedUser := domain.User{
		ID:     uuid.New(),
		Name:   "John Doe",
		Email:  "john@example.com",
		City:   "São Paulo",
		Phone:  "+5511999999999",
		Age:    30,
		Height: 170,
		Weight: 65.5,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "city", "phone", "age", "height", "weight"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email, expectedUser.City, expectedUser.Phone,
			expectedUser.Age, expectedUser.Height, expectedUser.Weight)

	mock.ExpectQuery("SELECT id, name, email, city, phone, age, height, weight FROM users").WillReturnRows(rows)

	users, err := repo.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, expectedUser, users[0])
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	expectedUser := domain.User{
		ID:     uuid.New(),
		Name:   "Jane Smith",
		Email:  "jane@example.com",
		City:   "Rio de Janeiro",
		Phone:  "+5521999999999",
		Age:    25,
		Height: 165,
		Weight: 55.0,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "city", "phone", "age", "height", "weight"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email, expectedUser.City, expectedUser.Phone,
			expectedUser.Age, expectedUser.Height, expectedUser.Weight)

	mock.ExpectQuery("SELECT id, name, email, city, phone, age, height, weight FROM users WHERE id =").
		WithArgs(expectedUser.ID).
		WillReturnRows(rows)

	user, err := repo.GetUserByID(expectedUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	user := domain.User{
		ID:     uuid.New(),
		Name:   "John Updated",
		Email:  "john.updated@example.com",
		City:   "Campinas",
		Phone:  "+5511987654321",
		Age:    35,
		Height: 175,
		Weight: 70.2,
	}

	mock.ExpectExec("UPDATE users").
		WithArgs(user.Name, user.Email, user.City, user.Phone, user.Age, user.Height, user.Weight, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)
	id := uuid.New()

	mock.ExpectExec("DELETE FROM users WHERE id =").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteUser(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
