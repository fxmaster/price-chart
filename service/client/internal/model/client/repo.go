package client

import (
	"context"
	"price-chart/pkg/util"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	KeySet         = "$set"
	KeyPush        = "$push"
	KeyPullAll     = "$pullAll"
	KeyPull        = "$pull"
	KeyEach        = "$each"
	KeyIn          = "$in"
	FieldID        = "_id"
	FieldChatID    = "chat_id"
	FieldFirstName = "first_name"
	FieldLastName  = "last_name"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
	FieldUsername  = "username"
	FieldChatType  = "chat_type"
	FieldUrls      = "urls"
)

type ClientRepository struct {
	Mongo util.Mongo
}

func (r ClientRepository) LoadByID(ctx context.Context, id string) (*Client, error) {
	var err error
	d := Client{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldID: id}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.Code(404), "client not found")
		}
		return nil, status.Error(codes.Code(500), err.Error())
	}

	return &d, nil
}

func (r ClientRepository) LoadByChatID(ctx context.Context, chatID uint64) (*Client, error) {
	var err error
	d := Client{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldChatID: chatID}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.Code(404), "client not found")
		}
		return nil, status.Error(codes.Code(500), err.Error())
	}

	return &d, nil
}

func (r ClientRepository) CheckByID(ctx context.Context, id string) (bool, error) {
	var err error
	d := Client{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldID: id}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, status.Error(codes.Code(500), err.Error())
	}

	return true, nil
}

func (r ClientRepository) CheckByUsername(ctx context.Context, username string) (bool, error) {
	var err error
	d := Client{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldUsername: username}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, status.Error(codes.Code(500), err.Error())
	}

	return true, nil
}

func (r ClientRepository) CheckByChatID(ctx context.Context, chatID uint64) (bool, error) {
	var err error
	d := Client{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldChatID: chatID}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, status.Error(codes.Code(500), err.Error())
	}

	return true, nil
}

func (r ClientRepository) Store(ctx context.Context, doc Client) error {
	if doc.areRequiredFieldsEmpty() {
		return status.Error(codes.Code(409), "required fields are empty")
	}

	data := bson.D{
		bson.E{Key: FieldID, Value: doc.ID},
		bson.E{Key: FieldCreatedAt, Value: primitive.NewDateTimeFromTime(doc.CreatedAt.UTC())},
		bson.E{Key: FieldUpdatedAt, Value: primitive.NewDateTimeFromTime(doc.UpdatedAt.UTC())},
		bson.E{Key: FieldChatID, Value: doc.ChatID},
		bson.E{Key: FieldChatType, Value: doc.ChatType},
		bson.E{Key: FieldFirstName, Value: doc.FirstName},
		bson.E{Key: FieldLastName, Value: doc.LastName},
		bson.E{Key: FieldUsername, Value: doc.Username},
	}

	if len(doc.Urls) != 0 {
		data = append(data, bson.E{Key: FieldUrls, Value: doc.Urls})
	}

	_, err := r.Mongo.Collection().InsertOne(ctx, data)
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (r ClientRepository) Delete(ctx context.Context, id string) error {
	doc, err := r.LoadByID(ctx, id)
	if err != nil {
		return err
	}

	res, err := r.Mongo.Collection().DeleteOne(ctx, bson.M{FieldID: doc.ID})
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}
	if res.DeletedCount != 1 {
		return status.Error(codes.Code(500), "DeletedCount is not 1:"+strconv.FormatInt(res.DeletedCount, 10))
	}

	return nil
}

func (r ClientRepository) AttachURL(ctx context.Context, id, url string) error {
	doc, err := r.LoadByID(ctx, id)
	if err != nil {
		return err
	}

	for i := range doc.Urls {
		if doc.Urls[i] == url {
			return status.Error(codes.Code(409), "client already has this url")
		}
	}

	_, err = r.Mongo.Collection().UpdateOne(ctx, bson.M{FieldID: doc.ID}, bson.M{
		KeyPush: bson.M{FieldUrls: url},
	})
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (r ClientRepository) DetachURL(ctx context.Context, id, url string) error {
	doc, err := r.LoadByID(ctx, id)
	if err != nil {
		return err
	}

	var ok bool
	for i := range doc.Urls {
		if doc.Urls[i] == url {
			ok = true
		}
	}
	if !ok {
		return status.Error(codes.Code(409), "client doesn't have this url")
	}

	_, err = r.Mongo.Collection().UpdateOne(ctx, bson.M{FieldID: doc.ID}, bson.M{
		KeyPull: bson.M{FieldUrls: url},
	})
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (r ClientRepository) Update(ctx context.Context, id string, dto *UpdateDTO) error {
	query := bson.D{}

	doc, err := r.LoadByID(ctx, id)
	if err != nil {
		return err
	}

	if dto.ChatID != 0 {
		if dto.ChatID == doc.ChatID {
			return status.Error(codes.Code(409), "client's chat_id already same")
		}
		query = append(query, bson.E{Key: KeySet, Value: bson.D{
			bson.E{Key: FieldChatID, Value: dto.ChatID}},
		})
	}
	if dto.FirstName != "" {
		if dto.FirstName == doc.FirstName {
			return status.Error(codes.Code(409), "client's first_name already same")
		}
		query = append(query, bson.E{Key: KeySet, Value: bson.D{
			bson.E{Key: FieldFirstName, Value: dto.FirstName}},
		})
	}
	if dto.LastName != "" {
		if dto.LastName == doc.LastName {
			return status.Error(codes.Code(409), "client's last_name already same")
		}
		query = append(query, bson.E{Key: KeySet, Value: bson.D{
			bson.E{Key: FieldLastName, Value: dto.LastName}},
		})
	}
	if dto.Username != "" {
		if dto.Username == doc.Username {
			return status.Error(codes.Code(409), "client's username already same")
		}
		query = append(query, bson.E{Key: KeySet, Value: bson.D{
			bson.E{Key: FieldUsername, Value: dto.Username}},
		})
	}
	if dto.ChatType != "" {
		if dto.ChatType == doc.ChatType {
			return status.Error(codes.Code(409), "client's chat_type already same")
		}
		query = append(query, bson.E{Key: KeySet, Value: bson.D{
			bson.E{Key: FieldChatType, Value: dto.ChatType}},
		})
	}

	if len(query) == 0 {
		return status.Error(codes.Code(409), "there is nothing to update")
	}

	query = append(query, bson.E{Key: KeySet, Value: bson.D{
		bson.E{Key: FieldUpdatedAt, Value: primitive.NewDateTimeFromTime(time.Now().UTC())}},
	})

	_, err = r.Mongo.Collection().UpdateOne(ctx, bson.M{FieldID: doc.ID}, query)
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}
