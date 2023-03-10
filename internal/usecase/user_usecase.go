package usecase

import "ddd-atm-simulation/internal/aggregate"

type UserUsecase interface {
	GetUser() (res aggregate.Users, err error)
	CreateUser(user aggregate.User) (uint64, error)
	GetUserByID(id uint64) (res aggregate.Users, err error)
	GetUserByName(name string) (res aggregate.Users, err error)
	DeleteUser(id uint64) error
}
