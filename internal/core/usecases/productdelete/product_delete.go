package productdelete

import (
	"context"
	"fmt"

	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/repositories"
)

type ProductDelete struct {
	productRepository ports.ProductRepository
}

func NewProductDelete(productRepository ports.ProductRepository) *ProductDelete {
	return &ProductDelete{
		productRepository: productRepository,
	}
}

func (p *ProductDelete) Delete(ctx context.Context, producID string) error {

	err := p.productRepository.Delete(ctx, producID)
	if err != nil {
		return fmt.Errorf("error to submit delete product: %w", err)
	}

	return nil
}
