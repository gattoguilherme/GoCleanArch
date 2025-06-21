package database

import (
	"GoCleanArch/internal/domain/entity"
	"errors"
	"sync"
)

// OrderRepositoryMock is a mock implementation of the OrderRepository interface.
type OrderRepositoryMock struct {
	mu     sync.Mutex
	orders map[int]*entity.Order
}

// NewOrderRepositoryMock creates a new OrderRepositoryMock.
func NewOrderRepositoryMock() *OrderRepositoryMock {
	return &OrderRepositoryMock{
		orders: make(map[int]*entity.Order),
	}
}

// Save saves an order to the mock database.
func (r *OrderRepositoryMock) Save(order *entity.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.OrderID] = order
	return nil
}

// GetByOrderID retrieves an order by its ID from the mock database.
func (r *OrderRepositoryMock) GetByOrderID(orderID int) (*entity.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	order, ok := r.orders[orderID]
	if !ok {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (r *OrderRepositoryMock) GetAll() ([]*entity.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orders := make([]*entity.Order, 0, len(r.orders))
	for _, order := range r.orders {
		orders = append(orders, order)
	}
	return orders, nil
}
