package productcreator

import (
	"context"
	"strings"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type ProductCreator struct {
	productRepository interfaces.ProductRepository
}

func NewProductCreator(productRepository interfaces.ProductRepository) *ProductCreator {
	return &ProductCreator{
		productRepository: productRepository,
	}
}

func (p *ProductCreator) Add(ctx context.Context, data *entities.Product) error {

	categoryName := strings.Title(strings.ToLower(data.CategoryID))
	categoryId, err := p.productRepository.GetCategoryID(ctx, categoryName)
	if err != nil {
		return err
	}

	data.CategoryID = categoryId
	err = p.productRepository.Add(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
