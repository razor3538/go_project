package services

import (
	"example.com/m/domain"
	"example.com/m/internal/repository"
	"example.com/m/tools"
	"strconv"
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

		accrualFloat, _ := strconv.ParseFloat(accrual, 32)
		println(strconv.ParseFloat(accrual, 32))
		if err != nil {
			return []domain.Order{}, err
		}

		orders[i].Accrual = accrual
		orders[i].Status = orderStatus
		_, err = orderRepo.ChangeStatus(order)
		if err != nil {
			return []domain.Order{}, err
		}
		_, err = balanceRepo.Add(userID, accrualFloat)
		if err != nil {
			return []domain.Order{}, err
		}
	}

	return orders, nil
}
