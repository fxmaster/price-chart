package app

import (
	"context"
	"log"
	"price-chart/service/client/internal/config"
	"price-chart/service/client/internal/delivery/grpc"
	"strconv"
)

func Run(ctx context.Context) {
	env := config.Environment()

	mongo, err := config.MongoClient()
	if err != nil {
		log.Fatalln("failed to get mongoClient: " + err.Error())
	}

	grpc.Init(ctx, "0.0.0.0:"+strconv.Itoa(env.App.Port), mongo)
}
