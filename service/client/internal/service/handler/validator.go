package handler

import (
	"net/url"
	pb "price-chart/pkg/protobuf/client"
	"price-chart/pkg/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateAdd(in *pb.AddRequest) error {
	if in.Username == "" {
		return status.Error(codes.Code(400), "request: username is empty")
	}
	if len(in.Username) < 3 {
		return status.Error(codes.Code(400), "request: username length shouldn't be less than 3")
	}
	if in.ChatId == 0 {
		return status.Error(codes.Code(400), "request: chat_id is 0")
	}
	if in.ChatType == "" {
		return status.Error(codes.Code(400), "request: chat_type is empty")
	}
	if in.FirstName == "" {
		return status.Error(codes.Code(400), "request: first_name is empty")
	}
	if len([]rune(in.FirstName)) < 3 {
		return status.Error(codes.Code(400), "request: first_name length shouldn't be less than 3")
	}
	if in.LastName != "" && len([]rune(in.LastName)) < 3 {
		return status.Error(codes.Code(400), "request: last_name length shouldn't be less than 3")
	}

	return nil
}

func validateRemove(in *pb.RemoveRequest) error {
	if !util.IsValidUUID(in.ClientId) {
		return status.Error(codes.Code(400), "request: client_id must be uuid type")
	}
	return nil
}

func validateUpdate(in *pb.UpdateRequest) error {
	if !util.IsValidUUID(in.ClientId) {
		return status.Error(codes.Code(400), "request: client_id must be uuid type")
	}
	if in.ChatId == 0 && in.ChatType == "" && in.FirstName == "" && in.LastName == "" && in.Username == "" {
		return status.Error(codes.Code(400), "request: nothing to update")
	}
	if in.FirstName != "" && len([]rune(in.FirstName)) < 3 {
		return status.Error(codes.Code(400), "request: first_name length shouldn't be less than 3")
	}
	if in.LastName != "" && len([]rune(in.LastName)) < 3 {
		return status.Error(codes.Code(400), "request: last_name length shouldn't be less than 3")
	}

	return nil
}

func validateAttachURL(in *pb.AttachUrlRequest) error {
	if !util.IsValidUUID(in.ClientId) {
		return status.Error(codes.Code(400), "request: client_id must be uuid type")
	}
	if _, err := url.ParseRequestURI(in.Url); err != nil {
		return status.Error(codes.Code(400), "request: url isn't valid")
	}
	return nil
}

func validateDetachURL(in *pb.DetachUrlRequest) error {
	if !util.IsValidUUID(in.ClientId) {
		return status.Error(codes.Code(400), "request: client_id must be uuid type")
	}
	if _, err := url.ParseRequestURI(in.Url); err != nil {
		return status.Error(codes.Code(400), "request: url isn't valid")
	}
	return nil
}
