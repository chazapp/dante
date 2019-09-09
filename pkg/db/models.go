package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Email string	`json:"email" bson:"email"`
	Password string	`json:"password" bson:"password"`
}

type Book struct {
	Title	string	`json:"title" bson:"title"`
	Author  string	`json:"author" bson:"author"`
	Tags	[]string 	`json:"tags" bson:"tags"`
}

type Quote struct {
	Book	primitive.ObjectID
}