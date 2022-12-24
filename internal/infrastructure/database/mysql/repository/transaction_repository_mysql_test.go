package repository

import (
	"ddd-atm-simulation/internal/aggregate"
	"ddd-atm-simulation/internal/enum"
	"ddd-atm-simulation/internal/repository"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type transactionRepositoryMysqlTest struct {
	suite.Suite
	mock                  sqlmock.Sqlmock
	transactionRepository repository.TransactionRepository
}

func (t *transactionRepositoryMysqlTest) TestGetTransaction() {
	transaction := aggregate.RebuildTransaction(1, 1, enum.Deposit, 1, 100)
	query := `SELECT id, user_id, flag, nominal FROM transactions`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"user_id",
			"flag",
			"nominal",
		}).AddRow(
			transaction.ID,
			transaction.UserID,
			transaction.Flag,
			transaction.Nominal,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		actualProduct, err := t.transactionRepository.GetTransaction()

		t.NotNil(actualProduct)
		t.NoError(err)
	})
	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New(""))

		actualActivity, err := t.transactionRepository.GetTransaction()

		t.Nil(actualActivity)
		t.Error(err)
	})
}
func (t *transactionRepositoryMysqlTest) TestGetTransactionByID() {
	transaction := aggregate.RebuildTransaction(1, 1, enum.Deposit, 1, 100)
	query := `SELECT id, user_id, flag, nominal FROM transactions WHERE id = ? LIMIT 1`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"user_id",
			"flag",
			"nominal",
		}).AddRow(
			transaction.ID,
			transaction.UserID,
			transaction.Flag,
			transaction.Nominal,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(transaction.ID).WillReturnRows(rows)

		actualProduct, err := t.transactionRepository.GetTransactionByID(transaction.ID)

		t.NotNil(actualProduct)
		t.NoError(err)
	})
	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(transaction.ID).WillReturnError(errors.New(""))

		actualTransaction, err := t.transactionRepository.GetTransactionByID(transaction.ID)

		t.Nil(actualTransaction)
		t.Error(err)
	})
}

func TestProductRepositoryMySQL(t *testing.T) {
	db, mock, _ := sqlmock.New()

	suite.Run(t, &transactionRepositoryMysqlTest{
		mock:                  mock,
		transactionRepository: NewMysqlTransactionRepository(db),
	})
}
