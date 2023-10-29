package productupdatehandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/productupdate"
	"github.com/gin-gonic/gin"
)

type ProductUpdateHandler struct {
	svc productupdate.ProductUpdate
}

func NewProductUpdateHandler(productUpdateUseCase *productupdate.ProductUpdate) *ProductUpdateHandler {
	return &ProductUpdateHandler{
		svc: *productUpdateUseCase,
	}
}

func (p *ProductUpdateHandler) Handle(c *gin.Context) {
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

	err = p.svc.Update(ctx, &product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "Product updated successfully"})
}
