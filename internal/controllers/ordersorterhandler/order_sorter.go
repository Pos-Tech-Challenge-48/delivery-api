package ordersorterhandler

import (
	"net/http"

	_ "github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
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

// OrderSorter godoc
// @Summary get all order
// @Description Get Order from DB
// @Tags order
// @Produce application/json
// @Success 200 {array} entities.Order "Order"
// @Failure 400 {object} string "invalid document"
// @Failure 404 {object} string "customer not find"
// @Failure 500 {object} string "general error"
// @Router /orders/sorted [get]
func (o *OrderSorterHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()

	list, err := o.OrderSorterUseCase.GetAllSortedByStatus(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": list})
}
