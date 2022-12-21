package services

import (
	"example.com/m/domain"
	"example.com/m/internal/repository"
)

// BalanceService struct
type BalanceService struct{}

// NewBalanceService return pointer to
// balance service struct with all methods
func NewBalanceService() *BalanceService {
	return &BalanceService{}
}

var balanceRepo = repository.NewBalanceRepo()

func (bs *BalanceService) Get(userID string) (domain.Balance, error) {
	balanceModel, err := balanceRepo.Get(userID)

	if err != nil {
		return domain.Balance{}, err
	}
	return balanceModel, nil
}
