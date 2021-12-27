package db

import (
	"back_server/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("sync.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
