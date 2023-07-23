package app

import (
	"search/controllers"
	"search/services"
	repositories "search/services/repositories"
)

type Dependencies struct {
	SearchController *controllers.SearchController
}

func BuildDependencies() *Dependencies {
	// repositories
	searchclient := repositories.NewSearchClient()
	//queueclient := repositories.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, "localhost", config.RABBITPORT)

	// services
	service := services.NewSearchService(searchclient)

	// controllers
	searchcontroller := controllers.NewSearchController(service)

	// consumers
	//go queueclient.ConsumeItems()
	//go queueclient.ConsumeUserUpdates(config.EXCHANGE, searchclient)

	return &Dependencies{
		SearchController: searchcontroller,
	}
}
