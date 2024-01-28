package services

import (
	"receive-service/core/models"
	"receive-service/infra/db"
)

type IOrderService interface {
	InsertOrder(payload *models.Order) (err error)
}

type OrderService struct {
	repository db.IOrderRepository
}

func NewOrderService() IOrderService {
	return &OrderService{
		repository: db.NewOrderRepository(),
	}
}

func (o *OrderService) InsertOrder(order *models.Order) (err error) {
	return o.repository.Insert(order)
}
