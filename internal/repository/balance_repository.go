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

func (br *BalanceRepo) Add(userID string, current float64) (domain.Balance, error) {
	var balance domain.Balance

	err := config.DB.
		Unscoped().
		Table("balances as b").
		Where("b.user_id = ?", userID).
		Scan(&balance).Error

	//if err != nil {
	//	balance.Current = current
	//	if err := config.DB.
	//		Create(&balance).
	//		Error; err != nil {
	//		return domain.Balance{}, err
	//	}
	//} else {
	//	balance.Current += current
	//	if err := config.DB.
	//		Update("current", balance.Current).
	//		Error; err != nil {
	//		return domain.Balance{}, err
	//	}
	//}

	return balance, err
}
