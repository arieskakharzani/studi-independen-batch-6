package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	data := u.db.Create(&user)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	users := &model.User{}
	data := u.db.Where(" username = ? AND password = ?", user.Username, user.Password).First(&users)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}
