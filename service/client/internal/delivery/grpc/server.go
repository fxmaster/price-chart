package grpc

import (
	"context"
	pb "price-chart/pkg/protobuf/client"
	"price-chart/service/client/internal/service/handler"
)

type Server struct {
	pb.UnimplementedClientServiceServer
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

func (s Server) Client(ctx context.Context, in *pb.ClientRequest) (*pb.ClientResponse, error) {
	out := &pb.ClientResponse{}
	err := s.Handler.Client(ctx, in, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (s Server) AttachUrl(ctx context.Context, in *pb.AttachUrlRequest) (*pb.EmptyResponse, error) {
	err := s.Handler.AttachURL(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}

func (s Server) DetachUrl(ctx context.Context, in *pb.DetachUrlRequest) (*pb.EmptyResponse, error) {
	err := s.Handler.DetachURL(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}
