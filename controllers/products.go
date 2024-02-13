package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrd108/go-rest-api/initializers"
)

type Product struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
}

var product Product

func init() {
	initializers.ConnectDB()
}

func ProductsList(c *gin.Context) {
	var products []Product
	// orm using built-in method
	initializers.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}
