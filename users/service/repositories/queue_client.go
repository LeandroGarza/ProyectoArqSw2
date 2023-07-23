package repositories

import e "users/utils/errors"

type QueueClient interface {
	SendMessage(userid int, action string, message string) e.ApiError
}
