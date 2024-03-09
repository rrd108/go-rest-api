package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrd108/go-rest-api/internal"
	"github.com/rrd108/go-rest-api/models"
	"gorm.io/gorm"
)

var db *gorm.DB
var allowedUnauthenticated = map[string][]string{"POST": {"/users/login"}, "GET": {"/products"}}

func init() {
	db = internal.ConnectDB()
}

func isAllowedUnauthenticated(method string, path string) bool {
	for _, url := range allowedUnauthenticated[method] {
		if url == path {
			return true
		}
	}
	return false
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if isAllowedUnauthenticated(c.Request.Method, c.Request.URL.Path) {
			c.Next()
			return
		}

		token := c.GetHeader("Token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authentication header"})
			c.Abort()
			return
		}

		var user models.User
		err := db.Table("users").Where("token = ?", token).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
				return
			}
			log.Fatal(err)
		}
		c.Set("user", user)

		c.Next()
	}
}
