package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	UserID primitive.ObjectID   `bson:"user_id"`
	Games  []primitive.ObjectID `bson:"games"`
}
