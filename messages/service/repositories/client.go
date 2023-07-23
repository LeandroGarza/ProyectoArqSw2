package repositories

import (
	"messages/model"
)

type MessageClient interface {
	CreateMessage(message model.Message) (model.Message, error)
	GetMessagesByItem(itemid string) (model.Messages, error)
	GetMessageById(id int) (model.Message, error)
	GetMessagesByUser(userid int) (model.Messages, error)
	DeleteMessage(id int) error
	DeleteMessagesByUser(userid int) error
	StartDbEngine()
}
