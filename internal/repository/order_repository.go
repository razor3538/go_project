package repository

import (
	"example.com/m/config"
	"example.com/m/domain"
)

// OrderRepo struct
type OrderRepo struct{}

// NewOrderRepo return pointer to user repository
// with all methods
func NewOrderRepo() *OrderRepo {
	return &OrderRepo{}
}

func (or *OrderRepo) Add(order domain.Order) (domain.Order, error) {
	if err := config.DB.
		Create(&order).
		Error; err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (or *OrderRepo) ChangeStatus(order domain.Order) (domain.Order, error) {
	if err := config.DB.
		Table("orders as o").
		Where("o.id = ?", order.ID).
		Update("status", order.Status).
		Error; err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (or *OrderRepo) GetByUser(userID string) ([]domain.Order, error) {
	var orders []domain.Order
	err := config.DB.
		Unscoped().
		Table("orders as o").
		Where("o.user_id = ?", userID).
		Order("created_at desc").
		Scan(&orders).Error
	if err != nil {
		println(err.Error())
	}
	return orders, err
}
