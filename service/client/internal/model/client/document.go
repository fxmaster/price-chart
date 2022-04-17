package client

import (
	"context"
	"time"
)

type Repository interface {
	CheckByID(ctx context.Context, id string) (bool, error)
	CheckByChatID(ctx context.Context, chatID uint64) (bool, error)
	CheckByUsername(ctx context.Context, username string) (bool, error)
	LoadByID(ctx context.Context, id string) (*Client, error)
	LoadByChatID(ctx context.Context, chatID uint64) (*Client, error)
	Store(ctx context.Context, client Client) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, dto *UpdateDTO) error
	AttachUrl(ctx context.Context, id, url string) error
	DetachUrl(ctx context.Context, id, url string) error
}

type Client struct {
	ID        string    `bson:"_id" json:"_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	ChatID    uint64    `bson:"chat_id" json:"chat_id"`
	FirstName string    `bson:"first_name" json:"first_name"`
	LastName  string    `bson:"last_name" json:"last_name"`
	Username  string    `bson:"username" json:"username"`
	ChatType  string    `bson:"chat_type" json:"chat_type"`
	Urls      []string  `bson:"urls" json:"urls"`
}

type UpdateDTO struct {
	ChatID    uint64
	FirstName string
	LastName  string
	Username  string
	ChatType  string
}

func (d Client) areRequiredFieldsEmpty() bool {
	return d.ID == "" || d.CreatedAt.IsZero() || d.UpdatedAt.IsZero() ||
		d.ChatID == 0 || d.FirstName == "" || d.LastName == "" || d.Username == "" ||
		d.ChatType == "" || len(d.Urls) == 0
}
