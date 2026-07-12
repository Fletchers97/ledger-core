package db

import (
	"context"
	"github.com/Fletchers97/ledger-core/db/models"
)

type Store interface {
	TransferFunds(ctx context.Context, fromID, toID string, amount int64) error
	CreateAccount(ctx context.Context, arg CreateAccountParams) (models.Account, error)
}

type CreateAccountParams struct {
	ID       string
	Balance  int64
	Currency string
}