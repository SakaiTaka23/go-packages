package usecase

import (
	"errors"
	"github.com/google/uuid"
	"gomock/model"
	"gomock/repository"
)

type UserUsecase interface {
	CreateUser(name string) (string, error)
	GetUser(id string) (model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) CreateUser(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	return u.userRepo.Create(name)
}

func (u *userUsecase) GetUser(id string) (model.User, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return model.User{}, errors.New("invalid id")
	}
	return u.userRepo.GetUser(id)
}
