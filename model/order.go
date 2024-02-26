package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64     `json:"order_id"`
	CustomerID  uuid.UUID  `json:"customer_id"` // comingfrom another service
	LineItems   []LineIten `json:"line_items"`
	CreatedAt   *time.Time `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type LineIten struct {
	ItemID   uuid.UUID `json:"item_id"` // comingfrom another service
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
