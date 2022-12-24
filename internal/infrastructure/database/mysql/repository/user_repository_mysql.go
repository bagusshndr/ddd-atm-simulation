package repository

import (
	"database/sql"
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/infrastructure/database/mysql/model"
	"ddd-to-do-list/internal/repository"
	"errors"

	"github.com/sirupsen/logrus"
)

type userRepositoryMYSQL struct {
	db *sql.DB
}

func NewMysqlUserRepository(Conn *sql.DB) repository.UserRepository {
	return &userRepositoryMYSQL{Conn}
}

func (m *userRepositoryMYSQL) fetch(query string, args ...interface{}) (aggregate.Users, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()
	userDTOs := []model.UserDTO{}
	for rows.Next() {
		t := model.UserDTO{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Amount,
		)

		if err != nil {
			return nil, err
		}

		userDTOs = append(userDTOs, t)
	}

	users := aggregate.Users{}
	for _, userDTO := range userDTOs {
		users = append(users, aggregate.RebuildUser(
			userDTO.ID,
			userDTO.Name,
			userDTO.Amount,
		))
	}

	return users, nil
}

func (m *userRepositoryMYSQL) GetUser() (res aggregate.Users, err error) {
	query := `SELECT id, name, amount FROM users`

	res, err = m.fetch(query)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *userRepositoryMYSQL) GetUserByID(id uint64) (res aggregate.Users, err error) {
	query := `SELECT id, name, amount FROM users WHERE id = ? LIMIT 1`

	res, err = m.fetch(query, id)
	if err != nil {
		return nil, errors.New("")
	}
	return
}

func (m *userRepositoryMYSQL) CreateUser(user aggregate.User) (uint64, error) {
	query := "INSERT INTO users (name, amount) VALUES(?, ?)"
	res, err := m.db.Exec(
		query,
		user.Name,
		user.Amount,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	uId := uint64(id)

	return uId, nil
}

func (m *userRepositoryMYSQL) UpdateUser(user aggregate.User) error {
	query := "UPDATE users SET name = ?, amount = ? WHERE id = ?"
	_, err := m.db.Exec(
		query,
		user.Name,
		user.Amount,
		user.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *userRepositoryMYSQL) DeleteUser(id uint64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := s.db.Exec(
		query,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
