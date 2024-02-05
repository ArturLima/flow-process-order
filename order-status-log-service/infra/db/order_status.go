package db

import (
	"context"
	"order-status-log-service/core/models"
	"os"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IOrderStatusRepository interface {
	Insert(os *models.OrderStatus) (err error)
	Get(orderId uuid.UUID) (orderStatus models.OrderStatus, err error)
}

type OrderStatusRepository struct {
	context *bun.DB
}

func NewOrderRepository() IOrderStatusRepository {
	return &OrderStatusRepository{
		context: NewDB(os.Getenv("DB_URL")),
	}
}

func (o *OrderStatusRepository) Insert(os *models.OrderStatus) (err error) {
	_, err = o.context.
		NewInsert().
		Model(os).
		Exec(context.Background())

	return
}

func (o *OrderStatusRepository) Get(order uuid.UUID) (orderStatus models.OrderStatus, err error) {
	err = o.context.
		NewSelect().
		Model(&orderStatus).
		Where("id = ? ", order).
		Scan(context.Background())
	return
}
