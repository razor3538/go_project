package services

import (
	"example.com/m/domain"
	"example.com/m/internal/repository"
)

// OrderService struct
type OrderService struct{}

// NewOrderService return pointer to
// user struct with all methods
func NewOrderService() *OrderService {
	return &OrderService{}
}

var orderRepo = repository.NewOrderRepo()

func (os *OrderService) Add(order domain.Order) (domain.Order, error) {
	orderModel, err := orderRepo.Add(order)

	if err != nil {
		return domain.Order{}, err
	}
	return orderModel, nil
}

func (os *OrderService) GetAllByUser(userID string) ([]domain.Order, error) {
	orderModel, err := orderRepo.GetByUser(userID)

	if err != nil {
		return []domain.Order{}, err
	}
	return orderModel, nil
}
