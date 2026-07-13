package main

import (
    "database/sql"
    "testing"
    _ "github.com/lib/pq"
	store "github.com/Fletchers97/ledger-core/db"
	"context"
)

func TestConnectToDB(t *testing.T) {
    connStr := "postgresql://root:secretpassword@localhost:5432/ledger?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        t.Fatal("Failed to open the connection:", err)
    }

    err = db.Ping() // This will attempt to connect to the database and verify the connection
    if err != nil {
        t.Fatal("Failed to reach the database:", err)
    }

    t.Log("Success! We have connected to the database.")
}

func TestCreateAccount(t *testing.T) {
	connStr := "postgresql://root:secretpassword@localhost:5432/ledger?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal("Failed to open the connection:", err)
	}

	params := map[string]interface{}{
		"Owner": "Bogdan",
		"Balance":    int64(100),
		"Currency": "USD",
	}

	myStore := store.NewSQLStore(db)
	acc, err := myStore.CreateAccount(context.Background(), store.CreateAccountParams{
		ID:       params["Owner"].(string),
		Balance:  params["Balance"].(int64),
		Currency: params["Currency"].(string),
	})
	if err != nil {
		t.Fatal("Failed to create account:", err)
	}

	if acc.ID == "" {
		t.Fatal("Expected account ID to be set, but it was empty")
	}

	if acc.Balance != 100 {
		t.Fatalf("Expected balance to be 100, but got %d", acc.Balance)
	}

	t.Log("Success! Account created with ID:", acc.ID)

}