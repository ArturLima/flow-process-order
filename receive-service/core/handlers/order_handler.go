package handlers

import (
	"encoding/json"
	"receive-service/core/models"
	"receive-service/core/services"
	"receive-service/infra/rabbitmq"
	"receive-service/utils"
)

type IOrderHandler interface {
	Execute(b []byte) (done bool)
}

type OrderHandler struct {
	service   services.IOrderService
	publisher rabbitmq.IPublisher
}

func NewOrderHandler() IOrderHandler {
	return &OrderHandler{
		service:   services.NewOrderService(),
		publisher: rabbitmq.NewPublisher(),
	}
}

func (o *OrderHandler) Execute(b []byte) (done bool) {

	var p models.Payload
	var os models.OrderStatus

	if err := json.Unmarshal(b, &p); err != nil {
		utils.FailWithError("failure to unmarshal object", err)
	}

	order := p.From()

	os.Order_id = order.Id

	orderStatus := os.From()

	err := o.publisher.SendToOrderStatus(orderStatus)
	if err != nil {
		utils.FailWithError("failure to send to order status job", err)
	}

	err = o.service.InsertOrder(order)
	if err != nil {
		utils.FailWithError("failure to unmarshal object", err)
	}

	//send to queue validate

	return true
}
