package app

import (
	config "users/config"
	controller "users/controller"
	services "users/service"
	repositories "users/service/repositories"
)

type Dependencies struct {
	UserController *controller.UserController
}

func BuildDependencies() *Dependencies {
	// repositories
	userclient := repositories.NewUserClient(config.DBUSER, config.DBPASS, config.DBHOST, config.DBPORT, config.DBNAME)
	userclient.StartDbEngine()
	queueclient := repositories.NewQueueClientImpl(config.RABBITUSER, config.RABBITPASSWORD, config.DBHOST, config.RABBITPORT)

	// services
	service := services.NewUserServiceImpl(userclient, queueclient)

	// controllers
	usercontroller := controller.NewUserController(service)

	return &Dependencies{
		UserController: usercontroller,
	}
}
