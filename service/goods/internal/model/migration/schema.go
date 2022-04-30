package migration

import (
	"context"
	"errors"
	"log"
	"price-chart/pkg/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateCollection(ctx context.Context, client util.Mongo, collection string) error {
	has, err := client.HasCollection(ctx, collection)
	if err != nil {
		return errors.New("failed to check if collection exists: " + err.Error())
	}
	if !has {
		log.Println("create collection goods")
		err = client.Database().CreateCollection(ctx, collection)
		if err != nil {
			return errors.New("failed to create collection: " + err.Error())
		}

		_, err = client.Collection().Indexes().CreateMany(ctx, []mongo.IndexModel{
			{Keys: bson.D{{Key: "url", Value: 1}}, Options: options.Index().SetSparse(true).SetUnique(true)},
			{Keys: bson.D{{Key: "prices", Value: 1}}},
		})
		if err != nil {
			return errors.New("failed to create indexes: " + err.Error())
		}
	}

	return nil
}
