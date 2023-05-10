package main

import (
	"fmt"
	"go-sqlmock/model"
	"go-sqlmock/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("example.db"), &gorm.Config{})
	_ = db.AutoMigrate(&model.User{})

	repo := repository.NewUserRepository(db)
	id, _ := repo.Create(model.User{ID: "1", Name: "John Doe"})
	fmt.Printf("user id: %s created\n", id)
	users, _ := repo.FindAll()
	fmt.Printf("%+v\n", users)
}
