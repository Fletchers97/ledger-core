package db

import "testing"

func TestNewSQLStore(t *testing.T) {
	store := NewSQLStore(nil)
	if store == nil {
		t.Errorf("Expected non-nil store, got nil")
	}

	if store.db != nil {
		t.Errorf("Expected db to be nil, got %v", store.db)
	}
}