package repository

import "GoCleanArch/internal/domain/entity"

// OrderRepository is an interface for interacting with order data.
type OrderRepository interface {
	Save(order *entity.Order) error
	GetByOrderID(orderID int) (*entity.Order, error)
	GetAll() ([]*entity.Order, error)
}

// OrderMessageQueue is an interface for sending order messages.
type OrderMessageQueue interface {
	Send(order *entity.Order) error
}
