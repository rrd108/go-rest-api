package initializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// HARDCODED
var username = "root"
var password = "123"
var dbname = "rest-api"

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbname
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}
}
