package goods

import (
	"context"
	"time"
)

type Repository interface {
	CheckByID(ctx context.Context, id string) (bool, error)
	CheckByURL(ctx context.Context, url string) (bool, error)
	CheckAnotherByURL(ctx context.Context, id, url string) (bool, error)
	LoadByID(ctx context.Context, id string) (*Goods, error)
	LoadByURL(ctx context.Context, url string) (*Goods, error)
	Store(ctx context.Context, goods Goods) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, dto *UpdateDTO) error
}

type Goods struct {
	ID        string    `bson:"_id" json:"_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	URL       string    `bson:"url" json:"url"`
	Status    string    `bson:"status" json:"status"`
	Prices    []*Price  `bson:"prices" json:"prices"`
}

type Price struct {
	Value     float64   `bson:"value" json:"value"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type UpdateDTO struct {
	URL    string
	Price  *float64
	Status string
}

func (d Goods) areRequiredFieldsEmpty() bool {
	return d.ID == "" || d.CreatedAt.IsZero() || d.UpdatedAt.IsZero() ||
		d.URL == "" || d.Status == "" || len(d.Prices) == 0
}
