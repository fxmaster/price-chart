package handler

import (
	"context"
	pb "price-chart/pkg/protobuf/parser"
	"price-chart/service/parser/internal/service/parser"
)

const Sel = ".item__price-once"

type GRPC struct {
	Parser parser.Parser
}

func (g GRPC) Price(ctx context.Context, in *pb.PriceRequest, out *pb.PriceResponse) error {

	resp, err := g.Parser.Parse(ctx, Sel, in.Url)
	if err != nil {
		return err
	}
	out.Price = resp

	return nil
}
