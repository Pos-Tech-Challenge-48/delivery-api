package customergetterhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/usecases/customergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customergetter"
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
		if errors.Is(err, customergetter.ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, domain.ErrCustomerInvalidDocument) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)

}
