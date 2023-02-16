package services

import (
	"example.com/m/domain"
	"example.com/m/internal/repository"
	"example.com/m/tools"
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
	orders, err := orderRepo.GetByUser(userID)

	if err != nil {
		return []domain.Order{}, err
	}

	for i, order := range orders {
		orderStatus, accrual, err := tools.OrderProcessed(order.Number)

		if err != nil {
			println(1)
			return orders, err
		}

		orders[i].Accrual = accrual
		orders[i].Status = orderStatus
		_, err = orderRepo.ChangeStatus(order)
		if err != nil {
			return orders, err
		}
		_, err = balanceRepo.Add(userID, accrual)
		if err != nil {
			return orders, err
		}
	}

	return orders, nil
}
