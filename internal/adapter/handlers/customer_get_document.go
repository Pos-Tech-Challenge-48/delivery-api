package handlers

import (
	"context"
	"net/http"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type CustomerGetterHandler struct {
	customerGetterUseCase ports.CustomerGetter
}

func NewCustomerGetterHandler(customerGetterUseCase ports.CustomerGetter) *CustomerGetterHandler {
	return &CustomerGetterHandler{
		customerGetterUseCase: customerGetterUseCase,
	}
}

func (h *CustomerGetterHandler) Handle(c *gin.Context) {

	ctx := context.Background()

	document := c.Query("document")
	customerInput := domain.Customer{
		Document: document,
	}

	customer, err := h.customerGetterUseCase.Get(ctx, &customerInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, customer)
}
