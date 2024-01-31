package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type OrderStatus struct {
	bun.BaseModel `bun:"table:order_status,alias:OS"`

	OrderId uuid.UUID `bun:"id,pk"`
	Status  int64     `bun:"status"`
}
