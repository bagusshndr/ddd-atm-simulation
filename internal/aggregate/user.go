package aggregate

import (
	"errors"
)

type Users []*User

type User struct {
	ID     uint64
	Name   string
	Amount float64
}

func (d *User) IncreaseAmount(nominal float64) error {
	if nominal == 0 {
		return errors.New("nominal is empty")
	}

	d.Amount = d.Amount + nominal
	return nil
}

func (d *User) DecreaseAmount(nominal float64) error {
	if nominal == 0 {
		return errors.New("nominal is empty")
	}

	d.Amount = d.Amount - nominal
	return nil
}

func NewUser(name string, amount float64) (*User, error) {
	return &User{
		Name:   name,
		Amount: amount,
	}, nil

}

func RebuildUser(id uint64, name string, amount float64) *User {
	return &User{
		ID:     id,
		Name:   name,
		Amount: amount,
	}
}
