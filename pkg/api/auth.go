package api

import (
	"context"
	"github.com/chazapp/dante/pkg/api/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type authForm struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type registerForm struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (api *API) Register(c *gin.Context) {
	var form registerForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filter := bson.D{{"email", form.Email}}
	result := api.db.Collection("users").FindOne(context.Background(), filter)
	var user db.User
	err := result.Decode(&user)
	if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusConflict, gin.H{"error": "account already exists."})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
		return
	}
	var account db.User
	account.Email = form.Email
	account.Password = string(hashedPassword)
	_, err = api.db.Collection("users").InsertOne(context.Background(), account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return JWT
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	at, err := accessToken.SignedString([]byte(os.Getenv("JWT_HMAC_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": at})
}

func (api *API) Login(c *gin.Context) {
	var form authForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filter := bson.D{{"email", form.Email}}
	result := api.db.Collection("users").FindOne(context.Background(), filter)
	var user db.User
	err := result.Decode(&user)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "bad credentials"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "bad credentials"})
		return
	}
	// Return JWT
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	at, err := accessToken.SignedString([]byte(os.Getenv("JWT_HMAC_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": at})
}
