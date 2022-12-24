package usecase

import (
	"ddd-atm-simulation/internal/aggregate"
	"ddd-atm-simulation/internal/repository"
)

type userUsecase struct {
	repoUser repository.UserRepository
}

func (u *userUsecase) GetUser() (res aggregate.Users, err error) {
	res = aggregate.Users{}
	return res, nil
}

func (u *userUsecase) CreateUser(user aggregate.User) (uint64, error) {
	return 0, nil
}

func (u *userUsecase) GetUserByID(id uint64) (res aggregate.Users, err error) {
	res = aggregate.Users{}
	return res, nil
}

func (u *userUsecase) DeleteUser(id uint64) error {
	return nil
}

func NewUserUsecase(repoUser repository.UserRepository) UserUsecase {
	return &userUsecase{
		repoUser: repoUser,
	}
}
