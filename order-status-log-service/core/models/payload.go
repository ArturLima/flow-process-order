package models

import (
	"github.com/google/uuid"
)

type Payload struct {
	Order_id uuid.UUID `json:"order_id"`
	Status   int       `json:"status"`
}

/*
func (p *Payload) From() (order *Order) {
	return &Order{
		Id:          uuid.New(),
		Product_id:  p.Product_id,
		Customer_id: p.Customer_id,
		Amout:       p.Amout,
		CreatedAt:   time.Now().UTC().Local(),
	}
}
*/
