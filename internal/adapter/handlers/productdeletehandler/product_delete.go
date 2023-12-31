package productdeletehandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/productdelete"
	"github.com/gin-gonic/gin"
)

type ProductDeleteHandler struct {
	svc productdelete.ProductDelete
}

func NewProductDeleteHandler(productDeleteUseCase *productdelete.ProductDelete) *ProductDeleteHandler {
	return &ProductDeleteHandler{
		svc: *productDeleteUseCase,
	}
}

// Products godoc
// @Summary delete product
// @Description delete product in DB
// @Param product body domain.Product true "Product"
// @Tags product
// @Produce application/json
// @Success 200
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /products [delete]
func (p *ProductDeleteHandler) Handle(c *gin.Context) {
	ctx := context.Background()
	var product domain.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	err = p.svc.Delete(ctx, product.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "Product deleted successfully"})
}
