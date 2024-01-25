package ordersorterhandler

import (
	"net/http"

	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/usecases/ordersorter"
	"github.com/gin-gonic/gin"
)

type OrderSorterHandler struct {
	OrderSorterUseCase interfaces.OrderSorter
}

func NewOrderSorterHandler(OrderSorterUseCase interfaces.OrderSorter) *OrderSorterHandler {
	return &OrderSorterHandler{
		OrderSorterUseCase: OrderSorterUseCase,
	}
}

func (o *OrderSorterHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()

	list, err := o.OrderSorterUseCase.GetAllSortedByStatus(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": list})
}
