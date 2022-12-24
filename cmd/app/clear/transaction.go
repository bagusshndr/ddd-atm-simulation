package clear

import (
	"ddd-atm-simulation/internal/aggregate"
	"ddd-atm-simulation/internal/usecase"
	"fmt"
	"os"
	"os/exec"
)

func CallClear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

type transactionCobra struct {
	usecaseUser        usecase.UserUsecase
	usecaseTransaction usecase.TransactionUsecase
	reqTransaction     ReqTransaction
	reqUser            ReqUser
}

type ReqTransaction struct {
	UserID        uint64
	Flag          int
	UserReceiveID uint64
	Nominal       float64
}

type ReqUser struct {
	ID     uint64
	Name   string
	Amount float64
}

var anotherTransaction int

func (t *transactionCobra) Transaction() {

	var choiceUser int
	var choice int

	fmt.Printf("\nEnter option for login or create account!\n\n")
	fmt.Printf("1. Sign In\n")
	fmt.Printf("2. Create Acoount\n")
	fmt.Scan(&choiceUser)

	switch choiceUser {
	case 1:
		fmt.Printf("Please enter your id: ")
		fmt.Scan(&t.reqUser.ID)
		getUser, err := t.usecaseUser.GetUserByID(t.reqUser.ID)
		if err != nil {
			fmt.Printf("Please enter your id: ")
			fmt.Scan(&t.reqUser.ID)
		}

		fmt.Printf("\nEnter any option to be served!\n\n")
		fmt.Printf("1. Check Amount\n")
		fmt.Printf("2. Desposit\n")
		fmt.Printf("3. Withdraw\n")
		fmt.Printf("4. Transfer\n\n")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: $%.2f", getUser[0].Amount)

			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
			fmt.Scan(&anotherTransaction)
			switch anotherTransaction {
			case 1:
				fmt.Printf("\nEnter any option to be served!\n\n")
				fmt.Printf("1. Check Amount\n")
				fmt.Printf("2. Desposit\n")
				fmt.Printf("3. Withdraw\n")
				fmt.Scan(&choice)
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 2:
			// Deposit
			fmt.Printf("Please enter amount to deposit: ")
			fmt.Scan(&t.reqTransaction.Nominal)

			a, _ := aggregate.NewTransactions(getUser[0].ID, 1, 0, t.reqTransaction.Nominal)
			t.usecaseTransaction.CreateTransaction(a)

			fmt.Printf("Thank you for deposit, new balance is: $%.2f", getUser[0].Amount)
			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
			fmt.Scan(&anotherTransaction)

			switch anotherTransaction {
			case 1:
				fmt.Printf("\nEnter any option to be served!\n\n")
				fmt.Printf("1. Check Amount\n")
				fmt.Printf("2. Desposit\n")
				fmt.Printf("3. Withdraw\n")
				fmt.Scan(&choice)
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 3:
			// Withdraw
			fmt.Printf("Please enter amount to withdraw: ")
			fmt.Scan(&t.reqTransaction.Nominal)

			a, _ := aggregate.NewTransactions(getUser[0].ID, 2, 0, t.reqTransaction.Nominal)
			t.usecaseTransaction.CreateTransaction(a)

			fmt.Printf("Success To Withdraw, new balance is: $%.2f", getUser[0].Amount)
			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
			fmt.Scan(&anotherTransaction)
			switch anotherTransaction {
			case 1:
				fmt.Printf("\nEnter any option to be served!\n\n")
				fmt.Printf("1. Check Amount\n")
				fmt.Printf("2. Desposit\n")
				fmt.Printf("3. Withdraw\n")
				fmt.Scan(&choice)
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		}

	case 2:
		fmt.Printf("Please enter your name: ")
		fmt.Scan(&t.reqUser.Name)

		a, _ := aggregate.NewUsers(t.reqUser.Name, 0)
		user, _ := t.usecaseUser.CreateUser(a)
		getUser, _ := t.usecaseUser.GetUserByID(user)

		fmt.Printf("Congratulation %s your account success to create", getUser[0].Name)
		fmt.Printf("Do you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
		fmt.Scan(&anotherTransaction)

		switch anotherTransaction {
		case 1:
			fmt.Printf("\nEnter any option to be served!\n\n")
			fmt.Printf("1. Check Amount\n")
			fmt.Printf("2. Desposit\n")
			fmt.Printf("3. Withdraw\n")
			fmt.Printf("4. Transfer\n\n")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				fmt.Printf("Your balance is: $%.2f", getUser[0].Amount)

				fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					fmt.Printf("\nEnter any option to be served!\n\n")
					fmt.Printf("1. Check Amount\n")
					fmt.Printf("2. Desposit\n")
					fmt.Printf("3. Withdraw\n")
					fmt.Scan(&choice)
				default:
					fmt.Println("\nThanks for using our service!!! \nHave a nice day")
				}
			case 2:
				// Deposit
				fmt.Printf("Please enter amount to deposit: ")
				fmt.Scan(&t.reqTransaction.Nominal)

				a, _ := aggregate.NewTransactions(getUser[0].ID, 1, 0, t.reqTransaction.Nominal)
				t.usecaseTransaction.CreateTransaction(a)

				fmt.Printf("Thank you for deposit, new balance is: $%.2f", getUser[0].Amount)
				fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
				fmt.Scan(&anotherTransaction)

				switch anotherTransaction {
				case 1:
					fmt.Printf("\nEnter any option to be served!\n\n")
					fmt.Printf("1. Check Amount\n")
					fmt.Printf("2. Desposit\n")
					fmt.Printf("3. Withdraw\n")
					fmt.Scan(&choice)
				default:
					fmt.Println("\nThanks for using our service!!! \nHave a nice day")
				}
			case 3:
				// Withdraw
				fmt.Printf("Please enter amount to withdraw: ")
				fmt.Scan(&t.reqTransaction.Nominal)

				a, _ := aggregate.NewTransactions(getUser[0].ID, 2, 0, t.reqTransaction.Nominal)
				t.usecaseTransaction.CreateTransaction(a)

				fmt.Printf("Success To Withdraw, new balance is: $%.2f", getUser[0].Amount)
				fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					fmt.Printf("\nEnter any option to be served!\n\n")
					fmt.Printf("1. Check Amount\n")
					fmt.Printf("2. Desposit\n")
					fmt.Printf("3. Withdraw\n")
					fmt.Scan(&choice)
				default:
					fmt.Println("\nThanks for using our service!!! \nHave a nice day")
				}
			}
		default:
			fmt.Println("\nThanks for using our service!!! \nHave a nice day")
		}

	}

	switch choice {
	case 1:
		fmt.Printf("Please enter your account id: ")
		fmt.Scan(&t.reqTransaction.UserID)

		user, _ := t.usecaseUser.GetUserByID(t.reqTransaction.UserID)

		fmt.Printf("Success To Withdraw, new balance is: $%.2f", user[0].Amount)
	case 2:
		// Deposit
		fmt.Printf("Please enter your account id: ")
		fmt.Scan(&t.reqTransaction.UserID)

		user, _ := t.usecaseUser.GetUserByID(t.reqTransaction.UserID)
		fmt.Printf("Hallo %s \n", user[0].Name)

		fmt.Printf("Please enter amount to deposit: ")
		fmt.Scan(&t.reqTransaction.Nominal)

		a, _ := aggregate.NewTransactions(t.reqTransaction.UserID, 1, 0, t.reqTransaction.Nominal)
		t.usecaseTransaction.CreateTransaction(a)

		fmt.Printf("Thank you for depositing, new balance is: $%.2f", user[0].Amount)

	case 3:
		// Withdraw
		fmt.Printf("Please enter your account id: ")
		fmt.Scan(&t.reqTransaction.UserID)

		user, _ := t.usecaseUser.GetUserByID(t.reqTransaction.UserID)
		fmt.Printf("Hallo %s \n", user[0].Name)
		fmt.Printf("Please enter amount to withdraw: ")
		fmt.Scan(&t.reqTransaction.Nominal)

		a, _ := aggregate.NewTransactions(t.reqTransaction.UserID, 2, 0, t.reqTransaction.Nominal)
		t.usecaseTransaction.CreateTransaction(a)

		fmt.Printf("Success To Withdraw, new balance is: $%.2f", user[0].Amount)
	}
}

func NewTransaction(usecaseUser usecase.UserUsecase, usecaseTransaction usecase.TransactionUsecase) *transactionCobra {
	return &transactionCobra{
		usecaseUser:        usecaseUser,
		usecaseTransaction: usecaseTransaction,
	}
}
