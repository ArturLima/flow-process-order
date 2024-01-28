package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Order struct {
	bun.BaseModel `bun:"table:order,alias:OR"`

	Id          uuid.UUID `bun:"id,pk"`
	Product_id  uuid.UUID `bun:"product_id"`
	Customer_id uuid.UUID `bun:"customer_id"`
	Amout       float64   `bun:"amount"`
	CreatedAt   time.Time `bun:"created_at"`
}
