package entities

func NewPayment(orderPayment *OrderPayment, order *Order) *Payment {
	return &Payment{
		Description:     orderPayment.Description,
		TotalAmount:     order.Amount,
		NotificationURL: orderPayment.NotificationURL,
		UserID:          "12345", //Default user to Mercado Pago for example
	}
}

type Payment struct {
	Description     string  `json:"description"`
	NotificationURL string  `json:"notification_url"`
	TotalAmount     float64 `json:"total_amount"`
	UserID          string  `json:"user_id"`
}

type PaymentResponse struct {
	ID     string `json:"id"`
	QRData string `json:"qr_data"`
}
