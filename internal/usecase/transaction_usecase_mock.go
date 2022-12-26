package usecase

import (
	"ddd-atm-simulation/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type TransactionUsecaseyMock struct {
	mock.Mock
}

func (m *TransactionUsecaseyMock) GetTransaction() (res aggregate.Transactions, err error) {
	args := m.Called()

	return args.Get(0).(aggregate.Transactions), args.Error(1)
}

func (m *TransactionUsecaseyMock) CreateTransaction(transaction aggregate.Transaction) (uint64, error) {
	args := m.Called(transaction)

	return args.Get(0).(uint64), args.Error(1)
}

func (m *TransactionUsecaseyMock) GetTransactionByID(id uint64) (res aggregate.Transactions, err error) {
	args := m.Called(id)

	return args.Get(0).(aggregate.Transactions), args.Error(1)
}

func (m *TransactionUsecaseyMock) DeleteTransaction(id uint64) error {
	args := m.Called(id)

	return args.Error(0)
}
