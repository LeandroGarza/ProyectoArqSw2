package service

import (
	"messages/dto"
	"messages/model"
	"messages/service/repositories"
	"messages/utils/errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type MessageServiceImpl struct {
	MessageClient repositories.MessageClient
}

func NewMessageServiceImpl(messageclient repositories.MessageClient) *MessageServiceImpl {
	messageclient.StartDbEngine()
	return &MessageServiceImpl{
		MessageClient: messageclient,
	}
}

func (s *MessageServiceImpl) CreateMessage(messagedto dto.MessageDto) (dto.MessageDto, errors.ApiError) {
	var message model.Message

	message.Userid = messagedto.Userid
	message.Itemid = messagedto.Itemid
	message.Content = messagedto.Content
	message.Createdat = time.Now().Format("2006-01-02 15:04:05")

	message, err := s.MessageClient.CreateMessage(message)
	if err != nil {
		return dto.MessageDto{}, errors.NewInternalServerApiError("Error creating new message", err)
	}
	messagedto.Id = message.Id

	return messagedto, nil
}

func (s *MessageServiceImpl) GetMessagesByItem(itemid string) (dto.MessagesDto, errors.ApiError) {
	var messagesdto dto.MessagesDto

	messages, err := s.MessageClient.GetMessagesByItem(itemid)
	if err != nil {
		return dto.MessagesDto{}, errors.NewInternalServerApiError("Error getting messages by item", err)
	}

	for _, message := range messages {
		var messagedto dto.MessageDto

		messagedto.Id = message.Id
		messagedto.Userid = message.Userid
		messagedto.Itemid = message.Itemid
		messagedto.Content = message.Content
		messagedto.Createdat = message.Createdat

		messagesdto = append(messagesdto, messagedto)
	}

	return messagesdto, nil
}

func (s *MessageServiceImpl) GetMessageById(id int) (dto.MessageDto, errors.ApiError) {
	var messagedto dto.MessageDto

	message, err := s.MessageClient.GetMessageById(id)
	if err != nil {
		return dto.MessageDto{}, errors.NewBadRequestApiError("Failed to get message by id")
	}

	messagedto.Id = message.Id
	messagedto.Userid = message.Userid
	messagedto.Itemid = message.Itemid
	messagedto.Content = message.Content
	messagedto.Createdat = message.Createdat

	return messagedto, nil
}

func (s *MessageServiceImpl) GetMessagesByUser(userid int) (dto.MessagesDto, errors.ApiError) {
	var messagesdto dto.MessagesDto

	messages, err := s.MessageClient.GetMessagesByUser(userid)
	if err != nil {
		return dto.MessagesDto{}, errors.NewBadRequestApiError("failed to get messages by user")
	}

	for _, message := range messages {
		var messagedto dto.MessageDto
		messagedto.Id = message.Id
		messagedto.Userid = message.Userid
		messagedto.Itemid = message.Itemid
		messagedto.Content = message.Content
		messagedto.Createdat = message.Createdat

		messagesdto = append(messagesdto, messagedto)
	}

	return messagesdto, nil
}

func (v *MessageServiceImpl) ValidateToken(authToken string) (*jwt.StandardClaims, errors.ApiError) {
	tokenString := strings.Split(authToken, " ")[1]
	claims := &jwt.StandardClaims{}
	jwtKey := []byte("tengohambre") // VARIABLE DE ENTORNO!

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, errors.NewInternalServerApiError("error parseando token", err)
	}

	if !token.Valid {
		return nil, errors.NewUnauthorizedApiError("token no valido")
	}

	// Verifica si el token ha expirado
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, errors.NewUnauthorizedApiError("token expirado")
	}

	return claims, nil
}

func (s *MessageServiceImpl) DeleteMessage(id int) errors.ApiError {
	err := s.MessageClient.DeleteMessage(id)
	if err != nil {
		return errors.NewBadRequestApiError("failed to delete message")
	}
	return nil
}

func (s *MessageServiceImpl) DeleteMessagesByUser(userid int) errors.ApiError {
	err := s.MessageClient.DeleteMessage(userid)
	if err != nil {
		return errors.NewBadRequestApiError("failed to delete messages")
	}
	return nil
}
