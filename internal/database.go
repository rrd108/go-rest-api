package internal

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// HARDCODED
var username = "root"
var password = "123"
var dbname = "rest-api"

func ConnectDB() *gorm.DB {
	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbname
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	return db
}
