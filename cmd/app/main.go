package main

import (
	"database/sql"
	"ddd-atm-simulation/internal/infrastructure/database/mysql/model"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	_CobraHandler "ddd-atm-simulation/cmd/app/clear"
	_Repo "ddd-atm-simulation/internal/infrastructure/database/mysql/repository"
	_Ucase "ddd-atm-simulation/internal/usecase"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var balance float32 = 0
// var anotherTransaction int

// func transaction() {

// 	var choice int

// 	fmt.Printf("\nEnter any option to be served!\n\n")
// 	fmt.Printf("1. Withdraw\n")
// 	fmt.Printf("2. Deposit\n")
// 	fmt.Printf("3. Balance\n")
// 	fmt.Printf("4. Transfer\n\n")
// 	fmt.Scan(&choice)

// 	// fmt.Print(choice)
// 	var amountToTransfer float32
// 	var amountToWithDraw float32
// 	var amountToDeposit float32

// 	switch choice {
// 	case 1:
// 		// Withdraw
// 		fmt.Printf("\nPlease enter amount to withdraw: ")
// 		fmt.Scan(&amountToWithDraw)

// 		if amountToWithDraw > balance {
// 			fmt.Printf("There is no insufficient funds in your account")
// 			fmt.Printf("Do you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
// 			fmt.Scan(&anotherTransaction)

// 			switch anotherTransaction {
// 			case 1:
// 				clear.CallClear()
// 				transaction()
// 			default:
// 				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
// 			}
// 		} else {
// 			balance -= amountToWithDraw
// 			fmt.Printf("You have withdrawn $%.2f and your new balance is $%.2f ", amountToWithDraw, balance)
// 			fmt.Printf("\n\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
// 			fmt.Scan(&anotherTransaction)

// 			switch anotherTransaction {
// 			case 1:
// 				clear.CallClear()
// 				transaction()
// 			default:
// 				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
// 			}
// 		}
// 	case 2:
// 		// Deposit
// 		fmt.Printf("\nPlease enter amount to deposit: ")
// 		fmt.Scan(&amountToDeposit)

// 		balance += amountToDeposit

// 		fmt.Printf("Thank you for depositing, new balance is: $%.2f", balance)
// 		fmt.Printf("\n\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
// 		fmt.Scan(&anotherTransaction)

// 		switch anotherTransaction {
// 		case 1:
// 			clear.CallClear()
// 			transaction()
// 		default:
// 			fmt.Println("\nThanks for using our service!!! \nHave a nice day")
// 		}

// 	case 3:
// 		fmt.Printf("\nYour bank balance is: $%.2f", balance)
// 		fmt.Printf("\n\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
// 		fmt.Scan(&anotherTransaction)

// 		switch anotherTransaction {
// 		case 1:
// 			clear.CallClear()
// 			transaction()
// 		default:
// 			fmt.Println("\nThanks for using our service!!! \nHave a nice day")
// 		}
// 	case 4:
// 		// Transfer
// 		fmt.Printf("\nPlease enter amount to transfer: ")
// 		fmt.Scan(&amountToTransfer)

// 		if amountToTransfer > balance {
// 			fmt.Printf("There is no insufficient funds in your account")
// 			fmt.Printf("Do you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
// 			fmt.Scan(&anotherTransaction)

// 			switch anotherTransaction {
// 			case 1:
// 				clear.CallClear()
// 				transaction()
// 			default:
// 				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
// 			}
// 		} else {
// 			balance -= amountToTransfer
// 			fmt.Printf("You have trasnfer to user $%.2f and your new balance is $%.2f ", amountToWithDraw, balance)
// 			fmt.Printf("\n\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
// 			fmt.Scan(&anotherTransaction)

// 			switch anotherTransaction {
// 			case 1:
// 				clear.CallClear()
// 				transaction()
// 			default:
// 				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
// 			}
// 		}
// 	}

// }

func main() {
	// fmt.Println("It is working")
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		if err = godotenv.Load(".env"); err != nil {
			return
		}
	}
	dbHost := os.Getenv(`MYSQL_HOST`)
	dbPort := os.Getenv(`MYSQL_PORT`)
	dbUser := os.Getenv(`MYSQL_USER`)
	dbPass := os.Getenv(`MYSQL_PASSWORD`)
	dbName := os.Getenv(`MYSQL_DBNAME`)

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	ds := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(ds), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.TransactionDTO{}, &model.UserDTO{})

	ar := _Repo.NewMysqlUserRepository(dbConn)
	tr := _Repo.NewMysqlTransactionRepository(dbConn)
	au := _Ucase.NewUserUsecase(ar)
	tu := _Ucase.NewTransactionUsecase(ar, tr)
	handler := _CobraHandler.NewTransaction(au, tu)
	fmt.Printf("\nWelcome to my ATM Simulator\n")
	handler.Transaction()
}
