package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func CreateMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()

	conn, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	return conn, mock
}
