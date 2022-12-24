package model

import "ddd-to-do-list/internal/enum"

type TransactionDTO struct {
	ID            uint64    `gorm:"id"`
	UserID        uint64    `gorm:"user_id"`
	Flag          enum.Flag `gorm:"flag"`
	UserRecieveID uint64    `gorm:"user_receive_id"`
	Nominal       float64   `gorm:"nominal"`
}

func (t *TransactionDTO) TableName() string {
	return "transactions"
}
