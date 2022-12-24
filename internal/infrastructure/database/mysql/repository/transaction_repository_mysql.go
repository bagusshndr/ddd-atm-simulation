package repository

import (
	"database/sql"
	"ddd-atm-simulation/internal/aggregate"
	"ddd-atm-simulation/internal/infrastructure/database/mysql/model"
	"ddd-atm-simulation/internal/repository"
	"errors"

	"github.com/sirupsen/logrus"
)

type transactionRepositoryMYSQL struct {
	db *sql.DB
}

func NewMysqlTransactionRepository(Conn *sql.DB) repository.TransactionRepository {
	return &transactionRepositoryMYSQL{Conn}
}

func (m *transactionRepositoryMYSQL) fetch(query string, args ...interface{}) (aggregate.Transactions, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()
	transactionDTOs := []model.TransactionDTO{}
	for rows.Next() {
		t := model.TransactionDTO{}
		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.Flag,
			// &t.UserRecieveID,
			&t.Nominal,
		)

		if err != nil {
			return nil, err
		}

		transactionDTOs = append(transactionDTOs, t)
	}

	transactions := aggregate.Transactions{}
	for _, transactionDTO := range transactionDTOs {
		transactions = append(transactions, aggregate.RebuildTransaction(
			transactionDTO.ID,
			transactionDTO.UserID,
			transactionDTO.Flag,
			0,
			transactionDTO.Nominal,
		))
	}

	return transactions, nil
}

func (m *transactionRepositoryMYSQL) GetTransaction() (res aggregate.Transactions, err error) {
	query := `SELECT id, user_id, flag, nominal FROM transactions`

	res, err = m.fetch(query)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *transactionRepositoryMYSQL) GetTransactionByID(id uint64) (res aggregate.Transactions, err error) {
	query := `SELECT id, user_id, flag, nominal FROM transactions WHERE id = ? LIMIT 1`

	res, err = m.fetch(query, id)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *transactionRepositoryMYSQL) CreateTransaction(transaction aggregate.Transaction) (uint64, error) {
	if transaction.Flag == 3 {
		query := "INSERT INTO transactions (user_id, flag, user_receive_id,nominal) VALUES(?, ?, ?, ?)"
		res, err := m.db.Exec(
			query,
			transaction.UserID,
			transaction.Flag,
			transaction.UserRecieveID,
			transaction.Nominal,
		)
		if err != nil {
			return 0, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}
		uId := uint64(id)

		return uId, nil
	}

	query := "INSERT INTO transactions (user_id, flag, nominal) VALUES (?, ?, ?)"
	res, err := m.db.Exec(
		query,
		transaction.UserID,
		transaction.Flag,
		transaction.Nominal,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	uId := uint64(id)

	return uId, nil
}

func (s *transactionRepositoryMYSQL) DeleteTransaction(id uint64) error {
	query := "DELETE FROM transactions WHERE id = ?"
	_, err := s.db.Exec(
		query,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
