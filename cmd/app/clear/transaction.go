package clear

import (
	"ddd-to-do-list/internal/usecase"
	"fmt"
)

type transactionCobra struct {
	usecaseUser        usecase.UserUsecase
	usecaseTransaction usecase.TransactionUsecase
}

func (t *transactionCobra) Transaction() {

	var choice int

	fmt.Printf("\nEnter any option to be served!\n\n")
	fmt.Printf("1. Check Amount\n")

	switch choice {
	case 1:
		transaction, _ := t.usecaseTransaction.GetTransaction()
		fmt.Println(transaction[0].ID)
		fmt.Printf("There is no insufficient funds in your account")
		fmt.Printf("Do you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
	}
}

func NewTransaction(usecaseUser usecase.UserUsecase, usecaseTransaction usecase.TransactionUsecase) *transactionCobra {
	return &transactionCobra{
		usecaseUser:        usecaseUser,
		usecaseTransaction: usecaseTransaction,
	}
}
