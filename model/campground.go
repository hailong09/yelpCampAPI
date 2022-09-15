package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Campground struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Geometry    [2]float32         `json:"geometry,omitempty" bson:"geometry,omitempty"`
	Price       float32            `json:"price,omitempty" bson:"price,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Location    string             `json:"location,omitempty" bson:"location,omitempty"`
	Author      primitive.ObjectID `json:"author,omitempty" bson:"author,omitempty"`
	Reviews []primitive.ObjectID   `json:"reviews,omitempty" bson:"reviews,omitempty"`
}
