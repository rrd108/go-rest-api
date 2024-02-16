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

func ProductsList(c *gin.Context) {
	var products []Product
	// orm using built-in method
	db.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func ProductsAdd(c *gin.Context) {
	var product Product
	c.BindJSON(&product)
	// orm using built-in method
	db.Create(&product)

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func ProductsEdit(c *gin.Context) {
	var product Product
	id := c.Param("id")
	// orm using built-in method
	db.First(&product, id)

	c.BindJSON(&product)
	// orm using built-in method
	db.Save(&product)

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func ProductsDelete(c *gin.Context) {
	var product Product
	id := c.Param("id")
	// orm using built-in method
	db.First(&product, id)
	// orm using built-in method
	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"deleted": id})
}
