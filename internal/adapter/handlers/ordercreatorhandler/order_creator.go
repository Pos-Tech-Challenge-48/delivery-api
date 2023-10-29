package ordercreatorhandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/usecases/ordercreator"
	"github.com/gin-gonic/gin"
)

type OrderCreatorHandler struct {
	OrderCreatorUseCase ports.OrderCreator
}

func NewOrderCreatorHandler(OrderCreatorUseCase ports.OrderCreator) *OrderCreatorHandler {
	return &OrderCreatorHandler{
		OrderCreatorUseCase: OrderCreatorUseCase,
	}
}

// CreateOrder godoc
// @Summary create Order
// @Description save Order in DB
// @Param Order body domain.Order true "Order"
// @Tags Order
// @Produce application/json
// @Success 201
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /Order [post]
func (h *OrderCreatorHandler) Handle(c *gin.Context) {
	var order domain.Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	if err := h.OrderCreatorUseCase.Create(ctx, &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
