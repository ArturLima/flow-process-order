package handlers

import (
	"encoding/json"
	"order-status-log-service/core/models"
	"order-status-log-service/utils"
)

type IOrderStatusHandler interface {
	Execute(b []byte) (done bool)
}

type OrderStatusHandler struct {
}

func NewOrderStatusHandler() IOrderStatusHandler {
	return &OrderStatusHandler{}
}

func (os *OrderStatusHandler) Execute(b []byte) (done bool) {
	var p models.Payload

	if err := json.Unmarshal(b, &p); err != nil {
		utils.FailWithError("failure to unmarshal object", err)
	}

	return true
}
