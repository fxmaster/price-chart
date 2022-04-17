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
		log.Println("create collection domains")
		err = client.Database().CreateCollection(ctx, collection)
		if err != nil {
			return errors.New("failed to create collection: " + err.Error())
		}

		_, err = client.Collection().Indexes().CreateMany(ctx, []mongo.IndexModel{
			{Keys: bson.D{{Key: "chat_id", Value: 1}}, Options: options.Index().SetSparse(true).SetUnique(true)},
			{Keys: bson.D{{Key: "username", Value: 1}}, Options: options.Index().SetSparse(true).SetUnique(true)},
			{Keys: bson.D{{Key: "phone", Value: 1}}, Options: options.Index().SetSparse(true).SetUnique(true)},
			{Keys: bson.D{{Key: "urls", Value: 1}}},
		})
		if err != nil {
			return errors.New("failed to create indexes: " + err.Error())
		}
	}

	return nil
}
