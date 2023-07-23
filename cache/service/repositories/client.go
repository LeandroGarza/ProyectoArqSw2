package repositories

import "cache/model"

type CacheClient interface {
	InsertUserData(id int, userdata model.Data) (model.Data, error)
	GetUserData(id int) (model.Data, error)
	DeleteUserData(id int) error
}
