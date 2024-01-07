package productupdate

import (
	"context"
	"fmt"
	"strings"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type ProductUpdate struct {
	productRepository ports.ProductRepository
}

func NewProductUpdate(productRepository ports.ProductRepository) *ProductUpdate {
	return &ProductUpdate{
		productRepository: productRepository,
	}
}

func (p *ProductUpdate) Update(ctx context.Context, data *entities.Product) error {

	categoryName := strings.Title(strings.ToLower(data.CategoryID))
	categoryId, err := p.productRepository.GetCategoryID(ctx, categoryName)
	if err != nil {
		return fmt.Errorf("error getting categoryID: %w", err)
	}

	data.CategoryID = categoryId
	err = p.productRepository.Update(ctx, data)
	if err != nil {
		return fmt.Errorf("error to updating product: %w", err)
	}

	err = p.productRepository.UpdateImage(ctx, data.ID, data.Image)
	if err != nil {
		return fmt.Errorf("error to updating image of product: %w", err)
	}

	return nil
}
