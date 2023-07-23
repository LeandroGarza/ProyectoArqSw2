package app

import (
	"cache/config"
	"cache/controller"
	"cache/service"
	"cache/service/repositories"
)

type Dependencies struct {
	Cachecontroller *controller.CacheController
}

func BuildDependencies() *Dependencies {
	cacheclient := repositories.NewCacheClientImpl(config.CACHEHOST, config.CACHEPORT)

	queueclient := repositories.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)
	go queueclient.Consumer(config.EXCHANGE, cacheclient)

	cacheservice := service.NewCacheServiceImpl(cacheclient)

	cachecontroller := controller.NewCacheController(cacheservice)
	return &Dependencies{
		Cachecontroller: cachecontroller,
	}
}
