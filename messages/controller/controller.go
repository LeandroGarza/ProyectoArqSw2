package controller

import (
	"messages/dto"
	"messages/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type MessageController struct {
	MessageService service.MessageService
}

func NewMessageController(messageservice service.MessageService) *MessageController {
	return &MessageController{
		MessageService: messageservice,
	}
}

func (ctrl *MessageController) CreateMessage(c *gin.Context) {
	var messagedto dto.MessageDto
	err := c.BindJSON(&messagedto)

	log.Debug(messagedto)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userid, err := strconv.Atoi(c.MustGet("userid").(string))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusConflict, err)
		return
	}
	messagedto.Userid = userid

	messagedto, er := ctrl.MessageService.CreateMessage(messagedto)
	if er != nil {
		log.Error(er)
		c.JSON(http.StatusBadRequest, er)
		return
	}
	c.JSON(http.StatusOK, messagedto)
}

func (ctrl *MessageController) GetMessagesByItem(c *gin.Context) {
	itemid := c.Param("itemid")
	messagesdto, err := ctrl.MessageService.GetMessagesByItem(itemid)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, messagesdto)
}

func (ctrl *MessageController) GetMessageById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	message, er := ctrl.MessageService.GetMessageById(id)
	if er != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, message)
}

func (ctrl *MessageController) GetMessageByUser(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	messagesdto, er := ctrl.MessageService.GetMessagesByUser(userid)
	if er != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, messagesdto)
}

func (ctrl *MessageController) DeleteMessage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	er := ctrl.MessageService.DeleteMessage(id)
	if er != nil {
		log.Error(er)
		c.JSON(http.StatusBadRequest, er)
	}

	c.JSON(http.StatusOK, id)
}

func (ctrl *MessageController) DeleteMessagesByUser(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	er := ctrl.MessageService.DeleteMessagesByUser(userid)
	if er != nil {
		log.Error(er)
		c.JSON(http.StatusBadRequest, er)
		return
	}

	c.JSON(http.StatusOK, userid)
}

func (ctrl *MessageController) ValidateToken(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	claims, err := ctrl.MessageService.ValidateToken(auth)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	c.Set("userid", claims.Id)
	c.Next()
}
