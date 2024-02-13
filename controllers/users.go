package controllers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrd108/go-rest-api/initializers"
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
	if err := c.ShouldBindJSON(&data); err != nil {
		// wrong JSON format, wrong data format
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := initializers.GetDB()
	defer db.Close()

	// Calculate the MD5 hash of the user's password
	hash := md5.Sum([]byte(data.Password))
	hashedUserPassword := hex.EncodeToString(hash[:])

	var user User
	var err error
	err = db.QueryRow("SELECT id, email, token FROM users WHERE email = ? AND password = ?", data.Email, hashedUserPassword).Scan(&user.ID, &user.Email, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
