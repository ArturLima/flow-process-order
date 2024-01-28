package models

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	Id          uuid.UUID `json:"id"`
	Product_id  uuid.UUID `json:"product_id"`
	Customer_id uuid.UUID `json:"customer_id"`
	Amout       float64   `json:"amout"`
}

func (p *Payload) From() (order *Order) {
	return &Order{
		Id:          uuid.New(),
		Product_id:  p.Product_id,
		Customer_id: p.Customer_id,
		Amout:       p.Amout,
		CreatedAt:   time.Now().UTC().Local(),
	}
}
