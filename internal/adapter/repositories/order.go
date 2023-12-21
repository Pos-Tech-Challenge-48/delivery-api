package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

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

func (r *OrderRepository) GetAll(ctx context.Context) ([]domain.Order, error) {
	queryParams := []interface{}{}

	query := `
	SELECT
		restaurant_order_id,
		restaurant_order_customer_id,
		status.status_name,
		restaurant_order_amount,
		restaurant_order.created_date_db,
		restaurant_order.last_modified_date_db,
		json_agg(json_build_object(
				'id',     order_item.order_item_id,
				'name', p.product_name
			)) AS order_items
	FROM restaurant_order
			join order_item on restaurant_order.restaurant_order_id = order_item.order_item_order_id
			join product p on order_item.order_item_product_id = p.product_id
			JOIN status ON	restaurant_order_status_id = status_id
	GROUP BY
		restaurant_order_id,
		restaurant_order_customer_id,
		status.status_name,
		restaurant_order_amount,
		restaurant_order.created_date_db,
		restaurant_order.last_modified_date_db;
	`

	rows, err := r.db.Query(query, queryParams...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]domain.Order, 0)
	var order domain.Order

	var orderProducts []byte

	for rows.Next() {
		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.Status,
			&order.Amount,
			&order.CreatedDate,
			&order.LastModifiedDate,
			&orderProducts,
		)

		if err != nil {
			return result, err
		}

		err = json.Unmarshal(orderProducts, &order.OrderProduct)
		if err != nil {
			return result, err
		}

		result = append(result, order)

	}

	return result, nil

}
