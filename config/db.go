package config

import (
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)

	}
	db.AutoMigrate(&domain.User{})
	DB = db
}
