package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rrd108/go-rest-api/controllers"
)

func main() {
	router := gin.Default()
	router.POST("/users/login", controllers.UserLogin)
	router.Run(":16108")
}
