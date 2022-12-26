package usecase

import (
	"ddd-atm-simulation/internal/aggregate"
	"ddd-atm-simulation/internal/repository"
)

type userUsecase struct {
	repoUser repository.UserRepository
}

func (u *userUsecase) GetUser() (res aggregate.Users, err error) {

	return res, nil
}

func (u *userUsecase) CreateUser(user aggregate.User) (uint64, error) {
	uid, err := u.repoUser.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func (u *userUsecase) GetUserByID(id uint64) (res aggregate.Users, err error) {
	res, err = u.repoUser.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userUsecase) GetUserByName(name string) (res aggregate.Users, err error) {
	res, _ = u.repoUser.GetUserByName(name)
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
