package entity

import "time"

// Order represents a customer order.
type Order struct {
	ID        string    `json:"id"`
	Data      string    `json:"Data"`
	OrderID   int       `json:"OrderId"`
	Status    string    `json:"Status"`
	Paid      bool      `json:"Paid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
