package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func (c* gin.Context) {
		if c.Request.RequestURI == "/login" || c.Request.RequestURI == "/register" {
			c.Next()
			return
		}
		bearer := c.GetHeader("Authorization")
		tokenString := strings.Replace(bearer, "Bearer ", "", -1)
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_HMAC_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		if exp, ok := claims["exp"].(int64); ok && exp < time.Now().Unix() {
			email, ok := claims["email"].(int64)
			if ok {
				c.Set("email", email)
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "expired token"})
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Next()
	}
}