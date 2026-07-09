package repository

import (
	"context"
	"fmt"
	"sync"
	"github.com/Fletchers97/ledger-core/internal/domain"
)

type InMemoryAccountRepo struct {
	mu       sync.RWMutex
	accounts map[string]*domain.Account
}

func NewInMemoryAccountRepo() *InMemoryAccountRepo {
	return &InMemoryAccountRepo{
		accounts: make(map[string]*domain.Account),
	}
}

func (r *InMemoryAccountRepo) Create(ctx context.Context, acc *domain.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.accounts[acc.ID] = acc
	return nil
}

func (r *InMemoryAccountRepo) GetByID(ctx context.Context, id string) (*domain.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	acc, ok := r.accounts[id]
	if !ok {
		return nil, fmt.Errorf("account not found")
	}
	return acc, nil
}

func (r *InMemoryAccountRepo) UpdateBalance(ctx context.Context, id string, newBalance int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if acc, ok := r.accounts[id]; ok {
		acc.Balance = newBalance
		return nil
	}
	return fmt.Errorf("account not found")
}

func (r *InMemoryAccountRepo) Delete(ctx context.Context, id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, ok := r.accounts[id]; !ok {
        return fmt.Errorf("account not found")
    }
    
    delete(r.accounts, id)
    return nil
}