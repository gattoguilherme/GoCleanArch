package usecase

import (
	"GoCleanArch/internal/domain/entity"
	"GoCleanArch/internal/domain/repository"
	"time"
)

// CreateOrderInputDTO is the data transfer object for creating an order.
type CreateOrderInputDTO struct {
	Data    string `json:"Data"`
	OrderID int    `json:"OrderId"`
	Status  string `json:"Status"`
}

// CreateOrderOutputDTO is the data transfer object for the result of creating an order.
type CreateOrderOutputDTO struct {
	Data    string `json:"Data"`
	OrderID int    `json:"OrderId"`
	Status  string `json:"Status"`
}

// CreateOrderUseCase is the use case for creating an order.
type CreateOrderUseCase struct {
	MessageQueue repository.OrderMessageQueue
}

// NewCreateOrderUseCase creates a new CreateOrderUseCase.
func NewCreateOrderUseCase(messageQueue repository.OrderMessageQueue) *CreateOrderUseCase {
	return &CreateOrderUseCase{MessageQueue: messageQueue}
}

// Execute executes the use case.
func (uc *CreateOrderUseCase) Execute(input CreateOrderInputDTO) (*CreateOrderOutputDTO, error) {
	order := entity.Order{
		Data:      input.Data,
		OrderID:   input.OrderID,
		Status:    input.Status,
		Paid:      false, // Default to false on creation
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := uc.MessageQueue.Send(&order)
	if err != nil {
		return nil, err
	}

	output := &CreateOrderOutputDTO{
		Data:    order.Data,
		OrderID: order.OrderID,
		Status:  order.Status,
	}

	return output, nil
}
