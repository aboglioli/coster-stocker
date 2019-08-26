package db

import (
	"context"
	"log"

	"github.com/aboglioli/coster-stocker/helpers/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func Get() (*mongo.Database, error) {
	if database == nil {
		c := config.Get()

		clientOptions := options.Client().ApplyURI(c.MongoURL)

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		database = client.Database("stock")
	}

	return database, nil
}
