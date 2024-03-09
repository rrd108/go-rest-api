package controllers

import (
	"github.com/rrd108/go-rest-api/internal"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = internal.ConnectDB()
}
