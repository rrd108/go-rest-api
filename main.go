package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rrd108/go-rest-api/controllers"
	"github.com/rrd108/go-rest-api/middlewares"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"OPTIONS", "GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Token", "Content-Type"},
		MaxAge:       300,
	}))
	router.Use(middlewares.TokenAuthMiddleware())

	router.GET("/products", controllers.ProductsList)
	router.POST("/products", controllers.ProductsAdd)
	router.PATCH("/products/:id", controllers.ProductsEdit)
	router.DELETE("/products/:id", controllers.ProductsDelete)

	router.POST("/users/login", controllers.UserLogin)
	router.GET("/users", controllers.UsersList)

	router.Run(":16108")
}
