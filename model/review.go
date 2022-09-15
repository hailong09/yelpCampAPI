package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Body   string `json:"body" bson:"body"`
	Rating int    `json:"rating" bson:"rating"`
	Author primitive.ObjectID `json:"author" bson:"author"`
}