package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rrd108/go-rest-api/controllers"
)

func main() {
	router := gin.Default()
	router.POST("/users/login", controllers.UserLogin)
	router.GET("/users", controllers.UsersList)
	router.Run(":16108")
}
