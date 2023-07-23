package services

import (
	"context"

	dtos "items/dtos"
	e "items/utils/errors"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError)
	InsertItem(ctx context.Context, Item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	InsertItems(ctx context.Context, Items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError)
	DeleteItemsByUserId(ctx context.Context, userid int) e.ApiError
	ValidateToken(authToken string) (*jwt.StandardClaims, error)
}
