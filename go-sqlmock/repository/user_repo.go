package repository

import (
	"go-sqlmock/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User) (string, error)
	Delete(id string) error
	FindAll() ([]model.User, error)
	UpdateName(id string, name string) error
}

type sqlUser struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &sqlUser{Conn: conn}
}

func (m *sqlUser) Create(user model.User) (string, error) {
	err := m.Conn.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (m *sqlUser) Delete(id string) error {
	err := m.Conn.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *sqlUser) FindAll() ([]model.User, error) {
	var users []model.User
	err := m.Conn.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *sqlUser) UpdateName(id string, name string) error {
	err := m.Conn.Model(&model.User{}).Where("id = ?", id).Update("name", name).Error
	if err != nil {
		return err
	}
	return nil
}
