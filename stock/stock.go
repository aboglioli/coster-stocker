package stock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Stock struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name" validate:"required"`
}

func NewStock(name string) *Stock {
	return &Stock{
		ID:   primitive.NewObjectID(),
		Name: name,
	}
}
