package app

import (
	clients "items/clients/rabbitmq"
	"items/config"
	controllers "items/controllers"
	service "items/services"
	repositories "items/services/repositories"
)

type Dependencies struct {
	ItemController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	ccache := repositories.NewCCache(config.CCMAXSIZE, config.CCITEMSTOPRUNE, config.CCDEFAULTTTL)
	mongo := repositories.NewMongoDB(config.DBHOST, config.DBPORT, config.COLLECTION)
	queue := clients.NewRabbitmq(config.RABBITHOST, config.RABBITPORT)

	// Services
	service := service.NewServiceImpl(ccache, mongo, queue)

	// consumer
	// go queue.ConsumeUserUpdate(config.EXCHANGE, ccache, mongo)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}
