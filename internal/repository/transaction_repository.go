package repository

import "ddd-atm-simulation/internal/aggregate"

type TransactionRepository interface {
	GetTransaction() (res aggregate.Transactions, err error)
	CreateTransaction(Transaction aggregate.Transaction) (uint64, error)
	GetTransactionByID(id uint64) (res aggregate.Transactions, err error)
	DeleteTransaction(id uint64) error
}
