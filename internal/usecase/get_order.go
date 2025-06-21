package usecase

import (
	"GoCleanArch/internal/domain/repository"
	"log"
)

// GetOrderByIDInputDTO is the data transfer object for getting an order by ID.
type GetOrderByIDInputDTO struct {
	OrderID int `json:"orderId"`
}

// GetOrderByIDOutputDTO is the data transfer object for the result of getting an order.
type GetOrderByIDOutputDTO struct {
	Data    string `json:"Data"`
	OrderID int    `json:"OrderId"`
	Status  string `json:"Status"`
	Paid    bool   `json:"Paid"`
}

// GetOrderByIDUseCase is the use case for getting an order by ID.
type GetOrderByIDUseCase struct {
	OrderRepository repository.OrderRepository
}

// NewGetOrderByIDUseCase creates a new GetOrderByIDUseCase.
func NewGetOrderByIDUseCase(orderRepository repository.OrderRepository) *GetOrderByIDUseCase {
	return &GetOrderByIDUseCase{OrderRepository: orderRepository}
}

// Execute executes the use case.
func (uc *GetOrderByIDUseCase) Execute(input GetOrderByIDInputDTO) (*GetOrderByIDOutputDTO, error) {
	log.Printf("Executing GetOrderByIDUseCase with OrderID: %d", input.OrderID)
	order, err := uc.OrderRepository.GetByOrderID(input.OrderID)
	if err != nil {
		return nil, err
	}

	output := &GetOrderByIDOutputDTO{
		Data:    order.Data,
		OrderID: order.OrderID,
		Status:  order.Status,
		Paid:    order.Paid,
	}

	return output, nil
}
