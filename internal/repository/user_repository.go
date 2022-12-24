package repository

import "ddd-to-do-list/internal/aggregate"

type UserRepository interface {
	GetUser() (res aggregate.Users, err error)
	CreateUser(user aggregate.User) (uint64, error)
	GetUserByID(id uint64) (res aggregate.Users, err error)
	UpdateUser(user aggregate.User) error
	DeleteUser(id uint64) error
}
