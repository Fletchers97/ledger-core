package repository

import (
	"context"

	"github.com/Fletchers97/ledger-core/internal/domain"
)


type AccountRepository interface {
	Create(ctx context.Context, account *domain.Account) error
	GetByID(ctx context.Context, id string) (*domain.Account, error)
	UpdateBalance(ctx context.Context, id string, newBalance int64) error
}
