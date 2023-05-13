package repository

import (
	"gomock/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(name string) (string, error)
	GetUser(id string) (model.User, error)
}

type userRepo struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepo {
	return &userRepo{Conn: conn}
}

func (u *userRepo) GetUser(id string) (model.User, error) {
	var user model.User
	err := u.Conn.Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *userRepo) Create(name string) (string, error) {
	user := model.User{Name: name}
	err := u.Conn.Create(&user).Error
	return user.ID, err
}
