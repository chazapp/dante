package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type API struct {
	db *mongo.Database
}

func StartRESTAPI(port int, db string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db))
	if err != nil {
		return err
	}
	router := gin.Default()
	addr := ":" + strconv.Itoa(port)
	api := &API{db: client.Database("dante")}
	router.Use(JWTAuthMiddleware())
	router.POST("/register", api.Register)
	router.POST("/login", api.Login)
	router.POST("/book", api.CreateBook)
	router.GET("/books", api.GetBooks)
	err = router.Run(addr)
	return err
}
