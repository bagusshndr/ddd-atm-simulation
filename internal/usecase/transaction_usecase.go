package usecase

import "ddd-atm-simulation/internal/aggregate"

type TransactionUsecase interface {
	CreateTransaction(transcation aggregate.Transaction) error
	GetTransaction() (res aggregate.Transactions, err error)
}
