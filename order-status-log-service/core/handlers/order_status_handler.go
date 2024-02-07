package handlers

import (
	"encoding/json"
	"order-status-log-service/core/models"
	"order-status-log-service/core/services"
	"order-status-log-service/utils"
)

type IOrderStatusHandler interface {
	Execute(b []byte) (done bool)
}

type OrderStatusHandler struct {
	service services.IOrderStatusService
}

func NewOrderStatusHandler() IOrderStatusHandler {
	return &OrderStatusHandler{
		service: services.NewOrderStatusService(),
	}
}

func (os *OrderStatusHandler) Execute(b []byte) (done bool) {
	var p models.Payload

	if err := json.Unmarshal(b, &p); err != nil {
		utils.FailWithError("failure to unmarshal object", err)
	}

	orderStatus := p.From()

	order, err := os.service.Get(orderStatus.OrderId)
	if err != nil {
		utils.FailWithError("get order", err)
	}

	if order.Status == 0 {
		os.service.update(orderStatus)
	}

	err = os.service.Insert(orderStatus)

	if err != nil {
		utils.FailWithError("error to create status", err)
	}

	return true
}
