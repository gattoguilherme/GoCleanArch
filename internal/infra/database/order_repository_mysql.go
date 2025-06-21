package database

import (
	"GoCleanArch/internal/domain/entity"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// OrderRepositoryMySQL implements the OrderRepository interface for MySQL.
type OrderRepositoryMySQL struct {
	DB *sql.DB
}

// NewOrderRepository creates a new MySQL order repository.
func NewOrderRepository(db *sql.DB) *OrderRepositoryMySQL {
	return &OrderRepositoryMySQL{DB: db}
}

// Save saves an order to the database.
func (r *OrderRepositoryMySQL) Save(order *entity.Order) error {
	stmt, err := r.DB.Prepare("INSERT INTO orders (id, data, order_id, status, paid, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID, order.Data, order.OrderID, order.Status, order.Paid, order.CreatedAt, order.UpdatedAt)
	return err
}

// GetByOrderID retrieves an order from the database by its ID.
func (r *OrderRepositoryMySQL) GetByOrderID(orderID int) (*entity.Order, error) {
	row := r.DB.QueryRow("SELECT id, data, order_id, status, paid, created_at, updated_at FROM orders WHERE id = ?", orderID)

	var order entity.Order
	err := row.Scan(&order.ID, &order.Data, &order.OrderID, &order.Status, &order.Paid, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No order found
		}
		return nil, err
	}

	return &order, nil
}

// GetAll retrieves all orders from the database.
func (r *OrderRepositoryMySQL) GetAll() ([]*entity.Order, error) {
	rows, err := r.DB.Query("SELECT id, data, order_id, status, paid, created_at, updated_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.ID, &order.Data, &order.OrderID, &order.Status, &order.Paid, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}
