package controller

import (
	dtos "items/dtos"
	service "items/services"
	e "items/utils/errors"
	"strconv"

	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Get(c *gin.Context) {
	item, apiErr := ctrl.service.Get(c.Request.Context(), c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *Controller) InsertItem(c *gin.Context) {
	var item dtos.ItemDto
	if err := c.BindJSON(&item); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := ctrl.service.InsertItem(c.Request.Context(), item)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (ctrl *Controller) InsertItems(c *gin.Context) {
	var itemsdto dtos.ItemsDto
	if err := c.BindJSON(&itemsdto); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	userid := c.MustGet("userid").(int)

	for i := range itemsdto {
		itemsdto[i].UserId = userid
	}

	items, apirErr := ctrl.service.InsertItems(c.Request.Context(), itemsdto)
	if apirErr != nil {
		c.JSON(apirErr.Status(), apirErr)
		return
	}

	c.JSON(http.StatusCreated, items)
}

func (ctrl *Controller) DeleteByUserId(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	er := ctrl.service.DeleteItemsByUserId(c.Request.Context(), userid)
	if er != nil {
		c.JSON(http.StatusBadRequest, er)
		return
	}

	c.JSON(http.StatusOK, "items deleted")
}

func (ctrl *Controller) ValidateToken(c *gin.Context) {

	auth := c.GetHeader("Authorization")

	claims, err := ctrl.service.ValidateToken(auth)
	if err != nil {
		apiErr := e.NewUnauthorizedApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.Set("userid", claims.Id)
	c.Next()
}
