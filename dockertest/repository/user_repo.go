package repository

import (
	"dockertest/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(name string) (uint, error)
	Delete(id uint) error
	FindAll() ([]model.User, error)
	UpdateName(id uint, name string) error
}

type sqlUser struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &sqlUser{Conn: conn}
}

func (m *sqlUser) Create(name string) (uint, error) {
	user := model.User{Name: name}
	err := m.Conn.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (m *sqlUser) Delete(id uint) error {
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

func (m *sqlUser) UpdateName(id uint, name string) error {
	err := m.Conn.Model(&model.User{}).Where("id = ?", id).Update("name", name).Error
	if err != nil {
		return err
	}
	return nil
}
