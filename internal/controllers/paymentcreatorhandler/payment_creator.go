package paymentcreatorhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/usecases/paymentcreator"
	"github.com/gin-gonic/gin"
)

type PaymentCreatorHandler struct {
	PaymentCreatorUseCase interfaces.PaymentCreator
}

func NewPaymentCreatorHandler(PaymentCreatorUseCase interfaces.PaymentCreator) *PaymentCreatorHandler {
	return &PaymentCreatorHandler{
		PaymentCreatorUseCase: PaymentCreatorUseCase,
	}
}


// CreatePayment godoc
// @Summary create payment
// @Description save Payment in DB
// @Param Payment body entities.Payment true "Payment"
// @Tags Payment
// @Produce application/json
// @Success 201
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /payment [post]
func (h *PaymentCreatorHandler) Handle(c *gin.Context) {
	var orderPayment entities.OrderPayment

	orderID := c.Param("order_id")

	if err := c.BindJSON(&orderPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	orderPayment.OrderID = orderID

	orderPaymentResponse, err := h.PaymentCreatorUseCase.Create(ctx, &orderPayment)
	if err != nil {
		if errors.Is(err, entities.ErrOrderNotExist) || errors.Is(err, entities.ErrOrderNotReady) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, orderPaymentResponse)
}
