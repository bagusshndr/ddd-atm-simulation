package usecase

import "ddd-to-do-list/internal/aggregate"

type UserUsecase interface {
	GetUser() (res aggregate.Users, err error)
	CreateUser(user aggregate.User) (uint64, error)
	GetUserByID(id uint64) (res aggregate.Users, err error)
	DeleteUser(id uint64) error
}
