package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrd108/go-rest-api/models"
	"gorm.io/gorm"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

	currentUser, exists := c.Get("user")
	if exists {
		user, ok := currentUser.(models.User)
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	}

	// orm using table
	var user models.User
	err = db.Table("users").Where("email = ? AND password = ?", data.Email, hashedUserPassword).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		log.Fatal(err)
	}
	if user.ID > 0 {
		c.Set("user", user)
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UsersList(c *gin.Context) {
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}
