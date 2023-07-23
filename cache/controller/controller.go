package controller

import (
	"cache/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CacheController struct {
	CacheService service.CacheService
}

func NewCacheController(cacheservice service.CacheService) *CacheController {
	return &CacheController{
		CacheService: cacheservice,
	}
}

func (ctrl *CacheController) GetUserData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("error converting param id to int")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userdto, er := ctrl.CacheService.GetUserData(id)
	if er != nil {
		fmt.Println("error getting user")
		c.JSON(http.StatusBadRequest, er)
		return
	}

	c.JSON(http.StatusOK, userdto)
}
