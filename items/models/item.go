package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID        primitive.ObjectID
	Title     string  `bson:"title"`
	UserId    int     `bson:"userid"`
	Image     string  `bson:"image"`
	Currency  string  `bson:"currency"`
	Price     float32 `bson:"price"`
	Sale_sate int     `bson:"state"`
	Condition string  `bson:"condition"`
	Address   string  `bson:"address"`
}

type Items []Item
