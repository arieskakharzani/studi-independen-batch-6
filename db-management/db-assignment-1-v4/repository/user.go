package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"errors"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := u.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CheckAvail(user model.User) error {
	query := "SELECT * FROM users WHERE username = $1 AND password = $2 LIMIT 1"
	err := u.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("User Not Found!")
		}
		return err
	}
	return nil
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
