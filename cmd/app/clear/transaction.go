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
backLogin:
	fmt.Printf("\nEnter option for login or create account!\n\n")
	fmt.Printf("1. Sign In\n")
	fmt.Printf("2. Create Acoount\n")
	fmt.Scan(&choiceUser)

	switch choiceUser {
	case 1:
	backSignin:
		fmt.Printf("\nPlease enter your id: ")
		fmt.Scan(&t.reqUser.ID)
		getUser, err := t.usecaseUser.GetUserByID(t.reqUser.ID)
		if err != nil {
			fmt.Printf("User not found")
			goto backSignin
		}

		fmt.Printf("\nEnter any option to be served %s !\n\n", getUser[0].Name)
		fmt.Printf("1. Check Amount\n")
		fmt.Printf("2. Desposit\n")
		fmt.Printf("3. Withdraw\n")
		fmt.Printf("4. Transfer\n\n")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("Your balance is: $%.2f", getUser[0].Amount)

			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
			fmt.Scan(&anotherTransaction)
			switch anotherTransaction {
			case 1:
				goto backSignin
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 2:
			// Deposit
			fmt.Printf("Please enter amount to deposit: ")
			fmt.Scan(&t.reqTransaction.Nominal)

			a, _ := aggregate.NewTransactions(getUser[0].ID, 1, 0, t.reqTransaction.Nominal)
			t.usecaseTransaction.CreateTransaction(a)
			getTransaction, _ := t.usecaseUser.GetUserByID(t.reqUser.ID)

			fmt.Printf("Thank you for deposit, new balance is: $%.2f", getTransaction[0].Amount)
			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
			fmt.Scan(&anotherTransaction)

			switch anotherTransaction {
			case 1:
				goto backSignin
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 3:
		backWithdraw:
			// Withdraw
			fmt.Printf("Please enter amount to withdraw: ")
			fmt.Scan(&t.reqTransaction.Nominal)

			getAmount, _ := t.usecaseUser.GetUserByID(t.reqUser.ID)
			if getAmount[0].Amount < t.reqTransaction.Nominal {
				fmt.Printf("Your balance is not enought to withdraw")
				fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					goto backSignin
				default:
					goto backWithdraw
				}
			}

			a, _ := aggregate.NewTransactions(getUser[0].ID, 2, 0, t.reqTransaction.Nominal)
			t.usecaseTransaction.CreateTransaction(a)
			getTransaction, _ := t.usecaseUser.GetUserByID(t.reqUser.ID)

			fmt.Printf("Success To Withdraw, new balance is: $%.2f", getTransaction[0].Amount)
			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
			fmt.Scan(&anotherTransaction)
			switch anotherTransaction {
			case 1:
				goto backSignin
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 4:
			// Transfer
		backUser:
			fmt.Printf("Please enter receive id: ")
			fmt.Scan(&t.reqTransaction.UserReceiveID)

			userReceive, err := t.usecaseUser.GetUserByID(t.reqTransaction.UserReceiveID)
			if err != nil {
				fmt.Printf("Receive ID not found \n")
				fmt.Printf("\nDo you want another transaction?\nPress 1 to input recieve id, 2 back to menu and 3 to logout\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					goto backUser
				case 2:
					goto backSignin
				default:
					fmt.Println("\nThanks for using our service!!! \nHave a nice day \n\n\n\n\n ", getUser[0].Name)
					goto backLogin
				}
			}

			if userReceive[0].ID == t.reqUser.ID {
				fmt.Printf("Cannot transfer to your own id \n")
				fmt.Printf("\nDo you want another transaction?\nPress 1 to input recieve id, 2 back to menu and 3 to logout\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					goto backUser
				case 2:
					goto backSignin
				default:
					fmt.Println("\nThanks for using our service!!! \nHave a nice day \n\n\n\n\n ", getUser[0].Name)
					goto backLogin
				}
			}

		backTransfer:
			fmt.Print("Please enter amount to transfer: ")
			fmt.Scan(&t.reqTransaction.Nominal)

			getAmount, _ := t.usecaseUser.GetUserByID(t.reqUser.ID)
			if getAmount[0].Amount < t.reqTransaction.Nominal {
				fmt.Printf("Your balance is not enought to transfer")
				fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					goto backSignin
				default:
					goto backTransfer
				}
			}

			a, _ := aggregate.NewTransactions(getUser[0].ID, 2, t.reqTransaction.UserReceiveID, t.reqTransaction.Nominal)
			errTransaction := t.usecaseTransaction.CreateTransaction(a)
			if errTransaction != nil {
				fmt.Printf("Failed to transfer")
			}

			fmt.Print("Success to transfer ", userReceive[0].Name)

			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
			fmt.Scan(&anotherTransaction)
			switch anotherTransaction {
			case 1:
				goto backSignin
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 5:
			// Transfer
			fmt.Println("\nThanks for using our service!!! \nHave a nice day ", getUser[0].Name)
			goto backLogin
		}

	case 2:
		fmt.Printf("Please enter your name: ")
		fmt.Scan(&t.reqUser.Name)

		a, _ := aggregate.NewUsers(t.reqUser.Name, 0)
		user, _ := t.usecaseUser.CreateUser(a)
		getUser, _ := t.usecaseUser.GetUserByID(user)

		fmt.Printf("Congratulation %s your account success to create ", getUser[0].Name)
		fmt.Printf("\nEnter any option to be served %s !\n\n", getUser[0].Name)
		fmt.Printf("1. Check Amount\n")
		fmt.Printf("2. Desposit\n")
		fmt.Printf("3. Withdraw\n")
		fmt.Printf("4. Transfer\n\n")
		fmt.Scan(&choice)
	back2:
		switch choice {
		case 1:
			fmt.Printf("Your balance is: $%.2f", getUser[0].Amount)
			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
			fmt.Scan(&anotherTransaction)

			switch anotherTransaction {
			case 1:
				fmt.Printf("\nEnter any option to be served %s !\n\n", getUser[0].Name)
				fmt.Printf("1. Check Amount\n")
				fmt.Printf("2. Desposit\n")
				fmt.Printf("3. Withdraw\n")
				fmt.Printf("4. Transfer\n")
				fmt.Printf("5. Logout\n\n")
				fmt.Scan(&choice)
				goto back2
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 2:
			// Deposit
			fmt.Printf("Please enter amount to deposit: ")
			fmt.Scan(&t.reqTransaction.Nominal)

			a, _ := aggregate.NewTransactions(getUser[0].ID, 1, 0, t.reqTransaction.Nominal)
			t.usecaseTransaction.CreateTransaction(a)

			getUserTransaction, _ := t.usecaseUser.GetUserByID(getUser[0].ID)
			fmt.Printf("Thank you for deposit, new balance is: $%.2f", getUserTransaction[0].Amount)
			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
			fmt.Scan(&anotherTransaction)

			switch anotherTransaction {
			case 1:
				fmt.Printf("\nEnter any option to be served %s !\n\n", getUserTransaction[0].Name)
				fmt.Printf("1. Check Amount\n")
				fmt.Printf("2. Desposit\n")
				fmt.Printf("3. Withdraw\n")
				fmt.Printf("4. Transfer\n")
				fmt.Printf("5. Logout\n\n")
				fmt.Scan(&choice)
				goto back2
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}

		case 3:
			// Withdraw
		backWithdraw2:
			fmt.Printf("Please enter amount to withdraw: ")
			fmt.Scan(&t.reqTransaction.Nominal)
			if getUser[0].Amount < t.reqTransaction.Nominal {
				fmt.Printf("Your balance is not enought to withdraw")
				fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					fmt.Printf("\nEnter any option to be served %s !\n\n", getUser[0].Name)
					fmt.Printf("1. Check Amount\n")
					fmt.Printf("2. Desposit\n")
					fmt.Printf("3. Withdraw\n")
					fmt.Printf("4. Transfer\n")
					fmt.Printf("5. Logout\n\n")
					fmt.Scan(&choice)
					goto back2
				default:
					goto backWithdraw2
				}
			}

			a, _ := aggregate.NewTransactions(getUser[0].ID, 2, 0, t.reqTransaction.Nominal)
			t.usecaseTransaction.CreateTransaction(a)

			getUserTransaction, _ := t.usecaseUser.GetUserByID(getUser[0].ID)

			fmt.Printf("Success To Withdraw, new balance is: $%.2f", getUserTransaction[0].Amount)
			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
			fmt.Scan(&anotherTransaction)
			switch anotherTransaction {
			case 1:
				fmt.Printf("\nEnter any option to be served %s !\n\n", getUserTransaction[0].Name)
				fmt.Printf("1. Check Amount\n")
				fmt.Printf("2. Desposit\n")
				fmt.Printf("3. Withdraw\n")
				fmt.Printf("4. Transfer\n")
				fmt.Printf("5. Logout\n\n")
				fmt.Scan(&choice)
				goto back2
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 4:
			// Transfer
		backUser2:
			fmt.Printf("Please enter receive id: ")
			fmt.Scan(&t.reqTransaction.UserReceiveID)

			userReceive, err := t.usecaseUser.GetUserByID(t.reqTransaction.UserReceiveID)
			if err != nil {
				fmt.Printf("Receive ID not found \n")
				fmt.Printf("\nDo you want another transaction?\nPress 1 to input recieve id, 2 back to menu and 3 to logout\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					goto backUser2
				case 2:
					fmt.Printf("\nEnter any option to be served %s !\n\n", getUser[0].Name)
					fmt.Printf("1. Check Amount\n")
					fmt.Printf("2. Desposit\n")
					fmt.Printf("3. Withdraw\n")
					fmt.Printf("4. Transfer\n")
					fmt.Printf("5. Logout\n\n")
					fmt.Scan(&choice)
					goto back2
				default:
					fmt.Println("\nThanks for using our service!!! \nHave a nice day \n\n\n\n\n ", getUser[0].Name)
					goto backLogin
				}
			}
			if userReceive[0].ID == getUser[0].ID {
				fmt.Printf("Cannot transfer to your own id \n")
				fmt.Printf("\nDo you want another transaction?\nPress 1 to input recieve id, 2 back to menu and 3 to logout\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					goto backUser2
				case 2:
					fmt.Printf("\nEnter any option to be served %s !\n\n", getUser[0].Name)
					fmt.Printf("1. Check Amount\n")
					fmt.Printf("2. Desposit\n")
					fmt.Printf("3. Withdraw\n")
					fmt.Printf("4. Transfer\n")
					fmt.Printf("5. Logout\n\n")
					fmt.Scan(&choice)
					goto back2
				default:
					fmt.Println("\nThanks for using our service!!! \nHave a nice day \n\n\n\n\n ", getUser[0].Name)
					goto backLogin
				}
			}

			fmt.Print("Please enter amount to transfer: ")
			fmt.Scan(&t.reqTransaction.Nominal)
		backTransfer2:
			getAmount, _ := t.usecaseUser.GetUserByID(t.reqUser.ID)
			if getAmount[0].Amount < t.reqTransaction.Nominal {
				fmt.Printf("Your balance is not enought to transfer")
				fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
				fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					fmt.Printf("\nEnter any option to be served %s !\n\n", getUser[0].Name)
					fmt.Printf("1. Check Amount\n")
					fmt.Printf("2. Desposit\n")
					fmt.Printf("3. Withdraw\n")
					fmt.Printf("4. Transfer\n")
					fmt.Printf("5. Logout\n\n")
					fmt.Scan(&choice)
					goto back2
				default:
					goto backTransfer2
				}
			}

			a, _ := aggregate.NewTransactions(getUser[0].ID, 2, t.reqTransaction.UserReceiveID, t.reqTransaction.Nominal)

			errTransaction := t.usecaseTransaction.CreateTransaction(a)
			if errTransaction != nil {
				fmt.Printf("Failed to transfer")
			}

			fmt.Print("Success to transfer ", userReceive[0].Name)

			fmt.Printf("\nDo you want another transaction?\nPress 1 to proceed and 2 to logout\n\n")
			fmt.Scan(&anotherTransaction)
			switch anotherTransaction {
			case 1:
				fmt.Printf("\nEnter any option to be served %s !\n\n", getUser[0].Name)
				fmt.Printf("1. Check Amount\n")
				fmt.Printf("2. Desposit\n")
				fmt.Printf("3. Withdraw\n")
				fmt.Printf("4. Transfer\n")
				fmt.Printf("5. Logout\n\n")
				fmt.Scan(&choice)
				goto back2
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		case 5:
			// Transfer
			fmt.Println("\nThanks for using our service!!! \nHave a nice day ", getUser[0].Name)
			goto backLogin
		}
	}

}

func NewTransaction(usecaseUser usecase.UserUsecase, usecaseTransaction usecase.TransactionUsecase) *transactionCobra {
	return &transactionCobra{
		usecaseUser:        usecaseUser,
		usecaseTransaction: usecaseTransaction,
	}
}
