package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/codingsluv/book-store/POS-api/models"
)

var DB *gorm.DB

func ConnectDB() {
	// import "gorm.io/driver/mysql"
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "root:codingsluv@tcp(127.0.0.1:3306)/pos-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db

	// Migrate the schema
	DB.AutoMigrate(&models.Cashier{})
}
