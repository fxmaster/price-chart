package handler

import (
	"context"
	pb "price-chart/pkg/protobuf/parser"
)

type Parser struct{}

func (Parser) Price(ctx context.Context, in *pb.PriceRequest, out *pb.PriceResponse) error {
	return nil
}
