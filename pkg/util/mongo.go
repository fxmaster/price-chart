package util

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo interface {
	Database() *mongo.Database
	Collection() *mongo.Collection
	Session() *mongo.Client
	HasCollection(ctx context.Context, val string) (bool, error)
}

type MongoClient struct {
	session  *mongo.Client
	db       *mongo.Database
	collName string
}

func NewMongoClient(ctx context.Context, dsn, db, collection string) (*MongoClient, error) {
	session, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	err = session.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		session:  session,
		db:       session.Database(db),
		collName: collection,
	}, nil
}

func (c MongoClient) Database() *mongo.Database {
	return c.db
}

func (c MongoClient) Collection() *mongo.Collection {
	return c.db.Collection(c.collName)
}

func (c MongoClient) Session() *mongo.Client {
	return c.session
}

func (c MongoClient) HasCollection(ctx context.Context, val string) (bool, error) {
	names, err := c.db.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return false, err
	}

	for _, name := range names {
		if name == val {
			return true, nil
		}
	}

	return false, nil
}
