package producthandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/usecases/contract"
	"github.com/gin-gonic/gin"
)

var (
	addRouter    = "POST"
	updateRouter = "PUT"
	deleteRouter = "DELETE"
	getAllRouter = "GET"
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
	var response interface{}
	var err error

	method := c.Request.Method

	switch method {
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
		response, err = p.GetAll(ctx, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": response})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
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

func (p *ProductHandler) GetAll(ctx context.Context, c *gin.Context) (list []domain.Product, err error) {
	category := c.Request.URL.Query().Get("category")

	list, err = p.svc.GetAll(ctx, category)
	if err != nil {
		return list, err
	}

	return list, nil
}
