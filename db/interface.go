package db

import (
	"context"
	"github.com/Fletchers97/ledger-core/db/models"
)

type Store interface {
	TransferFunds(ctx context.Context, fromID, toID string, amount int64) error
	CreateAccount(ctx context.Context, arg CreateAccountParams) (models.Account, error)
	GetAccount(ctx context.Context, id string) (models.Account, error)
	CountAccounts(ctx context.Context) (int64, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (models.Account, error)
}