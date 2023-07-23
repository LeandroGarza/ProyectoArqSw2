package repositories

import "users/model"

type Client interface {
	GetUserById(id int) (model.User, error)
	DeleteUser(user model.User) error
	GetUserByUsername(username string) (model.User, error)
	InsertUser(user model.User) model.User
}
