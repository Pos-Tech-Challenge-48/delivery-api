package producthandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/usecases/contract"
	"github.com/gin-gonic/gin"
)

type ProductAddHandler struct {
	svc contract.ProductContract
}

func NewProductAddHandler(productAddUseCase contract.ProductContract) *ProductAddHandler {
	return &ProductAddHandler{
		svc: productAddUseCase,
	}
}

func (p *ProductAddHandler) Handle(c *gin.Context) {
	ctx := context.Background()

	var product domain.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = product.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = p.svc.Add(ctx, &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Product created successfully"})
}
