package handler

import (
	"context"
	pb "price-chart/pkg/protobuf/client"
	"price-chart/pkg/util"
	"price-chart/service/client/internal/model/client"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPC struct {
	Repo client.Repository
}

func (g GRPC) Add(ctx context.Context, in *pb.AddRequest, out *pb.AddResponse) error {
	err := validateAdd(in)
	if err != nil {
		return err
	}

	ok, err := g.Repo.CheckByChatID(ctx, in.ChatId)
	if err != nil {
		return err
	}
	if ok {
		return status.Error(codes.Code(409), "client with such chat_id already exists")
	}

	ok, err = g.Repo.CheckByUsername(ctx, in.Username)
	if err != nil {
		return err
	}
	if ok {
		return status.Error(codes.Code(409), "client with such username already exists")
	}

	id := util.GenerateUUID().String()
	now := time.Now().UTC()

	model := client.Client{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		ChatID:    in.ChatId,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		ChatType:  in.ChatType,
	}
	err = g.Repo.Store(ctx, model)
	if err != nil {
		return err
	}

	out.ClientId = id

	return nil
}

func (g GRPC) Update(ctx context.Context, in *pb.UpdateRequest) error {
	err := validateUpdate(in)
	if err != nil {
		return err
	}

	dto := &client.UpdateDTO{}

	if in.ChatId != 0 {
		ok, err := g.Repo.CheckByChatID(ctx, in.ChatId)
		if err != nil {
			return err
		}
		if ok {
			return status.Error(codes.Code(409), "client with such chat_id already exists")
		}

		dto.ChatID = in.ChatId
	}
	if in.Username != "" {
		ok, err := g.Repo.CheckByUsername(ctx, in.Username)
		if err != nil {
			return err
		}
		if ok {
			return status.Error(codes.Code(409), "client with such username already exists")
		}

		dto.Username = in.Username
	}
	if in.ChatType != "" {
		dto.ChatType = in.ChatType
	}
	if in.FirstName != "" {
		dto.FirstName = in.FirstName
	}
	if in.LastName != "" {
		dto.LastName = in.LastName
	}

	err = g.Repo.Update(ctx, in.ClientId, dto)
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

	err = g.Repo.Delete(ctx, in.ClientId)
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

	err = validateHas(in)
	if err != nil {
		return err
	}

	switch inst := in.Arg.(type) {
	case *pb.HasRequest_Id:
		ok, err = g.Repo.CheckByID(ctx, inst.Id)
	case *pb.HasRequest_ChatId:
		ok, err = g.Repo.CheckByChatID(ctx, inst.ChatId)
	}
	if err != nil {
		return err
	}
	out.Ok = ok

	return nil
}

func (g GRPC) Client(ctx context.Context, in *pb.ClientRequest, out *pb.ClientResponse) error {
	var (
		model *client.Client
		err   error
	)

	err = validateClient(in)
	if err != nil {
		return err
	}

	switch inst := in.Arg.(type) {
	case *pb.ClientRequest_Id:
		model, err = g.Repo.LoadByID(ctx, inst.Id)
	case *pb.ClientRequest_ChatId:
		model, err = g.Repo.LoadByChatID(ctx, inst.ChatId)
	}
	if err != nil {
		return err
	}

	out.Client = &pb.Client{
		Id:        model.ID,
		ChatId:    model.ChatID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Username:  model.Username,
		ChatType:  model.ChatType,
		Urls:      model.Urls,
		CreatedAt: model.CreatedAt.Unix(),
		UpdatedAt: model.UpdatedAt.Unix(),
	}

	return nil
}

func (g GRPC) AttachUrl(ctx context.Context, in *pb.AttachUrlRequest) error {
	err := validateAttachUrl(in)
	if err != nil {
		return err
	}

	return g.Repo.AttachUrl(ctx, in.ClientId, in.Url)
}

func (g GRPC) DetachUrl(ctx context.Context, in *pb.DetachUrlRequest) error {
	err := validateDetachUrl(in)
	if err != nil {
		return err
	}

	return g.Repo.DetachUrl(ctx, in.ClientId, in.Url)
}
