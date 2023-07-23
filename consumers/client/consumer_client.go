package client

type ConsumerService interface {
	ConsumeUserUpdates()
	ConsumeItems()
}
