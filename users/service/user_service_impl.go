package services

import (
	"fmt"
	"users/dto"
	"users/model"
	client "users/service/repositories"
	e "users/utils/errors"
	utils "users/utils/login"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userclient  client.Client
	queueclient client.QueueClient
}

func NewUserServiceImpl(userclient client.Client, queueclient client.QueueClient) *UserServiceImpl {
	return &UserServiceImpl{
		userclient:  userclient,
		queueclient: queueclient,
	}
}

func (s *UserServiceImpl) GetUserById(id int) (dto.UserDto, e.ApiError) {
	var userdto dto.UserDto
	var user, err = s.userclient.GetUserById(id)
	if err != nil {
		return dto.UserDto{}, e.NewBadRequestApiError(err.Error())
	}

	userdto.Username = user.Username
	userdto.Email = user.Email
	userdto.Id = user.Id

	return userdto, nil
}

func (s *UserServiceImpl) InsertUser(userdto dto.UserDto) (dto.UserDto, e.ApiError) {
	var user model.User

	hashedpassword, er := bcrypt.GenerateFromPassword([]byte(userdto.Password), bcrypt.DefaultCost)
	if er != nil {
		return userdto, e.NewInternalServerApiError("Error generating password hash", er)
	}

	user.Username = userdto.Username
	user.Email = userdto.Email
	user.Id = userdto.Id
	user.Password = string(hashedpassword)

	user = s.userclient.InsertUser(user)
	userdto.Id = user.Id

	err := s.queueclient.SendMessage(userdto.Id, "create", fmt.Sprintf("%d", userdto.Id))
	if err != nil {
		return userdto, e.NewInternalServerApiError("Error sending user creation message", err)
	}
	return userdto, nil
}

func (s *UserServiceImpl) DeleteUser(id int) e.ApiError {
	err := s.queueclient.SendMessage(id, "delete", fmt.Sprintf("%d", id))
	if err != nil {
		return e.NewInternalServerApiError("Error sending user deletion message, aborting", err)
	}

	var user model.User
	user.Id = id
	er := s.userclient.DeleteUser(user)
	if er != nil {
		return e.NewInternalServerApiError("Error deleting user", er)
	}

	return nil
}

func (s *UserServiceImpl) Login(loginDto dto.LoginRequestDto) (dto.LoginResponseDto, e.ApiError) {
	var user model.User
	var loginResponseDto dto.LoginResponseDto
	user, err := s.userclient.GetUserByUsername(loginDto.Username)
	if err != nil {
		return loginResponseDto, e.NewBadRequestApiError("User not found")
	}

	if !utils.ComparePasswords(user.Password, loginDto.Password) {
		return dto.LoginResponseDto{}, e.NewUnauthorizedApiError("wrong password")
	}

	tokenstring, err := utils.GenerateToken(user.Id)
	if err != nil {
		return dto.LoginResponseDto{}, e.NewInternalServerApiError("Error generating token", err)
	}

	loginResponseDto.Token = tokenstring
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}
