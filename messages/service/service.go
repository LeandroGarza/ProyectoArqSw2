package service

import (
	"messages/dto"
	"messages/utils/errors"

	"github.com/golang-jwt/jwt"
)

type MessageService interface {
	CreateMessage(messagedto dto.MessageDto) (dto.MessageDto, errors.ApiError)
	GetMessagesByItem(itemid string) (dto.MessagesDto, errors.ApiError)
	GetMessageById(id int) (dto.MessageDto, errors.ApiError)
	GetMessagesByUser(userid int) (dto.MessagesDto, errors.ApiError)
	ValidateToken(authToken string) (*jwt.StandardClaims, errors.ApiError)
	DeleteMessage(id int) errors.ApiError
	DeleteMessagesByUser(userid int) errors.ApiError
}
