package repositories

import (
	"fmt"
	"messages/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type ClientImpl struct {
	Db *gorm.DB
}

func NewClientImpl(DBuser string, DBpass string, DBhost string, DBport int, DBname string) *ClientImpl {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", DBuser, DBpass, DBhost, DBport, DBname))
	if err != nil {
		panic(fmt.Sprintf("Error inicializing database, %v", err))
	}
	return &ClientImpl{
		Db: db,
	}
}

func (c *ClientImpl) StartDbEngine() {
	c.Db.AutoMigrate(&model.Message{})
	log.Info("Finishing migrating tables")
}

func (c *ClientImpl) CreateMessage(message model.Message) (model.Message, error) {
	result := c.Db.Create(&message)
	if result.Error != nil {
		log.Error(result.Error)
		return model.Message{}, result.Error
	}

	log.Debug("Message created: %v", message)
	return message, nil
}

func (c *ClientImpl) GetMessagesByItem(itemid string) (model.Messages, error) {
	var messages model.Messages
	result := c.Db.Where("itemid = ?", itemid).Find(&messages)
	if result.Error != nil {
		return model.Messages{}, result.Error
	}

	log.Debug("Messages found: %v", messages)
	return messages, nil
}

func (c *ClientImpl) GetMessageById(id int) (model.Message, error) {
	var message model.Message

	result := c.Db.Where("id = ?", id).First(&message)
	if result.Error != nil {
		return model.Message{}, result.Error
	}

	log.Debug("message found: %v", message)
	return message, nil
}

func (c *ClientImpl) GetMessagesByUser(userid int) (model.Messages, error) {
	var messages model.Messages

	result := c.Db.Where("userid = ?", userid).Find(&messages)
	if result.Error != nil {
		return model.Messages{}, result.Error
	}

	log.Debug("messages found: %v", messages)
	return messages, nil
}

func (c *ClientImpl) DeleteMessage(id int) error {
	var message model.Message
	c.Db.Where("id = ?", id).First(&message)
	result := c.Db.Delete(&message)
	return result.Error
}

func (c *ClientImpl) DeleteMessagesByUser(userid int) error {
	var messages model.Messages
	c.Db.Where("userid = ?", userid).Find(&messages)
	result := c.Db.Delete(&messages)
	return result.Error
}
