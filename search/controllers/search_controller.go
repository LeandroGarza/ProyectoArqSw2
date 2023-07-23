package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"search/dtos"
	service "search/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type SearchController struct {
	service service.Service
}

func NewSearchController(service service.Service) *SearchController {
	return &SearchController{
		service: service,
	}
}

func (sc *SearchController) Search(c *gin.Context) {
	query := c.Param("searchQuery")
	itemsDto, err := sc.service.Search(query)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)
}

func (sc *SearchController) SearchByUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	itemsdto, er := sc.service.SearchByUserId(id)
	if er != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, itemsdto)
}

func (sc *SearchController) InsertItems(c *gin.Context) {
	var itemsdto dtos.ItemsDto
	err := c.BindJSON(&itemsdto)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	itemsdto, er := sc.service.InsertItems(itemsdto)
	if er != nil {
		log.Error(er)
		c.JSON(http.StatusBadRequest, er)
		return
	}

	c.JSON(http.StatusOK, itemsdto)
}

func (sc *SearchController) DeleteAll(c *gin.Context) {
	err := sc.service.DeleteAll()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, err)
}

func (sc *SearchController) DeleteByUserId(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	er := sc.service.DeleteByUserId(userid)
	if er != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, er)
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Se eliminaron los documentos con userid: %v", userid))
}
