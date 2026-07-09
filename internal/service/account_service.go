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
	if acc.Balance < 0 {
		return domain.ErrInvalidBalance 
	}
	return s.repo.Create(ctx, acc)
}

// GetAccountByID retrieves an account by its ID
func (s *AccountService) GetAccountByID(ctx context.Context, id string) (*domain.Account, error) {
    return s.repo.GetByID(ctx, id)
}