package service

import (
	"cache/dto"
	"cache/utils/errors"
)

type CacheService interface {
	GetUserData(id int) (dto.UserDto, errors.ApiError)
}
