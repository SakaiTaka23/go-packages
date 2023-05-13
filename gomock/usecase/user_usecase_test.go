package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gomock/mock"
	"gomock/model"
	"testing"
)

func Test_userUsecase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createdID := uuid.New().String()
	mockRepo := mock.NewMockUserRepo(ctrl)
	mockRepo.EXPECT().Create("John Doe").Return(createdID, nil)

	uUsecase := NewUserUsecase(mockRepo)
	id, err := uUsecase.CreateUser("John Doe")

	assert.NoError(t, err)
	assert.Equal(t, createdID, id)
}

func Test_userUsecase_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	findID := "4AF455DD-EE72-4412-B830-82D87187A467"
	mockRepo := mock.NewMockUserRepo(ctrl)
	mockRepo.EXPECT().GetUser(findID).Return(model.User{ID: findID, Name: "John Doe"}, nil)

	uUsecase := NewUserUsecase(mockRepo)
	user, err := uUsecase.GetUser(findID)

	assert.NoError(t, err)
	assert.Equal(t, findID, user.ID)
}
