package usecase

import (
	"fmt"
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

func Test_userUsecase_CreateUser1(t *testing.T) {
	type fields struct {
		userRepo *mock.MockUserRepo
	}
	type args struct {
		name string
	}
	createdID := uuid.New().String()
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.userRepo.EXPECT().Create("John Doe").Return(createdID, nil)
			},
			args:    args{name: "John Doe"},
			want:    createdID,
			wantErr: assert.NoError,
		},
		{
			name:    "name must not be empty",
			prepare: nil,
			args:    args{name: ""},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{userRepo: mock.NewMockUserRepo(ctrl)}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			u := &userUsecase{
				userRepo: f.userRepo,
			}

			got, err := u.CreateUser(tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateUser(%v)", tt.args.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CreateUser(%v)", tt.args.name)
		})
	}
}
