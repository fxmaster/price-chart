package config

import (
	"context"
	"price-chart/pkg/util"
	"strconv"
	"sync"
	"time"
)

var (
	mongoClient *util.MongoClient
	mongoErr    error
	onceMongo   sync.Once
)

func MongoClient() (*util.MongoClient, error) {
	onceMongo.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		env := Environment().MongoParams
		dsn := "mongodb://" + env.User + ":" + env.Password + "@" + env.Host + ":" + strconv.Itoa(env.Port)

		mongoClient, mongoErr = util.NewMongoClient(ctx, dsn, env.DB, env.Collection)
	})

	return mongoClient, mongoErr
}
