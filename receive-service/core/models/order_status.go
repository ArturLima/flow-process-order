package models

import "github.com/google/uuid"

type OrderStatusEnum int

const (
	receive    OrderStatusEnum = 1
	validate   OrderStatusEnum = 2
	process    OrderStatusEnum = 3
	expedition OrderStatusEnum = 4
)

type OrderStatus struct {
	Order_id uuid.UUID       `json:"order_id"`
	Status   OrderStatusEnum `json:"status"`
}

func (os *OrderStatus) From() (order *OrderStatus) {
	return &OrderStatus{
		Order_id: os.Order_id,
		Status:   receive,
	}
}
