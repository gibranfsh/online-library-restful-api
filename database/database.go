package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3308)/go-restapi-fiber-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database.")
	}

	fmt.Println("Connection Opened to Database")
}