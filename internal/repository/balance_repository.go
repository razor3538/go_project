package repository

import (
	"example.com/m/config"
	"example.com/m/domain"
)

// BalanceRepo struct
type BalanceRepo struct{}

// NewBalanceRepo return pointer to balance repository
// with all methods
func NewBalanceRepo() *BalanceRepo {
	return &BalanceRepo{}
}

func (br *BalanceRepo) Get(userID string) (domain.Balance, error) {
	var balance domain.Balance
	err := config.DB.
		Unscoped().
		Table("balances as b").
		Where("b.user_id = ?", userID).
		Scan(&balance).Error

	return balance, err
}
