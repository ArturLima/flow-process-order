package services

import "receive-service/infra/db"

type IOrderService interface {
	InsertOrder(test string) (done bool)
}

type OrderService struct {
	repository db.IOrderRepository
}

func NewOrderService() IOrderService {
	return &OrderService{
		repository: db.NewOrderRepository(),
	}
}

func (o *OrderService) InsertOrder(test string) (done bool) {
	return o.repository.Insert(test)
}
