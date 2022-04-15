package handler

import (
	"context"
	pb "price-chart/pkg/protobuf/parser"
)

type Handler interface {
	Price(ctx context.Context, in *pb.PriceRequest, out *pb.PriceResponse) error
}
