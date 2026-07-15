package main

import (
    "database/sql"
    "testing"
    _ "github.com/lib/pq"
	store "github.com/Fletchers97/ledger-core/db"
	"context"
	"github.com/stretchr/testify/require"
	"fmt"
	"time"
	"os"
	"log"

)


var testQueries *store.Queries
func TestMain(m *testing.M) {
	connStr := "postgresql://root:secretpassword@localhost:5432/ledger?sslmode=disable"
    conn, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    // 2. Инициализируем глобальную переменную здесь!
    testQueries = store.New(conn)

    // 3. Запускаем все тесты
    os.Exit(m.Run())
}

func TestCreateAccount(t *testing.T) {
	connStr := "postgresql://root:secretpassword@localhost:5432/ledger?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal("Failed to open the connection:", err)
	}

	randomID := fmt.Sprintf("acc_%d", time.Now().UnixNano())
	params := map[string]interface{}{
		"Owner": randomID,
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

func TestGetAccount(t *testing.T) {
    randomID := fmt.Sprintf("acc_%d", time.Now().UnixNano())
    
    account, err := testQueries.CreateAccount(context.Background(), store.CreateAccountParams{
        ID:       randomID,
        Balance:  100,
        Currency: "USD",
    })
    require.NoError(t, err)

    //Get the account from the database
    account2, err := testQueries.GetAccount(context.Background(), account.ID)
    require.NoError(t, err)
    require.NotEmpty(t, account2)

    //Check that the retrieved account matches the created account
    require.Equal(t, account.ID, account2.ID)
    require.Equal(t, account.Balance, account2.Balance)
}

func TestUpdateAccount(t *testing.T) {
	randomID := fmt.Sprintf("acc_%d", time.Now().UnixNano())	
	account, err := testQueries.CreateAccount(context.Background(), store.CreateAccountParams{
		ID:       randomID,
		Balance:  500,
		Currency: "USD",
	})
	require.NoError(t, err)

	account2, err := testQueries.UpdateAccount(context.Background(), store.UpdateAccountParams{
		ID:     account.ID,
		Amount: 1000,
	})
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, int64(1000), account2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	randomID := fmt.Sprintf("acc_%d", time.Now().UnixNano())
	account, err := testQueries.CreateAccount(context.Background(), store.CreateAccountParams{
		ID:       randomID,
		Balance:  100,
		Currency: "USD",
	})
	require.NoError(t, err)

	err = testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}
