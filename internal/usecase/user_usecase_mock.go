package usecase

import (
	"ddd-atm-simulation/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type UserddddUsecaseMock struct {
	mock.Mock
}

func (m *UserddddUsecaseMock) GetUser() (res aggregate.Users, err error) {
	args := m.Called()

	return args.Get(0).(aggregate.Users), args.Error(1)
}

func (m *UserddddUsecaseMock) CreateUser(user aggregate.User) (uint64, error) {
	args := m.Called(user)

	return args.Get(0).(uint64), args.Error(1)
}

func (m *UserddddUsecaseMock) GetUserByID(id uint64) (res aggregate.Users, err error) {
	args := m.Called(id)

	return args.Get(0).(aggregate.Users), args.Error(1)
}

func (m *UserddddUsecaseMock) UpdateUser(user aggregate.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *UserddddUsecaseMock) DeleteUser(id uint64) error {
	args := m.Called(id)

	return args.Error(0)
}
