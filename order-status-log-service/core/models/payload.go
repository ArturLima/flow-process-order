package models

import (
	"github.com/google/uuid"
)

type Payload struct {
	Order_id uuid.UUID `json:"order_id"`
	Status   int       `json:"status"`
}

func (p *Payload) From() (order *OrderStatus) {
	return &OrderStatus{
		OrderId: p.Order_id,
		Status:  int64(p.Status),
	}
}
