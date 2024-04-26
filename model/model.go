package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID        		primitive.ObjectID `bson:"_id,omitempty"`
	Description 	string 						 `json:"description"`
	Amount 			  float64             `json:"amount"`
}