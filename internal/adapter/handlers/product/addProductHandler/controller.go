package producthandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/usecases/contract"
	"github.com/gin-gonic/gin"
)

var (
	addRouter    = "/v1/delivery/product/add"
	updateRouter = "/v1/delivery/product/update"
	deleteRouter = "/v1/delivery/product/delete"
	getAllRouter = "/v1/delivery/product/getAll"
)

type ProductHandler struct {
	svc contract.ProductContract
}

func NewProductHandler(productAddUseCase contract.ProductContract) *ProductHandler {
	return &ProductHandler{
		svc: productAddUseCase,
	}
}

func (p *ProductHandler) Handle(c *gin.Context) {
	ctx := context.Background()
	var response string
	var err error

	switch c.Request.URL.String() {
	case addRouter:
		response, err = p.AddHandle(ctx, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": response})
			return
		}
	case updateRouter:
		response, err = p.UpdateHandle(ctx, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": response})
			return
		}
	case deleteRouter:
		response, err = p.DeleteHandle(ctx, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": response})
			return
		}
	case getAllRouter:
		fmt.Println("GET ALL")
	}

	c.JSON(http.StatusOK, gin.H{"message": response})
}

func (p *ProductHandler) AddHandle(ctx context.Context, c *gin.Context) (retVal string, err error) {
	var product domain.Product

	err = c.BindJSON(&product)
	if err != nil {
		retVal = err.Error()
		return
	}

	err = product.Validate()
	if err != nil {
		retVal = err.Error()
		return
	}

	err = p.svc.Add(ctx, &product)
	if err != nil {
		retVal = err.Error()
		return
	}

	retVal = "Product created successfully"
	return
}

func (p *ProductHandler) UpdateHandle(ctx context.Context, c *gin.Context) (retVal string, err error) {
	var product domain.Product

	err = c.BindJSON(&product)
	if err != nil {
		retVal = err.Error()
		return
	}

	err = product.Validate()
	if err != nil {
		retVal = err.Error()
		return
	}

	err = p.svc.Update(ctx, &product)
	if err != nil {
		retVal = err.Error()
		return
	}

	retVal = "Product updated successfully"
	return
}

func (p *ProductHandler) DeleteHandle(ctx context.Context, c *gin.Context) (resp string, err error) {
	var product domain.Product

	err = c.BindJSON(&product)
	if err != nil {
		resp = err.Error()
		return
	}

	err = p.svc.Delete(ctx, product.ID)
	if err != nil {
		resp = err.Error()
		return
	}

	resp = "Product deleted successfully"
	return resp, nil
}
