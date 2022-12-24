package model

type UserDTO struct {
	ID     uint64  `gorm:"id"`
	Name   string  `gorm:"name"`
	Amount float64 `gorm:"amount"`
}

func (t *UserDTO) TableName() string {
	return "users"
}
