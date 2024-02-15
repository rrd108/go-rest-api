package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func UserLogin(c *gin.Context) {
	var data LoginData
	var err error
	if err = c.ShouldBindJSON(&data); err != nil {
		// wrong JSON format, wrong data format
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate the MD5 hash of the user's password
	hash := md5.Sum([]byte(data.Password))
	hashedUserPassword := hex.EncodeToString(hash[:])

	// orm using table
	err = db.Table("users").Where("email = ? AND password = ?", data.Email, hashedUserPassword).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UsersList(c *gin.Context) {
	var err error

	// let's check if the request conatins a token header and if not reponse with unauthorized
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// let's check if the token is valid
	// orm using table
	err = db.Table("users").Where("token = ?", token).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		log.Fatal(err)
	}

	// let's get the list of users
	var users []User
	// orm using built-in method
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}
