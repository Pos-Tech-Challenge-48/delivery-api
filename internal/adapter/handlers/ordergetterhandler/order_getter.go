package ordergetterhandler

import (
	"context"
	"net/http"

	_ "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/usecases/ordergetter"
	"github.com/gin-gonic/gin"
)

type OrderGetterHandler struct {
	OrderGetterUseCase ports.OrderGetter
}

func NewOrderGetterHandler(OrderGetterUseCase ports.OrderGetter) *OrderGetterHandler {
	return &OrderGetterHandler{
		OrderGetterUseCase: OrderGetterUseCase,
	}
}

// GetOrder godoc
// @Summary get all order
// @Description Get Order from DB
// @Tags order
// @Produce application/json
// @Success 200 {array} domain.Order "Order"
// @Failure 400 {object} string "invalid document"
// @Failure 404 {object} string "customer not find"
// @Failure 500 {object} string "general error"
// @Router /orders [get]
func (o *OrderGetterHandler) Handle(c *gin.Context) {
	ctx := context.Background()

	list, err := o.OrderGetterUseCase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"response": list})
}
