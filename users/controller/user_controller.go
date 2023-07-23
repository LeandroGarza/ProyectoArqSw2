package controller

import (
	"users/dto"
	service "users/service"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserController struct {
	service service.MessageService
}

func NewUserController(service service.MessageService) *UserController {
	return &UserController{
		service: service,
	}
}

func (ctrl *UserController) GetUserById(c *gin.Context) {
	log.Debug("user id: " + c.Param("id"))

	var userdto dto.UserDto
	id, _ := strconv.Atoi(c.Param("id"))
	userdto, err := ctrl.service.GetUserById(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userdto)
}

func (ctrl *UserController) InsertUser(c *gin.Context) {
	var userdto dto.UserDto
	err := c.BindJSON(&userdto)

	log.Debug(userdto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userdto, er := ctrl.service.InsertUser(userdto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userdto)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	log.Debug("user id: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.service.DeleteUser(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (ctrl *UserController) Login(c *gin.Context) {
	var logindto dto.LoginRequestDto
	err := c.BindJSON(&logindto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Debug(logindto)

	var loginresponsedto dto.LoginResponseDto
	loginresponsedto, er := ctrl.service.Login(logindto)
	if er != nil {
		if er.Status() == 400 {
			c.JSON(http.StatusBadRequest, er.Error())
			return
		}
		c.JSON(http.StatusForbidden, er.Error())
		return
	}
	c.JSON(http.StatusOK, loginresponsedto)
}
