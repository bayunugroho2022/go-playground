package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}