package aggregate

import (
	"ddd-atm-simulation/internal/enum"
)

type Transactions []*Transaction

type Transaction struct {
	ID            uint64
	UserID        uint64
	Flag          enum.Flag
	UserRecieveID uint64
	Nominal       float64
}

func NewTransaction(userID uint64, flag enum.Flag, userReceiveID uint64, nominal float64) (*Transaction, error) {
	return &Transaction{
		UserID:        userID,
		Flag:          flag,
		UserRecieveID: userID,
		Nominal:       nominal,
	}, nil
}

func RebuildTransaction(id, userID uint64, flag enum.Flag, userReceiveID uint64, nominal float64) *Transaction {
	return &Transaction{
		ID:            id,
		UserID:        userID,
		Flag:          flag,
		UserRecieveID: userID,
		Nominal:       nominal,
	}
}
