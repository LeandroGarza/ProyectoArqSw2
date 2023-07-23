package repositories

import (
	"context"

	dtos "items/dtos"
	e "items/utils/errors"
)

type Repository interface {
	Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError)
	InsertItem(ctx context.Context, Item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	InsertItems(ctx context.Context, Items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError)
	Update(ctx context.Context, Item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	Delete(ctx context.Context, id string) e.ApiError
	DeleteByUserId(ctx context.Context, userid int) e.ApiError
}
