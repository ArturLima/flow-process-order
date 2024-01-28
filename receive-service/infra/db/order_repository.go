package db

import (
	"context"
	"os"
	"receive-service/core/models"

	"github.com/uptrace/bun"
)

type IOrderRepository interface {
	Insert(payload *models.Order) (err error)
}

type OrderRepository struct {
	context *bun.DB
}

func NewOrderRepository() IOrderRepository {
	return &OrderRepository{
		context: NewDB(os.Getenv("DB_URL")),
	}
}

func (o *OrderRepository) Insert(payload *models.Order) (err error) {
	_, err = o.context.
		NewInsert().
		Model(payload).
		Exec(context.Background())

	return
}
