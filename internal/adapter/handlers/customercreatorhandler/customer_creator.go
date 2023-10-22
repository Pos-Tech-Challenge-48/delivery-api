package customercreatorhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/usecases/customercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customercreator"
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
		return
	}

	ctx := context.Background()

	if err := h.customerCreatorUseCase.Create(ctx, &customer); err != nil {
		if errors.Is(err, domain.ErrCustomerEmptyEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrCustomerEmptyName) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, customercreator.ErrAlreadyExistsCustomer) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrCustomerInvalidDocument) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, domain.ErrCustomerInvalidEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
