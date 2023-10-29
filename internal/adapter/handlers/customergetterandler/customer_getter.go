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

// CreateCustomer godoc
// @Summary get customer by document
// @Param   document     query    string     true        "Document"
// @Tags customer
// @Produce application/json
// @Success 200 {object} domain.Customer "Customer"
// @Failure 400 {object} string "invalid document"
// @Failure 404 {object} string "customer not find"
// @Failure 500 {object} string "general error"
// @Router /customer [get]
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

		if errors.Is(err, domain.ErrCustomerEmptyDocument) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)

}
