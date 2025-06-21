

import (
	"GoCleanArch/internal/domain/entity"
	"GoCleanArch/internal/domain/repository"
)

// GetAllOrdersUseCase retrieves all orders.
type GetAllOrdersUseCase struct {
	OrderRepository repository.OrderRepository
}

func NewGetAllOrdersUseCase(orderRepo repository.OrderRepository) *GetAllOrdersUseCase {
	return &GetAllOrdersUseCase{OrderRepository: orderRepo}
}

func (uc *GetAllOrdersUseCase) Execute() ([]*entity.Order, error) {
	return uc.OrderRepository.GetAll()
}
