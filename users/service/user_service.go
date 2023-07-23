package services

import (
	"users/dto"
	e "users/utils/errors"
)

type MessageService interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	DeleteUser(id int) e.ApiError
	InsertUser(userdto dto.UserDto) (dto.UserDto, e.ApiError)
	Login(logindto dto.LoginRequestDto) (dto.LoginResponseDto, e.ApiError)
	//InsertItemsById(itemdto dto.ItemsDto, token string) (dto.ItemsDto, e.ApiError)
}
