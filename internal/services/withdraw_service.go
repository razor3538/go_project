package services

import (
	"example.com/m/domain"
	"example.com/m/internal/repository"
)

// WithdrawService struct
type WithdrawService struct{}

// NewWithdrawService return pointer to
// withdraw service struct with all methods
func NewWithdrawService() *WithdrawService {
	return &WithdrawService{}
}

var withdrawService = repository.NewWithdrawRepo()

func (ws *WithdrawService) Pay(withdrawal domain.Withdrawals) (domain.Withdrawals, error) {
	result, err := withdrawService.Pay(withdrawal)

	if err != nil {
		return domain.Withdrawals{}, err
	}
	return result, nil
}

func (ws *WithdrawService) GetAllByUser(userID string) ([]domain.Withdrawals, error) {
	withdrawModel, err := withdrawService.Get(userID)

	if err != nil {
		return []domain.Withdrawals{}, err
	}
	return withdrawModel, nil
}
