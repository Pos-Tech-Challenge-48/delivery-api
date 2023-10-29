package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/repositories"
)

type ProductService struct {
	productRepository ports.ProductRepository
}

func NewProductService(productRepository ports.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) Add(ctx context.Context, data *domain.Product) error {

	categoryName := strings.Title(strings.ToLower(data.CategoryID))
	categoryId, err := p.productRepository.GetCategoryID(ctx, categoryName)
	if err != nil {
		return fmt.Errorf("error getting categoryID: %w", err)
	}

	data.CategoryID = categoryId
	err = p.productRepository.Add(ctx, data)
	if err != nil {
		return fmt.Errorf("error to submit create product: %w", err)
	}

	return nil
}
