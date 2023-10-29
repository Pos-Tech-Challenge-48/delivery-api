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

func (p *ProductService) Update(ctx context.Context, data *domain.Product) error {

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

	return nil
}

func (p *ProductService) Delete(ctx context.Context, producID string) error {

	err := p.productRepository.Delete(ctx, producID)
	if err != nil {
		return fmt.Errorf("error to submit delete product: %w", err)
	}

	return nil
}
