package api

import (
	"fmt"
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
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		var expiration float64
		expiration = claims["exp"].(float64)
		now := float64(time.Now().Unix())
		if now < expiration {
			email, ok := claims["email"].(string)
			if ok {
				c.Set("email", email)
				c.Next()
				return
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "expired token"})
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
}