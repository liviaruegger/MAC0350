package repository

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/liviaruegger/MAC0350/backend/internal/domain"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	user := domain.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		City:  "São Paulo",
		Phone: "+5511999999999",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Name, user.Email, user.City, user.Phone).
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

	// Use github.com/google/uuid for UUIDs
	id1 := uuid.New()
	id2 := uuid.New()

	expectedUsers := []domain.User{
		{
			ID:    id1,
			Name:  "John Doe",
			Email: "john.doe@example.com",
			City:  "São Paulo",
			Phone: "+5511999999999",
		},
		{
			ID:    id2,
			Name:  "Jane Doe",
			Email: "jane.doe@example.com",
			City:  "Rio de Janeiro",
			Phone: "+5521999999999",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "city", "phone"}).
		AddRow(expectedUsers[0].ID.String(), expectedUsers[0].Name, expectedUsers[0].Email, expectedUsers[0].City, expectedUsers[0].Phone).
		AddRow(expectedUsers[1].ID.String(), expectedUsers[1].Name, expectedUsers[1].Email, expectedUsers[1].City, expectedUsers[1].Phone)

	mock.ExpectQuery("SELECT id, name, email, city, phone FROM users").
		WillReturnRows(rows)

	users, err := repo.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllUsers_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	mock.ExpectQuery("SELECT id, name, email, city, phone FROM users").
		WillReturnError(assert.AnError)

	users, err := repo.GetAllUsers()
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllUsers_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	// Simulate a row with a wrong type (e.g., string instead of int for ID)
	rows := sqlmock.NewRows([]string{"id", "name", "email", "city", "phone"}).
		AddRow("not-an-int", "John Doe", "john.doe@example.com", "São Paulo", "+5511999999999")

	mock.ExpectQuery("SELECT id, name, email, city, phone FROM users").
		WillReturnRows(rows)

	users, err := repo.GetAllUsers()
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllUsers_RowsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "city", "phone"}).
		AddRow(1, "John Doe", "john.doe@example.com", "São Paulo", "+5511999999999")
	// Simulate error after reading rows
	rows.RowError(0, assert.AnError)

	mock.ExpectQuery("SELECT id, name, email, city, phone FROM users").
		WillReturnRows(rows)

	users, err := repo.GetAllUsers()
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	id := uuid.New()
	expectedUser := domain.User{
		ID:    id,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		City:  "São Paulo",
		Phone: "+5511999999999",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "city", "phone"}).
		AddRow(expectedUser.ID.String(), expectedUser.Name, expectedUser.Email, expectedUser.City, expectedUser.Phone)

	mock.ExpectQuery("SELECT id, name, email, city, phone FROM users WHERE id = \\$1").
		WithArgs(expectedUser.ID.String()).
		WillReturnRows(rows)

	user, err := repo.GetUserByID(expectedUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	id := uuid.New()
	expectedUser := domain.User{
		ID:    id,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		City:  "São Paulo",
		Phone: "+5511999999999",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "city", "phone"}).
		AddRow(expectedUser.ID.String(), expectedUser.Name, expectedUser.Email, expectedUser.City, expectedUser.Phone)

	mock.ExpectQuery("SELECT id, name, email, city, phone FROM users WHERE email = \\$1").
		WithArgs(expectedUser.Email).
		WillReturnRows(rows)

	user, err := repo.GetUserByEmail(expectedUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByEmail_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	email := "notfound@example.com"
	mock.ExpectQuery("SELECT id, name, email, city, phone FROM users WHERE email = \\$1").
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "city", "phone"}))

	user, err := repo.GetUserByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, domain.User{}, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	userID := uuid.New()
	user := domain.User{
		ID:    userID,
		Name:  "John Updated",
		Email: "john.updated@example.com",
		City:  "Campinas",
		Phone: "+5511988888888",
	}

	mock.ExpectExec("UPDATE users SET name = \\$1, email = \\$2, city = \\$3, phone = \\$4 WHERE id = \\$5").
		WithArgs(user.Name, user.Email, user.City, user.Phone, user.ID.String()).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	userID := uuid.New()

	mock.ExpectExec("DELETE FROM users WHERE id = \\$1").
		WithArgs(userID.String()).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteUser(userID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
