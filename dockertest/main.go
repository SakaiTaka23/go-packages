package main

import (
	"dockertest/model"
	"dockertest/repository"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:pass@tcp(127.0.0.1:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = db.AutoMigrate(&model.User{})

	repo := repository.NewUserRepository(db)
	id, _ := repo.Create("John Doe")
	fmt.Printf("user id: %s created\n", id)
	users, _ := repo.FindAll()
	fmt.Printf("%+v\n", users)
}
