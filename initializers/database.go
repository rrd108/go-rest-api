package initializers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// HARDCODED
var username = "root"
var password = "123"
var dbname = "rest-api"

func init() {
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")
}

func GetDB() *sql.DB {
	return db
}
