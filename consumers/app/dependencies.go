package app

import (
	"consumers/client"
	"consumers/config"
)

func BuildDependencies() {
	queueclient := client.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)
	go queueclient.ConsumeItems()
	go queueclient.ConsumeUserUpdates(config.EXCHANGE, config.ENDPOINTS)
}
