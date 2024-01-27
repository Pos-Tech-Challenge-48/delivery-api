package ordersorter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

//go:generate mockgen -destination=./../../../mocks/ordersortermock/order_sorter_mock.go -source=./order_sorter.go -package=ordersortermock
type OrderSorter interface {
	GetAllSortedByStatus(ctx context.Context) ([]entities.Order, error)
}
