package handlers

import (
	"context"
	"net/http"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type CustomerCreatorHandler struct {
	customerCreatorUseCase ports.CustomerCreator
}

func NewCustomerCreatorHandler(customerCreatorUseCase ports.CustomerCreator) *CustomerCreatorHandler {
	return &CustomerCreatorHandler{
		customerCreatorUseCase: customerCreatorUseCase,
	}
}

func (h *CustomerCreatorHandler) Handle(c *gin.Context) {
	var customer domain.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	ctx := context.Background()

	if err := h.customerCreatorUseCase.Create(ctx, &customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, customer)
}
