package repository

import (
	"ddd-atm-simulation/internal/repository"
	"ddd-atm-simulation/testdata/data"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type userRepositoryMysqlTest struct {
	suite.Suite
	mock           sqlmock.Sqlmock
	userRepository repository.UserRepository
}

func (t *userRepositoryMysqlTest) TestGetUser() {
	user := data.User()
	query := `SELECT id, name, amount FROM users`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
			"amount",
		}).AddRow(
			user.ID,
			user.Name,
			user.Amount,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		actualUser, err := t.userRepository.GetUser()

		t.NotNil(actualUser)
		t.NoError(err)
	})
	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New(""))

		actualUser, err := t.userRepository.GetUser()

		t.Nil(actualUser)
		t.Error(err)
	})
}

func (t *userRepositoryMysqlTest) TestGetUserByID() {
	user := data.User()
	query := `SELECT id, name, amount FROM users where id = ?`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
			"amount",
		}).AddRow(
			user.ID,
			user.Name,
			user.Amount,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(user.ID).WillReturnRows(rows)

		actualUser, err := t.userRepository.GetUserByID(user.ID)

		t.NotNil(actualUser)
		t.NoError(err)
	})
	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(user.ID).WillReturnError(errors.New("error"))

		actualUser, err := t.userRepository.GetUserByID(user.ID)

		t.Nil(actualUser)
		t.Error(err)
	})
}

func TestUserRepositoryMySQL(t *testing.T) {
	db, mock, _ := sqlmock.New()

	suite.Run(t, &userRepositoryMysqlTest{
		mock:           mock,
		userRepository: NewMysqlUserRepository(db),
	})
}
