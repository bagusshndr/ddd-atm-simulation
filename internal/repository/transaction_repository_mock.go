package repository

import (
	"ddd-atm-simulation/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) GetTransaction() (res aggregate.Transactions, err error) {
	args := m.Called()

	return args.Get(0).(aggregate.Transactions), args.Error(1)
}

func (m *TransactionRepositoryMock) CreateTransaction(transaction aggregate.Transaction) (uint64, error) {
	args := m.Called(transaction)

	return args.Get(0).(uint64), args.Error(1)
}

func (m *TransactionRepositoryMock) GetTransactionByID(id uint64) (res aggregate.Transactions, err error) {
	args := m.Called(id)

	return args.Get(0).(aggregate.Transactions), args.Error(1)
}

func (m *TransactionRepositoryMock) DeleteTransaction(id uint64) error {
	args := m.Called(id)

	return args.Error(0)
}
