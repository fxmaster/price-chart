package app

import (
	"context"
	"price-chart/service/parser/internal/config"
	"price-chart/service/parser/internal/delivery/grpc"
	"strconv"
)

func Run(ctx context.Context) {
	env := config.Environment()

	grpc.Init(ctx, "0.0.0.0:"+strconv.Itoa(env.App.Port))
}
