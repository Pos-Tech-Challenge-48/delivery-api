package productgetterhandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/productgetter"
	"github.com/gin-gonic/gin"
)

type ProductGetterHandler struct {
	svc productgetter.ProductGetter
}

func NewProductGetterHandler(productGetterUseCase *productgetter.ProductGetter) *ProductGetterHandler {
	return &ProductGetterHandler{
		svc: *productGetterUseCase,
	}
}

func (p *ProductGetterHandler) Handle(c *gin.Context) {
	ctx := context.Background()
	category := c.Request.URL.Query().Get("category")

	list, err := p.svc.GetAll(ctx, category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"response": list})
}
