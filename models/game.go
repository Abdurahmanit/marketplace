package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	ImageURL    string             `bson:"image_url"`
}
