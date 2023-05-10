package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-sqlmock/mock"
	"go-sqlmock/model"
	"regexp"
	"testing"
)

func Test_sqlUser_Create(t *testing.T) {
	conn, mockDB := mock.CreateMockDB()

	id := "1"
	name := "John"

	mockDB.ExpectBegin()
	mockDB.
		ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`name`) VALUES (?,?)")).
		WithArgs(id, name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	repo := NewUserRepository(conn)
	userID, err := repo.Create(model.User{ID: id, Name: name})

	expectations := mockDB.ExpectationsWereMet()

	assert.Nil(t, expectations)
	assert.Equal(t, nil, err)
	assert.Equal(t, id, userID)
}

func Test_sqlUser_Delete(t *testing.T) {
	conn, mockDB := mock.CreateMockDB()

	id := "1"

	mockDB.ExpectBegin()
	mockDB.
		ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE `users`.`id` = ?")).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	repo := NewUserRepository(conn)
	err := repo.Delete(id)

	expectations := mockDB.ExpectationsWereMet()

	assert.Nil(t, expectations)
	assert.Nil(t, err)
}

func Test_sqlUser_FindAll(t *testing.T) {
	conn, mockDB := mock.CreateMockDB()
	mockDB.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("1", "John").AddRow("2", "Doe"))

	repo := NewUserRepository(conn)
	users, err := repo.FindAll()

	expectations := mockDB.ExpectationsWereMet()

	assert.Nil(t, expectations)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "1", users[0].ID)
	assert.Equal(t, "John", users[0].Name)
}

func Test_sqlUser_UpdateName(t *testing.T) {
	conn, mockDB := mock.CreateMockDB()

	id := "1"
	name := "John"

	mockDB.ExpectBegin()
	mockDB.
		ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `name`=? WHERE id = ?")).
		WithArgs(name, id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	repo := NewUserRepository(conn)
	err := repo.UpdateName(id, name)

	expectations := mockDB.ExpectationsWereMet()

	assert.Nil(t, expectations)
	assert.Nil(t, err)
}
