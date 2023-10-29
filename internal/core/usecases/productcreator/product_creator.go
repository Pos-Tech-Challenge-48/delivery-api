package productcreator

import (
	"context"
	"strings"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/repositories"
	"github.com/pkg/errors"
)

type ProductCreator struct {
	productRepository ports.ProductRepository
}

func NewProductCreator(productRepository ports.ProductRepository) *ProductCreator {
	return &ProductCreator{
		productRepository: productRepository,
	}
}

func (p *ProductCreator) Add(ctx context.Context, data *domain.Product) error {

	categoryName := strings.Title(strings.ToLower(data.CategoryID))
	categoryId, err := p.productRepository.GetCategoryID(ctx, categoryName)
	if err != nil {
		return errors.Wrap(err, "error getting categoryID")
	}

	data.CategoryID = categoryId
	err = p.productRepository.Add(ctx, data)
	if err != nil {
		return errors.Wrap(err, "error to submit create product")
	}

	return nil
}
