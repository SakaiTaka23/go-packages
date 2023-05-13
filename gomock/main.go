package main

import (
	"fmt"
	"gomock/model"
	"gomock/repository"
	"gomock/usecase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	_ = db.AutoMigrate(&model.User{})

	uRepo := repository.NewUserRepository(db)
	uUcs := usecase.NewUserUsecase(uRepo)

	id, _ := uUcs.CreateUser("John Doe")
	fmt.Printf("created user id: %s\n", id)
	u, _ := uUcs.GetUser(id)
	fmt.Printf("user: %+v\n", u)
}
