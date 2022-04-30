package goods

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
	KeySet              = "$set"
	KeyNE               = "$ne"
	KeyPush             = "$push"
	KeyPullAll          = "$pullAll"
	KeyPull             = "$pull"
	KeyEach             = "$each"
	KeyIn               = "$in"
	FieldID             = "_id"
	FieldURL            = "url"
	FieldStatus         = "status"
	FieldPrices         = "prices"
	FieldCreatedAt      = "created_at"
	FieldUpdatedAt      = "updated_at"
	FieldPriceValue     = "prices.value"
	FieldPriceCreatedAt = "prices.created_at"
)

type GoodsRepository struct {
	Mongo util.Mongo
}

func (r GoodsRepository) LoadByID(ctx context.Context, id string) (*Goods, error) {
	var err error
	d := Goods{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldID: id}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.Code(404), "goods not found")
		}
		return nil, status.Error(codes.Code(500), err.Error())
	}

	return &d, nil
}

func (r GoodsRepository) LoadByURL(ctx context.Context, url string) (*Goods, error) {
	var err error
	d := Goods{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldURL: url}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.Code(404), "goods not found")
		}
		return nil, status.Error(codes.Code(500), err.Error())
	}

	return &d, nil
}

func (r GoodsRepository) CheckByID(ctx context.Context, id string) (bool, error) {
	var err error
	d := Goods{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldID: id}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, status.Error(codes.Code(500), err.Error())
	}

	return true, nil
}

func (r GoodsRepository) CheckByURL(ctx context.Context, url string) (bool, error) {
	var err error
	d := Goods{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldURL: url}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, status.Error(codes.Code(500), err.Error())
	}

	return true, nil
}

func (r GoodsRepository) CheckAnotherByURL(ctx context.Context, id, url string) (bool, error) {
	var err error
	d := Goods{}

	err = r.Mongo.Collection().FindOne(ctx, bson.M{FieldID: bson.M{KeyNE: id}, FieldURL: url}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, status.Error(codes.Code(500), err.Error())
	}

	return true, nil
}

func (r GoodsRepository) Store(ctx context.Context, doc Goods) error {
	if doc.areRequiredFieldsEmpty() {
		return status.Error(codes.Code(409), "required fields are empty")
	}

	data := bson.D{
		bson.E{Key: FieldID, Value: doc.ID},
		bson.E{Key: FieldCreatedAt, Value: primitive.NewDateTimeFromTime(doc.CreatedAt.UTC())},
		bson.E{Key: FieldUpdatedAt, Value: primitive.NewDateTimeFromTime(doc.UpdatedAt.UTC())},
		bson.E{Key: FieldURL, Value: doc.URL},
		bson.E{Key: FieldStatus, Value: doc.Status},
		bson.E{Key: FieldPrices, Value: doc.Prices},
	}

	_, err := r.Mongo.Collection().InsertOne(ctx, data)
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (r GoodsRepository) Delete(ctx context.Context, id string) error {
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

func (r GoodsRepository) Update(ctx context.Context, id string, dto *UpdateDTO) error {
	query := bson.D{}

	doc, err := r.LoadByID(ctx, id)
	if err != nil {
		return err
	}

	if dto.URL != "" {
		if dto.URL == doc.URL {
			return status.Error(codes.Code(409), "goods' url already same")
		}
		query = append(query, bson.E{Key: KeySet, Value: bson.D{
			bson.E{Key: FieldURL, Value: dto.URL}},
		})
	}
	if dto.Status != "" {
		if dto.Status == doc.Status {
			return status.Error(codes.Code(409), "goods' status already same")
		}
		query = append(query, bson.E{Key: KeySet, Value: bson.D{
			bson.E{Key: FieldStatus, Value: dto.Status}},
		})
	}
	if dto.Price != nil {
		query = append(query, bson.E{Key: KeyPush, Value: bson.M{FieldPrices: dto.Price}})
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
