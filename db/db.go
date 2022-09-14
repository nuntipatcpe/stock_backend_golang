package db

import (
	"main/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func SetupDB() {
	database, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Transaction{})
	db = database
}
