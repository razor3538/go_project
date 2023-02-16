package services

import (
	"example.com/m/domain"
	"example.com/m/internal/repository"
	"example.com/m/tools"
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

	orderStatus, accrual, err := tools.OrderProcessed(withdrawal.Order)

	if err != nil {
		println(1)
		return domain.Withdrawals{}, err
	}

	order := domain.Order{
		Base:    domain.Base{},
		Number:  withdrawal.Order,
		UserID:  withdrawal.UserID,
		Status:  orderStatus,
		Accrual: accrual,
	}

	_, err = orderRepo.ChangeStatus(order)
	if err != nil {
		return domain.Withdrawals{}, err
	}
	_, err = balanceRepo.Add(withdrawal.UserID.String(), accrual)
	if err != nil {
		return domain.Withdrawals{}, err
	}

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
