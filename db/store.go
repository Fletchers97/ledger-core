package db

import (
	"context"
	"database/sql"
)

type SQLStore struct {
	db *sql.DB
}

var _ Store = (*SQLStore)(nil)

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{db: db}
}

func (store *SQLStore) TransferFunds(ctx context.Context, fromID, toID string, amount int64) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = execTransfer(ctx, tx, fromID, -amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = execTransfer(ctx, tx, toID, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func execTransfer(ctx context.Context, tx *sql.Tx, accountID string, amount int64) error {
	_, err := tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, accountID)
	return err
}