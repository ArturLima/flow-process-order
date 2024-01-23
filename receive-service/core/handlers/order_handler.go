package handlers

import "receive-service/core/services"

type IOrderHandler interface {
	Execute(b []byte) (done bool)
}

type OrderHandler struct {
	service services.IOrderService
}

func NewOrderHandler() IOrderHandler {
	return &OrderHandler{
		service: services.NewOrderService(),
	}
}

func (o *OrderHandler) Execute(b []byte) (done bool) {

	return o.service.InsertOrder("teste")
}
