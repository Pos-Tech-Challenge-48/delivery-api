package api

import (
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	"github.com/google/uuid"
)

type PaymentGatewayAPI struct {
	HostAPI string
	APIKey  string
	Path    string
}

func NewPaymentGateway(hostAPI, APIKey, path string) *PaymentGatewayAPI {
	return &PaymentGatewayAPI{
		HostAPI: hostAPI,
		APIKey:  hostAPI,
		Path:    path,
	}
}

func (p *PaymentGatewayAPI) CreatePayment(payment *entities.Payment) (*entities.PaymentResponse, error) {
	return &entities.PaymentResponse{
		ID:     uuid.NewString(),
		QRData: "00020101021243650016COM.MERCADOLIBRE02013063638f1192a-5fd1-4180-a180-8bcae3556bc35204000053039865802BR5925IZABEL AAAA DE MELO6007BARUERI62070503***63040B6D",
	}, nil
}
