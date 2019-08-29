package stock

import (
	"context"

	"github.com/aboglioli/coster-stocker/helpers/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Insert(stock *Stock) (*Stock, error)
	// Update(stock *Stock) (*Stock, error)
	// FindAll() ([]*Stock, error)
	FindByID(id string) (*Stock, error)
}

type repositoryImpl struct {
	collection *mongo.Collection
}

func newRepository() (Repository, error) {
	db, err := db.Get()

	if err != nil {
		return nil, err
	}

	return &repositoryImpl{
		collection: db.Collection("stock"),
	}, nil
}

func (r *repositoryImpl) Insert(stock *Stock) (*Stock, error) {
	_, err := r.collection.InsertOne(context.Background(), stock)
	if err != nil {
		return nil, err
	}

	return stock, nil
}

func (r repositoryImpl) FindByID(id string) (*Stock, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	stock := &Stock{}
	filter := bson.D{
		{"_id", objID},
	}

	if err := r.collection.FindOne(context.Background(), filter).Decode(stock); err != nil {
		return nil, err
	}

	return stock, nil
}
