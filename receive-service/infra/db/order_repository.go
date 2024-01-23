package db

import (
	"os"

	"github.com/uptrace/bun"
)

type IOrderRepository interface {
	Insert(test string) (done bool)
}

type OrderRepository struct {
	context *bun.DB
}

func NewOrderRepository() IOrderRepository {
	return &OrderRepository{
		context: NewDB(os.Getenv("DB_URL")),
	}
}

func (o *OrderRepository) Insert(teste string) (done bool) {
	return true
}
