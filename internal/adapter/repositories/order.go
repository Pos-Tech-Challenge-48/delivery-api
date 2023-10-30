package repositories

import (
	"context"
	"database/sql"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	_ "github.com/lib/pq"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Save(ctx context.Context, order *domain.Order) error {
	query := `
	INSERT INTO restaurant_order (restaurant_order_id, restaurant_order_customer_id, restaurant_order_status_id, restaurant_order_amount)
	VALUES ($1, $2, (SELECT status_id from Status where status_name = $3), $4)
`

	_, err := r.db.Exec(
		query,
		order.ID,
		order.CustomerID,
		order.Status,
		order.Amount)

	if err != nil {
		return err
	}

	for _, p := range order.OrderProduct {

		query := `
        INSERT INTO order_item (order_item_product_id, order_item_order_id)
        VALUES ($1, $2)
    `
		_, err := r.db.Exec(query, p.ID, order.ID)
		if err != nil {
			return err
		}

	}

	return err
}
