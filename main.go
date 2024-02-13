package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rrd108/go-rest-api/controllers"
	"github.com/rrd108/go-rest-api/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/products", controllers.ProductsList)
	router.POST("/users/login", controllers.UserLogin)
	router.GET("/users", controllers.UsersList)
	router.Run(":16108")
}
