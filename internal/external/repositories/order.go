package repositories

import (
	"context"
	"database/sql"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
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

func (r *OrderRepository) Save(ctx context.Context, order *entities.Order) error {
	query := `
	INSERT INTO restaurant_order (restaurant_order_id, restaurant_order_customer_id, restaurant_order_status_id, restaurant_order_amount)
	VALUES ($1, $2, (SELECT status_id from status where status_name = $3), $4)
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

func (r *OrderRepository) GetAll(ctx context.Context) ([]entities.Order, error) {
	queryParams := []interface{}{}

	query := `
		SELECT
			restaurant_order_id,
			restaurant_order_customer_id,
			status.status_name,
			restaurant_order_amount,
			restaurant_order.created_date_db,
			restaurant_order.last_modified_date_db
		FROM restaurant_order
		JOIN status ON
				restaurant_order_status_id = status_id
		ORDER BY restaurant_order.created_date_db ASC;
	`

	rows, err := r.db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]entities.Order, 0)
	var order entities.Order

	for rows.Next() {

		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.Status,
			&order.Amount,
			&order.CreatedDate,
			&order.LastModifiedDate,
		)
		if err != nil {
			return result, err
		}

		productQuery := `
			SELECT
			order_item_product_id,
			product.product_name
			FROM order_item
			JOIN product ON order_item.order_item_product_id = product.product_id
			WHERE order_item_order_id = $1;
		`

		productRows, err := r.db.Query(productQuery, &order.ID)

		if err != nil {
			return nil, err
		}

		defer productRows.Close()

		products := make([]entities.OrderProduct, 0)
		var product entities.OrderProduct

		for productRows.Next() {
			err := productRows.Scan(
				&product.ID,
				&product.Name,
			)
			if err != nil {
				return result, err
			}

			products = append(products, product)

		}

		order.OrderProduct = products

		result = append(result, order)
	}

	return result, nil

}

// TODO
func (r *OrderRepository) GetByID(ctx context.Context, orderID string) (*entities.Order, error) {
	return &entities.Order{
		ID:     "12345",
		Status: "Recebido",
	}, nil
}

// TODO
func (r *OrderRepository) Update(ctx context.Context, order *entities.Order) error {
	return nil
}
