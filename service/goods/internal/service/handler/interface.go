package handler

import (
	"context"
	pb "price-chart/pkg/protobuf/goods"
)

type Handler interface {
	Add(ctx context.Context, in *pb.AddRequest, out *pb.AddResponse) error
	Update(ctx context.Context, in *pb.UpdateRequest) error
	Remove(ctx context.Context, in *pb.RemoveRequest) error
	Has(ctx context.Context, in *pb.HasRequest, out *pb.HasResponse) error
	Goods(ctx context.Context, in *pb.GoodsRequest, out *pb.GoodsResponse) error
}
