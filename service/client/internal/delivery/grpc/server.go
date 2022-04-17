package grpc

import (
	"context"
	pb "price-chart/pkg/protobuf/client"
)

type Server struct {
	pb.UnimplementedClientServiceServer
}

func (s Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{}, nil
}

func (s Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, nil
}

func (s Server) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, nil
}

func (s Server) Has(ctx context.Context, in *pb.HasRequest) (*pb.HasResponse, error) {
	return &pb.HasResponse{}, nil
}

func (s Server) Client(ctx context.Context, in *pb.ClientRequest) (*pb.ClientResponse, error) {
	return &pb.ClientResponse{}, nil
}

func (s Server) AttachUrl(ctx context.Context, in *pb.AttachUrlRequest) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, nil
}

func (s Server) DetachUrl(ctx context.Context, in *pb.DetachUrlRequest) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, nil
}
