package api

import (
	"context"
	"github.com/chazapp/dante/pkg/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type bookForm struct {
	Title	string `json:"title"`
	Author	string	`json:"author"`
	Tags	[]string `json:"tags"`
}
func (api *API) CreateBook(c *gin.Context) {
	var form bookForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var book db.Book
	book.Title = form.Title
	book.Tags = form.Tags
	book.Author = form.Author
	res, err := api.db.Collection("books").InsertOne(
		context.Background(),
		book,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"book": res.InsertedID})
	return
}

func (api *API) GetBooks(c *gin.Context) {
	cur, err := api.db.Collection("books").Find(
		context.Background(),
		bson.D{},
		nil,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var books []db.Book
	for cur.Next(context.Background()) {
		var book db.Book
		err := cur.Decode(&book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
	return
}
