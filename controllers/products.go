package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func ProductsList(c *gin.Context) {
	var products []Product
	// orm using built-in method
	db.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func ProductsAdd(c *gin.Context) {
	c.BindJSON(&product)
	// orm using built-in method
	db.Create(&product)

	c.JSON(http.StatusCreated, gin.H{"product": product})
}
