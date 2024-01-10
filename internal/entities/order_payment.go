package entities

import (
	"errors"
	"time"
)

var (
	ErrOrderNotExist = errors.New("customer empty email")
	ErrOrderNotReady = errors.New("customer invalid email")
)

type OrderPayment struct {
	OrderID         string     `json:"order_id"`
	ExpirationDate  *time.Time `json:"expiration_date"`
	Description     string     `json:"description"`
	NotificationURL string     `json:"notification_url"`
}

type OrderPaymentResponse struct {
	PaymentID string `json:"payment_id"`
	QRCode    string `json:"qr_code"`
}
