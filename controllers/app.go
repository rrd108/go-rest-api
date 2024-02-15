package controllers

import (
	"github.com/rrd108/go-rest-api/internal"
	"github.com/rrd108/go-rest-api/models"
	"gorm.io/gorm"
)

var user models.User
var db *gorm.DB

func init() {
	db = internal.ConnectDB()
}
