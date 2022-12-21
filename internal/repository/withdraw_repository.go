package repository

import (
	"errors"
	"example.com/m/config"
	"example.com/m/domain"
)

// WithdrawRepo struct
type WithdrawRepo struct{}

// NewWithdrawRepo return pointer to withdraw repository
// with all methods
func NewWithdrawRepo() *WithdrawRepo {
	return &WithdrawRepo{}
}

func (wr *WithdrawRepo) Pay(withdrawal domain.Withdrawals) (domain.Withdrawals, error) {
	var balance domain.Balance

	err := config.DB.
		Unscoped().
		Table("balances as b").
		Where("b.user_id = ?", withdrawal.UserID).
		Scan(&balance).Error

	if withdrawal.Sum > balance.Current {
		return domain.Withdrawals{}, errors.New("сумма для снятия не достаточна")
	}
	println(withdrawal.Order)

	if err := config.DB.
		Create(&withdrawal).
		Error; err != nil {
		return domain.Withdrawals{}, err
	}
	return withdrawal, err
}

func (wr *WithdrawRepo) Get(userID string) ([]domain.Withdrawals, error) {
	var balance []domain.Withdrawals
	err := config.DB.
		Unscoped().
		Table("withdrawals as b").
		Where("b.user_id = ?", userID).
		Scan(&balance).Error

	return balance, err
}
