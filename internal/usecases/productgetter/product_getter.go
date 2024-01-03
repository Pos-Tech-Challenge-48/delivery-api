package productgetter

import (
	"context"
	"strings"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type ProductGetter struct {
	productRepository interfaces.ProductRepository
}

func NewProductGetter(productRepository interfaces.ProductRepository) *ProductGetter {
	return &ProductGetter{
		productRepository: productRepository,
	}
}

func (p *ProductGetter) GetAll(ctx context.Context, category string) ([]entities.Product, error) {

	categoryName := strings.Title(strings.ToLower(category))

	list, err := p.productRepository.GetAll(ctx, categoryName)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}
