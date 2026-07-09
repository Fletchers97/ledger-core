package service

import (
	"context"
	"github.com/Fletchers97/ledger-core/internal/domain"
	"github.com/Fletchers97/ledger-core/internal/repository"
)

// AccountService is responsible for account operations
type AccountService struct {
	repo repository.AccountRepository
}

// NewAccountService creates a new service
func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

// CreateAccount creates a new account
func (s *AccountService) CreateAccount(ctx context.Context, acc *domain.Account) error {
	// Здесь можно добавить проверку: например, нельзя создать счет с отрицательным балансом
	if acc.Balance < 0 {
		return domain.ErrInvalidBalance // Потребуется добавить этот тип ошибки в domain
	}
	return s.repo.Create(ctx, acc)
}