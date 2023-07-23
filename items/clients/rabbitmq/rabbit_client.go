package rabbitmq

import (
	"context"

	dtos "items/dtos"
	"items/services/repositories"
)

type QueueClient interface {
	PublishItem(ctx context.Context, item dtos.ItemDto) error
	PublishItems(ctx context.Context, items dtos.ItemsDto) error
	ConsumeUserUpdate(exchange string, ccache repositories.Repository, mongo repositories.Repository)
}
