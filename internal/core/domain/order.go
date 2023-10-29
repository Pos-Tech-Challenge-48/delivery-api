package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID               string         `json:"id"`
	CustomerID       string         `json:"customer_id"`
	Status           string         `json:"status"`
	Amount           float64        `json:"amount"`
	OrderProduct     []OrderProduct `json:"products"`
	CreatedDate      time.Time      `json:"created_date_db"`
	LastModifiedDate time.Time      `json:"last_modified_date_db"`
}

type OrderProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewOrder(customerID string, products []OrderProduct) *Order {
	return &Order{
		ID:               uuid.NewString(),
		CustomerID:       customerID,
		Status:           "Recebido",
		OrderProduct:     products,
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
	}
}

func (*Order) Validate() error {
	return nil
}
