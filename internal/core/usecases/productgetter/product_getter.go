package productgetter

import (
	"context"
	"strings"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/repositories"
	"github.com/pkg/errors"
)

type ProductGetter struct {
	productRepository ports.ProductRepository
}

func NewProductGetter(productRepository ports.ProductRepository) *ProductGetter {
	return &ProductGetter{
		productRepository: productRepository,
	}
}

func (p *ProductGetter) GetAll(ctx context.Context, category string) ([]domain.Product, error) {

	categoryName := strings.Title(strings.ToLower(category))

	list, err := p.productRepository.GetAll(ctx, categoryName)
	if err != nil {
		return nil, errors.Wrap(err, "error getting list of product")
	}

	return list, nil
}
