package handler

import (
	"context"
	"net/url"
	pb "price-chart/pkg/protobuf/goods"
	"price-chart/pkg/util"
	"price-chart/service/goods/internal/model/goods"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPC struct {
	Repo goods.Repository
}

func (g GRPC) Add(ctx context.Context, in *pb.AddRequest, out *pb.AddResponse) error {
	err := validateAdd(in)
	if err != nil {
		return err
	}

	ok, err := g.Repo.CheckByURL(ctx, in.Url)
	if err != nil {
		return err
	}
	if ok {
		return status.Error(codes.Code(409), "goods with such url already exists")
	}

	id := util.GenerateUUID().String()
	now := time.Now().UTC()

	model := goods.Goods{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		URL:       in.Url,
		Prices: []*goods.Price{
			{
				Value:     in.Price,
				CreatedAt: now,
			},
		},
	}
	err = g.Repo.Store(ctx, model)
	if err != nil {
		return err
	}
	out.GoodsId = id

	return nil
}

func (g GRPC) Update(ctx context.Context, in *pb.UpdateRequest) error {
	err := validateUpdate(in)
	if err != nil {
		return err
	}

	dto := &goods.UpdateDTO{}

	if in.Url != "" {
		ok, err := g.Repo.CheckAnotherByURL(ctx, in.GoodsId, in.Url)
		if err != nil {
			return err
		}
		if ok {
			return status.Error(codes.Code(409), "goods with such url already exists")
		}

		dto.URL = in.Url
	}
	if in.Status != "" {
		dto.Status = in.Status
	}
	if in.Price != nil {
		dto.Price = &in.Price.Value
	}

	err = g.Repo.Update(ctx, in.GoodsId, dto)
	if err != nil {
		return err
	}

	return nil
}

func (g GRPC) Remove(ctx context.Context, in *pb.RemoveRequest) error {
	err := validateRemove(in)
	if err != nil {
		return err
	}

	err = g.Repo.Delete(ctx, in.GoodsId)
	if err != nil {
		return err
	}

	return nil
}

func (g GRPC) Has(ctx context.Context, in *pb.HasRequest, out *pb.HasResponse) error {
	var (
		ok  bool
		err error
	)

	switch inst := in.Arg.(type) {
	case *pb.HasRequest_Id:
		if !util.IsValidUUID(inst.Id) {
			return status.Error(codes.Code(400), "request: arg.id must be uuid type")
		}
		ok, err = g.Repo.CheckByID(ctx, inst.Id)
	case *pb.HasRequest_Url:
		if _, err := url.ParseRequestURI(inst.Url); err != nil {
			return status.Error(codes.Code(400), "request: arg.url is invalid")
		}
		ok, err = g.Repo.CheckByURL(ctx, inst.Url)
	}
	if err != nil {
		return err
	}
	out.Ok = ok

	return nil
}

func (g GRPC) Goods(ctx context.Context, in *pb.GoodsRequest, out *pb.GoodsResponse) error {
	var (
		model *goods.Goods
		err   error
	)

	switch inst := in.Arg.(type) {
	case *pb.GoodsRequest_Id:
		if !util.IsValidUUID(inst.Id) {
			return status.Error(codes.Code(400), "request: arg.id must be uuid type")
		}
		model, err = g.Repo.LoadByID(ctx, inst.Id)
	case *pb.GoodsRequest_Url:
		if _, err := url.ParseRequestURI(inst.Url); err != nil {
			return status.Error(codes.Code(400), "request: arg.url is invalid")
		}
		model, err = g.Repo.LoadByURL(ctx, inst.Url)
	}
	if err != nil {
		return err
	}
	prices := make([]*pb.Price, 0, len(model.Prices))
	for _, p := range model.Prices {
		prices = append(prices, &pb.Price{
			Value:     p.Value,
			CreatedAt: p.CreatedAt.Unix(),
		})
	}

	out.Goods = &pb.Goods{
		Id:        model.ID,
		Url:       model.URL,
		Status:    model.Status,
		Prices:    prices,
		CreatedAt: model.CreatedAt.Unix(),
		UpdatedAt: model.UpdatedAt.Unix(),
	}

	return nil
}
