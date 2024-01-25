package entities

import (
	"errors"
	"time"
)

var (
	ErrOrderNotExist                 = errors.New("order not exists")
	ErrOrderNotReady                 = errors.New("order not ready")
	ErrOrderNotReadyToConfirmPayment = errors.New("order not ready to confirm the payment")
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
