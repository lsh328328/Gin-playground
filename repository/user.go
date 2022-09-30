package repository

import (
	"github.com/Gin-playground/database"
	"github.com/Gin-playground/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	db := database.InitDB()
	db.AutoMigrate(&model.User{})

	return &UserRepository{DB: db}
}
