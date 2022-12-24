package usecase

import "ddd-to-do-list/internal/aggregate"

type TransactionUsecase interface {
	CreateTransaction(transcation aggregate.Transaction) error
	GetTransaction() (res aggregate.Transactions, err error)
}
