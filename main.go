package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rrd108/go-rest-api/controllers"
	"github.com/rrd108/go-rest-api/middlewares"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.TokenAuthMiddleware())

	router.GET("/products", controllers.ProductsList)
	router.POST("/users/login", controllers.UserLogin)
	router.GET("/users", controllers.UsersList)

	router.Run(":16108")
}
