package db

import (
	"context"
	"database/sql"
	"github.com/Fletchers97/ledger-core/db/models"
)

type SQLStore struct {
	db *sql.DB
}


func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{db: db}
}

func (store *SQLStore) TransferFunds(ctx context.Context, fromID, toID string, amount int64) error {
	// Start a transaction
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Withdraw from the source account
	err = execTransfer(ctx, tx, fromID, -amount)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Credit to the destination account
	err = execTransfer(ctx, tx, toID, amount)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Commit the transaction
	return tx.Commit()
}

func execTransfer(ctx context.Context, tx *sql.Tx, accountID string, amount int64) error {
	query := "UPDATE accounts SET balance = balance + $1 WHERE id = $2"

	_, err := tx.ExecContext(ctx, query, amount, accountID)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) CreateAccount(ctx context.Context, arg CreateAccountParams) (models.Account, error) {
	var account models.Account

	// Insert the new account into the database
	// $1, $2, and $3 are placeholders for the ID, Balance, and Currency values respectively
	// RETURNING * returns the data of the newly created account, which we scan into the account variable
	query := "INSERT INTO accounts (id, balance, currency) VALUES ($1, $2, $3) RETURNING id, balance, currency, created_at"

	err := s.db.QueryRowContext(ctx, query, arg.ID, arg.Balance, arg.Currency).Scan(
		&account.ID,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	return account, err
}