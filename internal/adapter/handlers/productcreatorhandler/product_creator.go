package productcreatorhandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/productcreator"
	"github.com/gin-gonic/gin"
)

type ProductCreatorHandler struct {
	svc productcreator.ProductCreator
}

func NewProductCreatorHandler(productAddUseCase *productcreator.ProductCreator) *ProductCreatorHandler {
	return &ProductCreatorHandler{
		svc: *productAddUseCase,
	}
}

func (p *ProductCreatorHandler) Handle(c *gin.Context) {
	ctx := context.Background()
	var product domain.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	err = product.Validate()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	err = p.svc.Add(ctx, &product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "Product created successfully"})
}
