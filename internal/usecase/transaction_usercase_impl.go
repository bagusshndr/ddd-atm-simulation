package usecase

import (
	"ddd-atm-simulation/internal/aggregate"
	"ddd-atm-simulation/internal/repository"
)

type transactionUsecase struct {
	repoUser        repository.UserRepository
	repoTransaction repository.TransactionRepository
}

func (u *transactionUsecase) GetTransaction() (res aggregate.Transactions, err error) {
	res, _ = u.repoTransaction.GetTransaction()
	return res, nil
}

func (u *transactionUsecase) CreateTransaction(transaction aggregate.Transaction) error {
	user, err := u.repoUser.GetUserByID(transaction.UserID)
	if err != nil {
		return err
	}

	switch transaction.Flag {
	case 1:
		_, err := u.repoTransaction.CreateTransaction(transaction)
		if err != nil {
			return err
		}
		user[0].IncreaseAmount(transaction.Nominal)
		u.repoUser.UpdateUser(*user[0])
	case 2:
		_, err := u.repoTransaction.CreateTransaction(transaction)
		if err != nil {
			return err
		}
		user[0].DecreaseAmount(transaction.Nominal)
		u.repoUser.UpdateUser(*user[0])
	case 3:
		_, err := u.repoUser.GetUserByID(transaction.UserRecieveID)
		if err != nil {
			return err
		}
		u.repoTransaction.CreateTransaction(transaction)
		u.repoUser.GetUserByID(user[0].ID)
	}

	return nil
}

func NewTransactionUsecase(repoUser repository.UserRepository, repoTransaction repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		repoUser:        repoUser,
		repoTransaction: repoTransaction,
	}
}
