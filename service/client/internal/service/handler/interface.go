package handler

import (
	"context"
	pb "price-chart/pkg/protobuf/client"
)

type Handler interface {
	Add(ctx context.Context, in *pb.AddRequest, out *pb.AddResponse) error
	Update(ctx context.Context, in *pb.UpdateRequest) error
	Remove(ctx context.Context, in *pb.RemoveRequest) error
	Has(ctx context.Context, in *pb.HasRequest, out *pb.HasResponse) error
	Client(ctx context.Context, in *pb.ClientRequest, out *pb.ClientResponse) error
	AttachURL(ctx context.Context, in *pb.AttachUrlRequest) error
	DetachURL(ctx context.Context, in *pb.DetachUrlRequest) error
}
