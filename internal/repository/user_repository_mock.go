package repository

import (
	"ddd-atm-simulation/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUser() (res aggregate.Users, err error) {
	args := m.Called()

	return args.Get(0).(aggregate.Users), args.Error(1)
}

func (m *UserRepositoryMock) CreateUser(user aggregate.User) (uint64, error) {
	args := m.Called(user)

	return args.Get(0).(uint64), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByID(id uint64) (res aggregate.Users, err error) {
	args := m.Called(id)

	return args.Get(0).(aggregate.Users), args.Error(1)
}

func (m *UserRepositoryMock) UpdateUser(user aggregate.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *UserRepositoryMock) DeleteUser(id uint64) error {
	args := m.Called(id)

	return args.Error(0)
}
