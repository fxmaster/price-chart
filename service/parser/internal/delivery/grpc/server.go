package grpc

import (
	"context"
	pb "price-chart/pkg/protobuf/parser"
	"price-chart/service/parser/internal/service/handler"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedParserServiceServer
	Handler handler.Handler
}

func (s Server) Price(ctx context.Context, in *pb.PriceRequest) (*pb.PriceResponse, error) {

	err := validatePrice(in)
	if err != nil {
		return nil, err
	}
	out := &pb.PriceResponse{}

	err = s.Handler.Price(ctx, in, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func validatePrice(in *pb.PriceRequest) error {
	if in.Url == "" {
		return status.Error(codes.Code(400), "request: url is empty")
	}

	return nil
}
