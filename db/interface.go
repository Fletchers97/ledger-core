package db

import "context"

type Store interface {
	TransferFunds(ctx context.Context, fromID, toID string, amount int64) error
}