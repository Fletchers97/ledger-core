package main

import (
	"context"
	"fmt"
	"github.com/Fletchers97/ledger-core/internal/domain"
	"github.com/Fletchers97/ledger-core/internal/repository"
	"github.com/Fletchers97/ledger-core/internal/service"
)

func main() {
	repo := repository.NewInMemoryAccountRepo()
	
	
	svc := service.NewAccountService(repo)

	//Create a new account
	acc := &domain.Account{
		ID:       "1",
		Balance:  1000, // 10 dollars in cents
		Currency: "CAD",
	}

	
	err := svc.CreateAccount(context.Background(), acc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	
	foundAcc, err := svc.GetAccountByID(context.Background(), "1")
	if err != nil {
    	fmt.Println("Error searching:", err)
	} else {
    	fmt.Println("Account found:", foundAcc.ID, "with balance:", foundAcc.Balance)
	}

	fmt.Println("Success! Account created with balance:", acc.Balance)
}
