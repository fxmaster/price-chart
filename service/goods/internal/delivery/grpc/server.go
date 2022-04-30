package grpc

import (
	"context"
	pb "price-chart/pkg/protobuf/goods"
	"price-chart/service/goods/internal/service/handler"
)

type Server struct {
	pb.UnimplementedGoodsServiceServer
	Handler handler.Handler
}

func (s Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	out := &pb.AddResponse{}
	err := s.Handler.Add(ctx, in, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (s Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.EmptyResponse, error) {
	err := s.Handler.Update(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}

func (s Server) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.EmptyResponse, error) {
	err := s.Handler.Remove(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}

func (s Server) Has(ctx context.Context, in *pb.HasRequest) (*pb.HasResponse, error) {
	out := &pb.HasResponse{}
	err := s.Handler.Has(ctx, in, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (s Server) Goods(ctx context.Context, in *pb.GoodsRequest) (*pb.GoodsResponse, error) {
	out := &pb.GoodsResponse{}
	err := s.Handler.Goods(ctx, in, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
