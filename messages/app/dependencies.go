package app

import (
	"messages/config"
	"messages/controller"
	"messages/service"
	"messages/service/repositories"
)

type Dependencies struct {
	Controller *controller.MessageController
}

func BuildDependencies() *Dependencies {
	// clients
	messageclient := repositories.NewClientImpl(config.DBUSER, config.DBPASS, config.DBHOST, config.DBPORT, config.DBNAME)
	//queueclient := repositories.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)

	//go queueclient.ConsumeUserUpdates(config.EXCHANGE, messageclient)

	// services
	messageservice := service.NewMessageServiceImpl(messageclient)

	// controller
	messagecontroller := controller.NewMessageController(messageservice)

	return &Dependencies{
		Controller: messagecontroller,
	}
}
