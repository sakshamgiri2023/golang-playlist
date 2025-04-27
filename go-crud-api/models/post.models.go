package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"` // MongoDB ObjectID
	Title string             `bson:"title"`
	Body  string             `bson:"body"`
}
