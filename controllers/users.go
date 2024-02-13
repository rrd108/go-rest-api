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

var db = initializers.GetDB()
var user User
var err error

func UserLogin(c *gin.Context) {
	var data LoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		// wrong JSON format, wrong data format
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate the MD5 hash of the user's password
	hash := md5.Sum([]byte(data.Password))
	hashedUserPassword := hex.EncodeToString(hash[:])

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

func UsersList(c *gin.Context) {
	// let's check if the request conatins a token header and if not reponse with unauthorized
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	log.Println(token)

	// let's check if the token is valid
	err = db.QueryRow("SELECT id FROM users WHERE token = ?", token).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		log.Fatal(err)
	}

	// let's get the list of users
	rows, err := db.Query("SELECT id, email,token FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//return users as a json array
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Token)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
