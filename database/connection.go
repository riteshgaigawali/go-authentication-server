package database

import (
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	conn, err := gorm.Open(mysql.Open("ritesh:ritesh@/test"), &gorm.Config{})
	if err != nil {
		panic("Could not connect...")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
