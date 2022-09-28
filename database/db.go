package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	envError := godotenv.Load("./../.env")

	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	var DB_USERNAME = os.Getenv("DB_USERNAME")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_PORT = os.Getenv("DB_PORT")

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database: error = %v", err)
		return nil
	}

	fmt.Println("Database connected...")
	return db
}
