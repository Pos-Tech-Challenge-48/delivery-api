package paymentwebhookhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/usecases/paymentwebhook"
	"github.com/gin-gonic/gin"
)

type PaymentWebhookHandler struct {
	PaymentWebhookUseCase interfaces.PaymentWebhook
}

func NewPaymentWebhookHandler(PaymentWebhookUseCase interfaces.PaymentWebhook) *PaymentWebhookHandler {
	return &PaymentWebhookHandler{
		PaymentWebhookUseCase: PaymentWebhookUseCase,
	}
}

// Payment Webhook godoc
// @Summary create Payment
// @Description update payment and set paid in DB
// @Param order_id path string true "Order ID"
// @Param OrderPayment body entities.OrderPaymentRequest true "OrderPaymentRequest"
// @Tags payment
// @Produce application/json
// @Success 201
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /orders/:order_id/payment/webhook [post]
func (h *PaymentWebhookHandler) Handle(c *gin.Context) {

	orderID := c.Param("order_id")

	ctx := context.Background()

	var paymentWebhookBody entities.OrderPaymentRequest

	if err := c.BindJSON(&paymentWebhookBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.PaymentWebhookUseCase.Update(ctx, orderID, paymentWebhookBody)
	if err != nil {
		if errors.Is(err, entities.ErrOrderNotExist) || errors.Is(err, entities.ErrOrderNotReadyToConfirmPayment) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
